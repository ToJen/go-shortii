package shorten

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

// validateURL checks urlInput against a RegularExpression
func validateURL(urlInput string) (matched bool) {
	var validURL = regexp.MustCompile(`^(http(s)?\:\/\/)?(www\.)?[\w-]+\.[a-z]+$`)
	matched = validURL.MatchString(urlInput)
	return
}

// Shorten makes a POST request to TinyURL's API using
// urlInput and returns the shortened url
func Shorten(urlInput string) (string, error) {
	// validate input
	valid := validateURL(urlInput)
	if !valid {
		return "", errors.New("invalid URL")
	}

	tinyURLAPI := "https://tinyurl.com/api-create.php/"
	response, err := http.PostForm(tinyURLAPI, url.Values{
		"url": {urlInput},
	})
	if err != nil {
		fmt.Printf("There's something wrong with %s: \n%v", urlInput, err)
	}

	defer response.Body.Close() // close when program is complete

	shortURL, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("There's something wrong with the API response: \n%v", err)
	}

	return string(shortURL), nil
}
