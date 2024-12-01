package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexellis/hmac"
)

func main() {
	var inputVar string
	var secretVar string
	flag.StringVar(&inputVar, "message", "", "message to create a digest from")
	flag.StringVar(&secretVar, "secret", "", "secret for the digest")
	flag.Parse()

	if len(strings.TrimSpace(secretVar)) == 0 {
		panic("--secret is required")
	}
	fmt.Printf("Computing hash for: %q\nSecret: %q\n", inputVar, secretVar)
	digest := hmac.Sign([]byte(inputVar), []byte(secretVar))
	fmt.Printf("Digest: %x\n", digest)

	fmt.Printf("Computing hash for: %q\nSecret: %q\n", inputVar, secretVar)
}
