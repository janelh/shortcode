package store

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	testStore = &Store{}
)

func init() {
	rdsPwd := os.Getenv("RDS_PASSWORD")
	rdsHost := os.Getenv("RDS_HOST")
	store := InitStoreClient(rdsHost, rdsPwd)
	testStore = store
}

func TestInitStoreClient(t *testing.T) {
	assert.True(t, testStore != nil)
}

func TestCreateAndRetrieve(t *testing.T) {
	url := "https://go.dev/doc/tutorial/add"
	shortcode := "s9g0Pfs5"
	CreateUrl(shortcode, url)
	result := RetrieveUrl(shortcode)
	assert.EqualValues(t, result, url, "URL and short retrieval and/or creation failed")
}
