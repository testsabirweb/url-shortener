package store

import (
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"sync"
)

var (
	storeService = &StorageService{
		urlMap:      make(map[string]string),
		domainCount: make(map[string]int),
	}
)

type StorageService struct {
	urlMap      map[string]string
	domainCount map[string]int
	mutex       sync.Mutex
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

	// Extract domain name from original URL and update domain count
	domain := extractDomain(originalUrl)
	storeService.domainCount[domain]++

	fmt.Printf("Saved shortUrl: %s - originalUrl: %s\n", shortUrl, originalUrl)
}

// extractDomain extracts the domain name from the given URL.
func extractDomain(originalUrl string) string {
	parsedUrl, err := url.Parse(originalUrl)
	if err != nil {
		return ""
	}
	domain := parsedUrl.Hostname()
	// Remove "www." if present
	return strings.TrimPrefix(domain, "www.")
}

// RetrieveInitialUrl retrieves the original URL associated with the given short URL.
func RetrieveInitialUrl(shortUrl string) (string, error) {
	storeService.mutex.Lock()
	defer storeService.mutex.Unlock()

	// Retrieve the original URL from the in-memory map
	originalUrl, exists := storeService.urlMap[shortUrl]
	if !exists {
		err := errors.New("short URL not found")
		return "", err
	}

	return originalUrl, nil
}

// GetTopDomains returns the top 3 domain names shortened the most.
func GetTopDomains() map[string]int {
	storeService.mutex.Lock()
	defer storeService.mutex.Unlock()

	// Create a slice to store the domain names and their counts
	type domainCountPair struct {
		domain string
		count  int
	}
	domainCounts := make([]domainCountPair, 0, len(storeService.domainCount))

	// Populate the slice
	for domain, count := range storeService.domainCount {
		domainCounts = append(domainCounts, domainCountPair{domain: domain, count: count})
	}

	// Sort the slice by count in descending order
	sort.Slice(domainCounts, func(i, j int) bool {
		return domainCounts[i].count > domainCounts[j].count
	})

	// Get the top 3 domains
	topDomains := make(map[string]int)
	for i := 0; i < len(domainCounts) && i < 3; i++ {
		topDomains[domainCounts[i].domain] = domainCounts[i].count
	}

	return topDomains
}
