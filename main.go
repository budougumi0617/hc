package main

import "net/url"

// Result shows executed result.
type Result struct {
	// Page is target URL.
	Page url.URL
	// HBC is Hatena Bookmark Count.
	HBC int
	// Err is error if API call failed.
	Err error
}

func main() {
}
