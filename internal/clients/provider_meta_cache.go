/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ovh/terraform-provider-ovh/v2/ovh"
	"github.com/pkg/errors"
)

const (
	// Configured provider metas are evicted after this long without use.
	// Entries are keyed by the effective provider configuration, so a live
	// entry can never hold stale credentials; the TTL only bounds memory as
	// OAuth access tokens rotate and produce new keys.
	metaCacheIdleTTL = time.Hour

	errConfigureSDKProvider = "cannot configure the OVH SDKv2 provider"
)

// metaEntry holds a configured SDKv2 provider meta (an *ovh.Config). The
// embedded sync.Once makes the entry act as a promise: the first caller for a
// key configures the provider while the rest block and share the result, so a
// burst of reconciles on a cold cache costs a single OVH API call.
type metaEntry struct {
	once sync.Once
	meta any
	err  error

	// lastUsed is only read and written while providerMetaCache.mu is held.
	lastUsed time.Time
}

// providerMetaCache caches configured SDKv2 provider metas. Without it we would
// run ovh.ConfigureContextFunc -- which calls Config.loadAndValidate and hits
// GET /auth/details -- on every Connect of every managed resource.
type providerMetaCache struct {
	mu      sync.Mutex
	entries map[string]*metaEntry
}

// Global cache so configured providers are reused across reconciliations.
var globalProviderMetaCache = &providerMetaCache{
	entries: make(map[string]*metaEntry),
}

// getOrConfigure returns the cached meta for key, calling configure on a miss.
// configure runs at most once per key, outside the cache lock.
func (c *providerMetaCache) getOrConfigure(key string, configure func() (any, error)) (any, error) {
	c.mu.Lock()
	e, ok := c.entries[key]
	if !ok {
		e = &metaEntry{}
		c.entries[key] = e
	}
	e.lastUsed = time.Now()
	c.mu.Unlock()

	e.once.Do(func() {
		e.meta, e.err = configure()
	})

	if e.err != nil {
		// Don't latch the failure: drop the entry so the next Connect retries
		// instead of replaying the same error forever.
		c.mu.Lock()
		if c.entries[key] == e {
			delete(c.entries, key)
		}
		c.mu.Unlock()
		return nil, e.err
	}

	return e.meta, nil
}

// cleanup removes entries that have not been used within the idle TTL.
func (c *providerMetaCache) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, e := range c.entries {
		if now.Sub(e.lastUsed) > metaCacheIdleTTL {
			delete(c.entries, key)
		}
	}
}

// configureSDKProvider builds an OVH SDKv2 provider, configures it with cfg and
// returns the resulting provider meta. upjet configures the Framework provider
// itself but has no equivalent step for SDKv2, where it passes terraform.Setup's
// Meta straight through to every resource CRUD function.
func configureSDKProvider(ctx context.Context, cfg map[string]any) (any, error) {
	p := ovh.Provider()
	if diags := p.Configure(ctx, tfsdk.NewResourceConfigRaw(cfg)); diags.HasError() {
		return nil, errors.New(errConfigureSDKProvider + ": " + diagnosticsMessage(diags))
	}

	meta := p.Meta()
	if meta == nil {
		return nil, errors.New(errConfigureSDKProvider + ": provider returned no meta")
	}

	return meta, nil
}

// diagnosticsMessage renders error diagnostics into a single message.
func diagnosticsMessage(diags diag.Diagnostics) string {
	msgs := make([]string, 0, len(diags))
	for _, d := range diags {
		if d.Severity != diag.Error {
			continue
		}
		if d.Detail == "" {
			msgs = append(msgs, d.Summary)
			continue
		}
		msgs = append(msgs, d.Summary+": "+d.Detail)
	}
	return strings.Join(msgs, "; ")
}

// providerMetaCacheKey derives a cache key from the ProviderConfig identity and
// the effective provider configuration. Hashing the configuration means a
// rotated OAuth access token yields a new key, so a cached meta is never handed
// out with credentials that have since changed.
func providerMetaCacheKey(cacheKey string, cfg map[string]any) (string, error) {
	// encoding/json sorts map keys, so this is deterministic.
	raw, err := json.Marshal(cfg)
	if err != nil {
		return "", errors.Wrap(err, "cannot serialize provider configuration for caching")
	}

	sum := sha256.Sum256(raw)
	return cacheKey + "#" + hex.EncodeToString(sum[:]), nil
}

// configuredProviderMeta returns the provider meta to hand to upjet's SDKv2
// external client, reusing a cached one when the configuration is unchanged.
func configuredProviderMeta(ctx context.Context, cacheKey string, cfg map[string]any) (any, error) {
	key, err := providerMetaCacheKey(cacheKey, cfg)
	if err != nil {
		return nil, err
	}
	return globalProviderMetaCache.getOrConfigure(key, func() (any, error) {
		return configureSDKProvider(ctx, cfg)
	})
}

func init() {
	// Periodically evict provider metas that have fallen out of use.
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			globalProviderMetaCache.cleanup()
		}
	}()
}
