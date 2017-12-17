package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "ERROR: No URL to encode received. Try ue -help \r\n")
		return
	} else if len(os.Args) > 2 {
		fmt.Fprintf(os.Stderr, "ERROR: Multiple arguments received. Try ue -help \r\n")
		for i, a := range os.Args[1:] {
			fmt.Fprintf(os.Stderr, "%d %s\r\n", i+1, a)
		}
		return
	}

	userUrl := os.Args[1]
	if userUrl == "-help" || userUrl == "--help" {
		displaySampleOfUsage()
		return
	}

	// Split the URL in its parts...
	u, err := url.Parse(userUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\r\n", err)
		return
	}

	// ...and rebuild it but forcing the query string parameters
	// to be encoded
	queryString := u.Query().Encode()
	if queryString != "" {
		queryString = "?" + queryString
	}
	u.RawQuery = ""
	encodedUrl := u.String() + queryString
	fmt.Print(encodedUrl)
}

func displaySampleOfUsage() {
	simpleUrl := "http://localhost/something?q=hello world"
	complexUrl := "http://localhost:8983/solr/bibdata/select?debugQuery=false&q=title:\\\"silver buckle\\\""
	fmt.Printf("URL encodes the value passed.\r\n")
	fmt.Printf("Samples of usage: \r\n\r")
	fmt.Printf("\r\n")
	fmt.Printf("ue \"%s\"\r\n", simpleUrl)
	fmt.Printf("ue \"%s\"\r\n", complexUrl)
	fmt.Printf("curl $(ue \"%s\")\r\n", complexUrl)
	fmt.Printf("\r\n")
}
