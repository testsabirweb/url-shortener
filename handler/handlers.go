package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/testsabirweb/url-shortener/shortener"
	"github.com/testsabirweb/url-shortener/store"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

// ResolveCollision handles potential collisions by generating a new short URL if needed
func ResolveCollision(originalUrl, shortUrl string, attempt int) (string, error) {
	if attempt > 5 { // Limit recursion to prevent stack overflow
		return "", fmt.Errorf("too many collisions, cannot generate unique short URL")
	}

	initialUrl, err := store.RetrieveInitialUrl(shortUrl)
	if err.Error() == "short URL not found" { // This indicates that the short URL is unique
		return shortUrl, nil
	}

	if initialUrl == originalUrl {
		return shortUrl, nil
	}

	// Generate a new short URL to avoid collision
	newShortUrl := shortener.GenerateShortLink(originalUrl + shortUrl + fmt.Sprint(attempt))
	return ResolveCollision(originalUrl, newShortUrl, attempt+1)
}

// CreateShortUrl handles the creation of a new short URL
func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	host := "http://localhost:3000/"

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl)
	initialUrl, err := store.RetrieveInitialUrl(shortUrl)
	if err != nil && err.Error() != "short URL not found" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	if err != nil && err.Error() == "short URL not found" {
		store.SaveUrlMapping(shortUrl, creationRequest.LongUrl)
		c.JSON(http.StatusOK, gin.H{
			"message":   "short URL created successfully",
			"short_url": host + shortUrl,
		})
		return
	}

	if initialUrl == creationRequest.LongUrl {
		c.JSON(http.StatusOK, gin.H{
			"message":   "short URL already present",
			"short_url": host + shortUrl,
		})
		return
	}

	// Handle collision
	shortUrl, err = ResolveCollision(creationRequest.LongUrl, shortUrl, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl)
	c.JSON(200, gin.H{
		"message":   "short URL created successfully",
		"short_url": host + shortUrl,
	})
}

// HandleTopDomains returns the top 3 most frequent domains
func HandleTopDomains(c *gin.Context) {
	topDomains := store.GetTopDomains()
	c.JSON(200, gin.H{
		"message": "Top 3 Domains are",
		"domains": topDomains,
	})
}

// HandleShortUrlRedirect handles the redirection for a short URL
func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl, err := store.RetrieveInitialUrl(shortUrl)
	if err != nil {
		if err.Error() == "short URL not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "short URL not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve initial URL"})
		return
	}
	c.Redirect(http.StatusFound, initialUrl)
}
