/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/pkg/errors"
)

// testMeta stands in for the *ovh.Config that configureSDKProvider returns.
const testMeta = "meta"

func newTestCache() *providerMetaCache {
	return &providerMetaCache{entries: make(map[string]*metaEntry)}
}

func TestGetOrConfigureCachesResult(t *testing.T) {
	c := newTestCache()
	var calls int32

	configure := func() (any, error) {
		atomic.AddInt32(&calls, 1)
		return testMeta, nil
	}

	for i := 0; i < 3; i++ {
		meta, err := c.getOrConfigure("key", configure)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if meta != testMeta {
			t.Fatalf("got meta %v, want %q", meta, testMeta)
		}
	}

	if got := atomic.LoadInt32(&calls); got != 1 {
		t.Errorf("configure called %d times, want 1", got)
	}
}

func TestGetOrConfigureConfiguresPerKey(t *testing.T) {
	c := newTestCache()
	var calls int32

	for _, key := range []string{"a", "b", "a"} {
		if _, err := c.getOrConfigure(key, func() (any, error) {
			atomic.AddInt32(&calls, 1)
			return key, nil
		}); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	if got := atomic.LoadInt32(&calls); got != 2 {
		t.Errorf("configure called %d times, want 2 (one per distinct key)", got)
	}
}

// A burst of concurrent Connects on a cold cache must result in a single
// configuration call, since configuring hits the OVH API.
func TestGetOrConfigureSingleFlight(t *testing.T) {
	c := newTestCache()
	var calls int32

	release := make(chan struct{})
	configure := func() (any, error) {
		atomic.AddInt32(&calls, 1)
		<-release
		return testMeta, nil
	}

	const n = 20
	var wg sync.WaitGroup
	errs := make([]error, n)
	metas := make([]any, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			metas[i], errs[i] = c.getOrConfigure("key", configure)
		}(i)
	}

	// Let the goroutines pile up on the Once before letting configure return.
	time.Sleep(50 * time.Millisecond)
	close(release)
	wg.Wait()

	if got := atomic.LoadInt32(&calls); got != 1 {
		t.Errorf("configure called %d times, want 1", got)
	}
	for i := 0; i < n; i++ {
		if errs[i] != nil {
			t.Fatalf("caller %d got error: %v", i, errs[i])
		}
		if metas[i] != testMeta {
			t.Errorf("caller %d got meta %v, want %q", i, metas[i], testMeta)
		}
	}
}

// A failed configuration must not be latched: the next Connect has to retry
// rather than replay the same error forever.
func TestGetOrConfigureDoesNotCacheErrors(t *testing.T) {
	c := newTestCache()
	var calls int32

	configure := func() (any, error) {
		if atomic.AddInt32(&calls, 1) == 1 {
			return nil, errors.New("boom")
		}
		return testMeta, nil
	}

	if _, err := c.getOrConfigure("key", configure); err == nil {
		t.Fatal("expected an error from the first call")
	}

	c.mu.Lock()
	_, cached := c.entries["key"]
	c.mu.Unlock()
	if cached {
		t.Error("failed entry was left in the cache")
	}

	meta, err := c.getOrConfigure("key", configure)
	if err != nil {
		t.Fatalf("retry failed: %v", err)
	}
	if meta != testMeta {
		t.Fatalf("got meta %v, want %q", meta, testMeta)
	}
	if got := atomic.LoadInt32(&calls); got != 2 {
		t.Errorf("configure called %d times, want 2", got)
	}
}

func TestCleanupEvictsIdleEntries(t *testing.T) {
	c := newTestCache()

	configure := func() (any, error) { return testMeta, nil }
	if _, err := c.getOrConfigure("idle", configure); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, err := c.getOrConfigure("fresh", configure); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	c.mu.Lock()
	c.entries["idle"].lastUsed = time.Now().Add(-2 * metaCacheIdleTTL)
	c.mu.Unlock()

	c.cleanup()

	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.entries["idle"]; ok {
		t.Error("idle entry was not evicted")
	}
	if _, ok := c.entries["fresh"]; !ok {
		t.Error("fresh entry was evicted")
	}
}

func TestProviderMetaCacheKey(t *testing.T) {
	base := map[string]any{
		"endpoint":         "ovh-eu",
		"user_agent_extra": "Crossplane/edixos/provider-ovh",
		"access_token":     "token-1",
	}

	key, err := providerMetaCacheKey("pc:ns/secret", base)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Same configuration, freshly built map with keys in a different order.
	same, err := providerMetaCacheKey("pc:ns/secret", map[string]any{
		"access_token":     "token-1",
		"endpoint":         "ovh-eu",
		"user_agent_extra": "Crossplane/edixos/provider-ovh",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if key != same {
		t.Error("key is not stable for an equivalent configuration")
	}

	// A rotated OAuth token must produce a new key so that we never hand out a
	// meta holding the previous token.
	rotated, err := providerMetaCacheKey("pc:ns/secret", map[string]any{
		"endpoint":         "ovh-eu",
		"user_agent_extra": "Crossplane/edixos/provider-ovh",
		"access_token":     "token-2",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if key == rotated {
		t.Error("rotated credentials reused the same cache key")
	}

	// Distinct ProviderConfigs must not share an entry.
	other, err := providerMetaCacheKey("other:ns/secret", base)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if key == other {
		t.Error("different ProviderConfigs share a cache key")
	}
}
