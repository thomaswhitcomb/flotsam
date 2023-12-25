package main

import (
	"bytes"
	"testing"
)

var cannedURLs = []string{
	"This is URL 1",
	"This is URL 2",
	"This is URL 3"}

func TestOneURLLoaded(t *testing.T) {
	url := 0
	r := bytes.NewReader([]byte(cannedURLs[url]))
	urls := loadURLs(r)
	if len(urls) != 1 {
		t.Errorf("Bad length returned: %d\n", len(urls))
	}
	if urls[0] != cannedURLs[url] {
		t.Errorf("Incorrect URL: %s\n", urls[url])
	}
}
func TestManyURLLoaded(t *testing.T) {

	buffer := cannedURLs[0]
	for _, s := range cannedURLs[1:] {
		buffer = buffer + "\n" + s
	}

	r := bytes.NewReader([]byte(buffer))
	urls := loadURLs(r)
	if len(urls) != len(cannedURLs) {
		t.Errorf("Bad length returned: %d\n", len(urls))
	}

	for i, _ := range cannedURLs {
		if urls[i] != cannedURLs[i] {
			t.Errorf("Incorrect URL.  url[%d]: %s\n", i, urls[i])
			break
		}
	}
}
func TestNoURLLoaded(t *testing.T) {
	r := bytes.NewReader([]byte(""))
	urls := loadURLs(r)
	if len(urls) != 0 {
		t.Errorf("Bad length returned: %d\n", len(urls))
	}
}
