package main

import (
	"flag"
)

var (
	algorithmFlag = flag.String("algorithm", "", "The hashing algorithm to use")
	inputFlag     = flag.String("input", "", "The input to hash")
)

func init() {
	flag.Parse()
}

func main() {
	switch *algorithmFlag {
	case "double_md5_sorted":
		hash := doubleMd5Sorted(*inputFlag)
		println(hash)
	case "sha256_raw":
		hash := sha256Raw(*inputFlag)
		println(hash)
	default:
		println("Unknown algorithm")
	}
}
