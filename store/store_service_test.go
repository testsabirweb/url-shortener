package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	_store := InitializeStore()
	testStoreService = _store
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.infracloud.io/"
	shortURL := "something"

	SaveUrlMapping(shortURL, initialLink)
	originalUrl, err := RetrieveInitialUrl(shortURL)

	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, initialLink, originalUrl, "Retrieved URL should match the initial link")
}

func TestExtractDomain(t *testing.T) {
	tests := []struct {
		originalUrl string
		expected    string
	}{
		{"https://www.infracloud.io", "infracloud.io"},
		{"https://infracloud.io", "infracloud.io"},
		{"http://www.example.com", "example.com"},
		{"http://example.com", "example.com"},
		{"ftp://www.ftpserver.com", "ftpserver.com"},
		{"ftp://ftpserver.com", "ftpserver.com"},
		{"invalid-url", ""},                // Invalid URL
		{"https://sub.domain.com", "sub.domain.com"}, // Subdomain
	}

	for _, test := range tests {
		domain := extractDomain(test.originalUrl)
		assert.Equal(t, test.expected, domain)
	}
}

func TestGetTopDomains(t *testing.T) {
	// Reset domainCount for a clean test environment
	storeService.domainCount = map[string]int{
		"infracloud.io":  5,
		"example.com":    10,
		"test.com":       3,
		"another.com":    7,
	}

	expectedTopDomains := map[string]int{
		"example.com":  10,
		"another.com":  7,
		"infracloud.io": 5,
	}

	topDomains := GetTopDomains()

	assert.Equal(t, expectedTopDomains, topDomains)
}