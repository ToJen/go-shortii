package shorten

import (
	"testing"
)

func TestURL(t *testing.T) {
	url := "https://www.titan-suite.com"
	shortURL, err := Shorten(url)
	if err != nil {
		t.Error(err)
	}
	t.Log(shortURL)
}
