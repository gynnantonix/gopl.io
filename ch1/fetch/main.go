// Fetch prints the content found at each specified URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") {
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			panic(err)
		}
		count, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d bytes written\n", count)
	}
}
