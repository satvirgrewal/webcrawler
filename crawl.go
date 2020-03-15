package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/jackdanger/collectlinks"
)

var visited = make(map[string]bool)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Please provide the starting URL")
		os.Exit(1)
	} else {
		fmt.Printf("Passed URL: %v \n", args)
	}
	queue := make(chan string)
	go func() {

		queue <- args[0]
	}()
	for url := range queue {
		enqueue(url, queue)
	}
}

func enqueue(url string, queue chan string) {
	fmt.Printf("Processing link: %v \n", url)
	visited[url] = true
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	client := http.Client{Transport: transport}

	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	links := collectlinks.All(resp.Body)
	for _, link := range links {
		absolute := fixUrl(link, url)
		if absolute != "" {
			if !visited[absolute] {
				go func() {
					queue <- absolute
				}()
			}

		}
	}
}

func fixUrl(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}
