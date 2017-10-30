package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var algo = flag.String("algo", "sha256", "hash algorithm")

func main() {
	src := []byte("a")

	flag.Parse()

	switch *algo {
	case "sha256":
		fmt.Println(sha256.Sum256(src))
	case "sha384":
		fmt.Println(sha512.Sum384(src))
	case "sha512":
		fmt.Println(sha512.Sum512(src))
	default:
		fmt.Fprintf(os.Stderr, "Unsupported algorithm: %s\n", *algo)
		os.Exit(-1)
	}

}
