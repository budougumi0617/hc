package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"strconv"
)

// Entry shows executed result.
type Entry struct {
	// Page is target URL.
	Page url.URL
	// HBC is Hatena Bookmark Count.
	HBC int
	// Err is error if API call failed.
	Err error
}

func readLines(in io.Reader) []string {
	var ss []string
	s := bufio.NewScanner(in)
	for s.Scan() {
		ss = append(ss, strconv.Quote(s.Text()))
	}
	if s.Err() != nil {
		log.Fatal(s.Err())
	}
	return ss
}

func main() {
	/**
	  $ curl -D - -X GET http://api.b.st-hatena.com/entry.count?url=https%3A%2F%2Fbudougumi0617.github.io%2F2019%2F05%2F12%2Fpass-aws-solution-architect-associate%2F
	  HTTP/1.1 200 OK
	  Content-Type: text/plain
	  Content-Length: 3
	  Connection: keep-alive
	  Date: Mon, 24 Jun 2019 12:21:51 GMT
	  Server: nginx
	  Cache-Control: public, max-age=3600, s-maxage=3600
	  X-Cache: Miss from cloudfront
	  Via: 1.1 4ca8d239c2b4b1a578fa3c7797e67c11.cloudfront.net (CloudFront)
	  X-Amz-Cf-Pop: NRT57-C3
	  X-Amz-Cf-Id: 3wS1whM3YI4I_PWIriHF6jGjZ5YkVXpGVAMbUSFarfz8qeUnI6osTw==

	  268%
	*/
	ss := readLines(os.Stdin)
	for _, s := range ss {
		fmt.Println(s)
	}
	// TODO build request
	// TODO Execute GET
	// TODO Parse response
	// TODO show results
	// TODO Repeat for input URLs
}
