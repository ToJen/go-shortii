package tests

import (
	"shortii/shorten"
	"testing"
)

func TestURL(t *testing.T) {
	url := "https://www.titan-suite.com"
	shortURL, err := shorten.Shorten(url)
	if err != nil {
		t.Error(err)
	}
	t.Log(shortURL)
}
