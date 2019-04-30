package main

import (
	"bufio"
	"fmt"
	"os"
	"shortii/shorten"
	"strings"
)

func main() {
	const NAME = "shortii"
	const VERSION = "0.1.1"
	fmt.Printf("Welcome to %s v%s, the easiest URL shortener out there\n\n", NAME, VERSION)

	// read user input
	fmt.Println("Enter a URL:\t")

	reader := bufio.NewReader(os.Stdin)
	urlInput, _ := reader.ReadString('\n')
	urlInput = strings.TrimSuffix(urlInput, "\n") // get rid of newline character

	shortURL, err := shorten.Shorten(urlInput)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Here you go:\t%s\n", shortURL)
	}

}
