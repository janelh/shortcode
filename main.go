package main

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

type url struct {
	ShortCode string `json:"shortCode"`
	Url       string `json:"url"`
}

var urls = []url{
	{ShortCode: "S9fkSIfj", Url: "https://stackoverflow.com/questions/6109225/echoing-the-last-command-run-in-bash"},
	{ShortCode: "lFiwp93K", Url: "https://www.pccasegear.com/category/113_1361/keyboards/ducky-keyboards"},
}

func getUrls(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, urls)
}

func postUrls(context *gin.Context) {
	var newUrl url

	if err := context.BindJSON(&newUrl); err != nil {
		return
	}
	// Hash url, url safe encode hash and store first 8 chars of encoded string
	h := sha256.New()
	h.Write([]byte(newUrl.Url))
	encodedString := base64.URLEncoding.EncodeToString(h.Sum(nil))
	newUrl.ShortCode = encodedString[:8]

	urls = append(urls, newUrl)
	context.IndentedJSON(http.StatusCreated, newUrl)
}

func getUrlByShortCode(context *gin.Context) {
	id := context.Param("shortCode")
	for _, url := range urls {
		if url.ShortCode == id {
			context.IndentedJSON(http.StatusOK, url)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "url not found"})
}

func main() {
	router := gin.Default()
	router.GET("urls", getUrls)
	router.POST("urls", postUrls)
	router.GET("urls/:shortCode", getUrlByShortCode)

	router.Run("localhost:9000")
}
