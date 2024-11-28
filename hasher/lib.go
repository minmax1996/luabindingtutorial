package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
)

func doubleMd5Sorted(input string) string {
	// Parse the input URL
	urlParsed, err := url.Parse(input)
	if err != nil {
		return "error" + err.Error()
	}

	// Sort the query parameters
	queryParams := urlParsed.Query()
	keys := make([]string, 0, len(queryParams))
	for key := range queryParams {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Create the sorted query string
	sortedParams := []string{}
	for _, key := range keys {
		sortedParams = append(sortedParams, key+"="+queryParams.Get(key))
	}
	sortedParamsStr := strings.Join(sortedParams, "&")
	// Hash the sorted query string twice
	firstHash := md5.Sum([]byte(sortedParamsStr))
	secondHash := md5.Sum(firstHash[:])
	return hex.EncodeToString(secondHash[:])
}

func sha256Raw(input string) string {
	// Parse the input URL
	urlParsed, err := url.Parse(input)
	if err != nil {
		return "error" + err.Error()
	}

	// Hash the raw query string
	hash := sha256.New()
	hash.Write([]byte(urlParsed.RawQuery))
	return hex.EncodeToString(hash.Sum(nil))
}
