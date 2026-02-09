/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/ovh/go-ovh/ovh"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	// Cache TTL for OAuth access tokens (50 minutes, tokens valid for 1 hour)
	tokenCacheTTL = 50 * time.Minute
)

// tokenCache holds cached OAuth access tokens to avoid repeated authentication
type tokenCache struct {
	mu    sync.RWMutex
	cache map[string]*cachedToken
}

// cachedToken represents a cached access token with expiration
type cachedToken struct {
	accessToken string
	expiresAt   time.Time
}

// Global token cache to reuse OAuth access tokens across reconciliations
var globalTokenCache = &tokenCache{
	cache: make(map[string]*cachedToken),
}

// get retrieves a cached token if it exists and is not expired
func (c *tokenCache) get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	cached, exists := c.cache[key]
	if !exists {
		return "", false
	}

	// Check if expired (with 5-minute buffer)
	if time.Now().Add(5 * time.Minute).After(cached.expiresAt) {
		return "", false
	}

	return cached.accessToken, true
}

// set stores a token in the cache with TTL
func (c *tokenCache) set(key, token string, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = &cachedToken{
		accessToken: token,
		expiresAt:   time.Now().Add(ttl),
	}
}

// cleanup removes expired entries from the cache
func (c *tokenCache) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, cached := range c.cache {
		if now.After(cached.expiresAt) {
			delete(c.cache, key)
		}
	}
}

// oauthTokenURLs maps OVH API endpoint URLs to their OAuth token endpoints.
// Mirrors the unexported tokensURLs from github.com/ovh/go-ovh/ovh.
var oauthTokenURLs = map[string]string{
	ovh.OvhEU: "https://www.ovh.com/auth/oauth2/token",
	ovh.OvhCA: "https://ca.ovh.com/auth/oauth2/token",
	ovh.OvhUS: "https://us.ovhcloud.com/auth/oauth2/token",
}

// resolveEndpoint converts an endpoint name (e.g. "ovh-eu") to the full URL,
// or returns it as-is if it already looks like a URL.
func resolveEndpoint(endpoint string) string {
	if strings.Contains(endpoint, "/") {
		return endpoint
	}
	if url, ok := ovh.Endpoints[endpoint]; ok {
		return url
	}
	return endpoint
}

// resolveTokenURL returns the OAuth token URL for the given API endpoint.
func resolveTokenURL(endpoint string) (string, error) {
	apiURL := resolveEndpoint(endpoint)
	if tokenURL, ok := oauthTokenURLs[apiURL]; ok {
		return tokenURL, nil
	}
	return "", fmt.Errorf("no OAuth token URL known for endpoint %q", endpoint)
}

// getOrExchangeToken retrieves a cached token or performs OAuth exchange
// using golang.org/x/oauth2/clientcredentials (the same library the go-ovh SDK uses).
func getOrExchangeToken(ctx context.Context, cacheKey, clientID, clientSecret, endpoint string) (string, error) {
	// Try to get from cache first
	if token, found := globalTokenCache.get(cacheKey); found {
		return token, nil
	}

	// Resolve the OAuth token endpoint URL
	tokenURL, err := resolveTokenURL(endpoint)
	if err != nil {
		return "", err
	}

	// Use the same clientcredentials flow as the go-ovh SDK
	conf := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
		Scopes:       []string{"all"},
	}

	// This actually performs the HTTP POST to get the access token
	token, err := conf.TokenSource(ctx).Token()
	if err != nil {
		return "", fmt.Errorf("OAuth token exchange failed: %w", err)
	}

	if token.AccessToken == "" {
		return "", fmt.Errorf("OAuth exchange succeeded but access token is empty")
	}

	// Cache the token; use the server-reported expiry if available
	ttl := tokenCacheTTL
	if !token.Expiry.IsZero() {
		serverTTL := time.Until(token.Expiry) - 5*time.Minute // 5-min safety margin
		if serverTTL > 0 && serverTTL < ttl {
			ttl = serverTTL
		}
	}
	globalTokenCache.set(cacheKey, token.AccessToken, ttl)

	return token.AccessToken, nil
}

func init() {
	// Start a background goroutine to periodically clean up expired cache entries
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			globalTokenCache.cleanup()
		}
	}()
}
