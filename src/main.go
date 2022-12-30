package main

import (
	"github.com/gin-gonic/gin"
	"github.com/janelh/shortcode/src/shortcode"
	"github.com/janelh/shortcode/src/store"
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

	err := store.CreateUrl(shortcode, newUrl.Url)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Url not stored"})
	}

	baseUrl := os.Getenv("BASE_URL")
	context.IndentedJSON(http.StatusCreated, baseUrl+shortcode)
}

func getUrl(context *gin.Context) {
	shortCode := context.Param("shortCode")
	redirectUrl, err := store.RetrieveUrl(shortCode)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "url not found"})
	}

	context.Redirect(301, redirectUrl)
}

func checkEnvVars(vars []string) {
	for _, v := range vars {
		if _, exists := os.LookupEnv(v); !exists {
			log.Fatalf("%s environment variable not found", v)
		}
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot load .env file")
	}

	checkEnvVars([]string{"RDS_PWD", "RDS_HOST", "BASE_URL"})

	rdsPwd := os.Getenv("RDS_PWD")
	rdsHost := os.Getenv("RDS_HOST")

	store.InitStoreClient(rdsHost, rdsPwd)

	router := gin.Default()
	router.POST("urls", postUrl)
	router.GET("urls/:shortCode", getUrl)

	err = router.Run()
	if err != nil {
		log.Fatalf("Failed to start Gin http server - %s", err)
	}
}
