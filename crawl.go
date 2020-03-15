package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Please provide the starting URL")
		os.Exit(1)
	} else {
		fmt.Printf("Passed URL: %v \n", args)
		retrieve(args[0])
	}

}

func retrieve(url string) {
	resp, err := http.Get(url)
	if err != nil {
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(body))
}
