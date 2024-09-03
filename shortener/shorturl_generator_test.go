package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortLinkGenerator(t *testing.T) {
	userId := "user1"
	initialLink_1 := "https://www.infracloud.io/"
	shortLink_1 := GenerateShortLink(initialLink_1, userId)

	initialLink_2 := "https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/"
	shortLink_2 := GenerateShortLink(initialLink_2, userId)

	initialLink_3 := "https://dev.to/justlorain/go-how-to-write-a-worker-pool-1h3b"
	shortLink_3 := GenerateShortLink(initialLink_3, userId)

	assert.Equal(t, shortLink_1, "WtCEaFgL")
	assert.Equal(t, shortLink_2, "ByJG4ZeU")
	assert.Equal(t, shortLink_3, "b7B9TQd6")
}
