package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

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
	/*
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
	*/
	n := flag.Int(
		"n",
		3,
		"number of goroutines processing CSV files")
	flag.Parse()

	urls := loadURLs(os.Stdin)
	ppl := jetsam.Pipeline{
		BufferSize:     100,
		ProcessorCount: *n,
		Sources:        urls,
		Reducer:        reduce,
	}
	ppl.Provision()
	result := ppl.Run()
	fmt.Printf(
		"Average: %s\n",
		result.Results["average"])

	fmt.Printf(
		"Median: %s\n",
		result.Results["median"])
	fmt.Printf(
		"Number of people with median age: %s\n",
		result.Results["count"])
	fmt.Printf(
		"Names: %s\n",
		result.Results["names"])
}
