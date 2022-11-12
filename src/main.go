package main

import (
	"example.com/src/shortcode"
	"example.com/src/store"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

	err := store.CreateUrl(shortcode, newUrl.Url)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Url not stored"})
	}

	hostUrl := "localhost:8080"
	context.IndentedJSON(http.StatusCreated, hostUrl+shortcode)
}

func getUrl(context *gin.Context) {
	shortCode := context.Param("shortCode")
	redirectUrl, err := store.RetrieveUrl(shortCode)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "url not found"})
	}

	context.Redirect(301, redirectUrl)
}

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Cannot load .env file")
	//}
	//
	//rdsPwd := os.Getenv("RDS_PASSWORD")
	//rdsHost := os.Getenv("RDS_HOST")
	//
	//store.InitStoreClient(rdsHost, rdsPwd)

	router := gin.Default()
	router.POST("urls", postUrl)
	router.GET("urls/:shortCode", getUrl)

	err := router.Run()
	if err != nil {
		log.Fatalf("Failed to start Gin http server - %s", err)
	}
}
