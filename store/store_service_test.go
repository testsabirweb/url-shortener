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
	originalUrl, _ := RetrieveInitialUrl(shortURL)
	assert.Equal(t, initialLink, originalUrl)
}
