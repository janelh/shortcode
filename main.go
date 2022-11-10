package main

import (
	"example.com/shortcode"
	"example.com/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type url struct {
	Url string `json:"url"`
}

func postUrl(context *gin.Context) {
	var newUrl url

	if err := context.BindJSON(&newUrl); err != nil {
		return
	}

	shortcode := shortcode.GenerateShortcode(newUrl.Url)

	store.CreateUrl(shortcode, newUrl.Url)
	hostUrl := "localhost:9001"
	context.IndentedJSON(http.StatusCreated, hostUrl+shortcode)
}

func getUrl(context *gin.Context) {
	shortCode := context.Param("shortCode")
	redirectUrl := store.RetrieveUrl(shortCode)
	if redirectUrl != "" {
		context.Redirect(301, redirectUrl)
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "url not found"})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot load .env file")
	}

	rdsPwd := os.Getenv("RDS_PASSWORD")
	rdsHost := os.Getenv("RDS_HOST")

	store.InitStoreClient(rdsHost, rdsPwd)

	router := gin.Default()
	router.POST("urls", postUrl)
	router.GET("urls/:shortCode", getUrl)

	router.Run("localhost:9001")
}
