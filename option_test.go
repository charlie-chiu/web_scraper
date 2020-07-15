package scraper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	o := &Options{
		Region:  8,
		Section: "100,102",
	}

	wantURL := "https://rent.591.com.tw/?firstRow=0&kind=0&order=&orderType=&region=8&section=100%2C102"
	gotURL, err := o.URL()

	assert.Nil(t, err)
	assert.Equal(t, gotURL, wantURL)
}
