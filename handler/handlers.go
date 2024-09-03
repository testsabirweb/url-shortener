package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/testsabirweb/url-shortener/shortener"
	"github.com/testsabirweb/url-shortener/store"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	host := "http://localhost:3000/"

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl)
	initialUrl, _ := store.RetrieveInitialUrl(shortUrl)
	if initialUrl == creationRequest.LongUrl {
		c.JSON(200, gin.H{
			"message":   "short url already present",
			"short_url": host + shortUrl,
		})
	} else {
		store.SaveUrlMapping(shortUrl, creationRequest.LongUrl)

		c.JSON(200, gin.H{
			"message":   "short url created successfully",
			"short_url": host + shortUrl,
		})
	}
}

func HandleTopDomains(c *gin.Context) {
	topDomains := store.GetTopDomains()
	c.JSON(200, gin.H{
		"message": "Top 3 Domains are",
		"doamins": topDomains,
	})

}
func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl, _ := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
