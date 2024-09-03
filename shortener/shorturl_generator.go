package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

// GenerateShortLink takes an initial URL as input and returns an 8-character shortened version of the URL.
func GenerateShortLink(initialLink, userId string) string {
	// Generate a SHA-256 hash of the input URL
	urlHashBytes := sha256Of(initialLink + userId)

	// Convert the hash bytes to a big integer and then to a uint64 number
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()

	// Encode the number using Base58 encoding
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))

	// Return the first 8 characters of the encoded string as the shortened link
	return finalString[:8]
}

// sha256Of generates a SHA-256 hash for the given input string.
func sha256Of(input string) []byte {
	// Create a new SHA-256 hash algorithm instance
	algorithm := sha256.New()

	// Write the input data to the hash algorithm
	algorithm.Write([]byte(input))

	// Compute and return the hash as a byte slice
	return algorithm.Sum(nil)
}

// base58Encoded encodes the given byte slice using Base58 encoding (Bitcoin alphabet).
func base58Encoded(bytes []byte) string {
	// Define the Base58 encoding scheme (Bitcoin variant)
	encoding := base58.BitcoinEncoding

	// Encode the input bytes using the Base58 scheme
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		// Print the error and exit the program if encoding fails
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Return the encoded string
	return string(encoded)
}
