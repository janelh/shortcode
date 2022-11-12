package shortcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateShortcode(t *testing.T) {
	expectedResult := "39xyOmxz"
	url := "https://github.com/rameshsunkara/go-rest-api-example/tree/main/internal/db"
	result := GenerateShortcode(url)
	assert.EqualValues(t, result, expectedResult, "Shortcode generated does not match expected result")
}
