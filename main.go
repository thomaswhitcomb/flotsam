package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/thomaswhitcomb/jetsam"
)

func loadURLs(in io.Reader) []string {
	urls := []string{}
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	return urls
}

func main() {
	myLoader := func(url string) (io.Reader, error) {
		if strings.HasPrefix(url, "file:") {
			name := url[5:]
			fmt.Println("name", name)
			file, err := os.Open(name)
			if err != nil {
				return nil, err
			}
			return bufio.NewReader(file), nil
		}
		return nil, errors.New("Source must have file scheme")
	}
	urls := loadURLs(os.Stdin)
	ppl := jetsam.Pipeline{
		Loader:         myLoader,
		ItemDepth:      100,
		ProcessorCount: 3,
		Sources:        urls,
	}
	ppl.Provision()
	ppl.Run()
}
