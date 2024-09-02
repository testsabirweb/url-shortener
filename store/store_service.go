package store

import (
	"errors"
	"fmt"
	"sync"
)

var (
	storeService = &StorageService{
		urlMap: make(map[string]string),
	}
	// ctx = context.Background()
)

type StorageService struct {
	urlMap map[string]string
	mutex  sync.Mutex
}

// InitializeStore initializes the in-memory store.
// It simulates the initialization of a storage service like Redis but for an in-memory map.
func InitializeStore() *StorageService {
	fmt.Println("In-memory store initialized successfully")
	return storeService
}

// SaveUrlMapping saves the mapping between the original URL and the generated short URL.
func SaveUrlMapping(shortUrl string, originalUrl string) {
	storeService.mutex.Lock()
	defer storeService.mutex.Unlock()

	// Save the mapping in the in-memory map
	storeService.urlMap[shortUrl] = originalUrl

	fmt.Printf("Saved shortUrl: %s - originalUrl: %s\n", shortUrl, originalUrl)
}

// RetrieveInitialUrl retrieves the original URL associated with the given short URL.
func RetrieveInitialUrl(shortUrl string) string {
	storeService.mutex.Lock()
	defer storeService.mutex.Unlock()

	// Retrieve the original URL from the in-memory map
	originalUrl, exists := storeService.urlMap[shortUrl]
	if !exists {
		err := errors.New("short URL not found")
		panic(fmt.Sprintf("Failed to RetrieveInitialUrl | Error: %v - shortUrl: %s\n", err, shortUrl))
	}

	return originalUrl
}
