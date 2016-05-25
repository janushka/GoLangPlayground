/*
// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"strings"
)
func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
			fmt.Println(url)
		} else {
			fmt.Println("Prefix exists!")
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		result, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "copy: copying %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\t%s", resp.Status, result)
	}
}*/
