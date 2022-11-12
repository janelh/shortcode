package shortcode

import (
	"crypto/sha256"
	"encoding/base64"
)

func GenerateShortcode(url string) string {
	hash := sha256.New()
	hash.Write([]byte(url))
	encodedString := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	shortcode := encodedString[:8]

	return shortcode
}
