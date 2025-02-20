package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"os"
)

func main() {
	// The base64 encoded string
	encodedStr := "H4sIACfXsmcAA1NQgIBH0zqACJWNLEYIcMEYCE3YWHQyBskg7N4hw0WUu4kWgAvZaeQhEi1BlyTVQCyG+bn7B3iiiJJpkEuoTyiqKJkGUew9bMaRmOgoiVNS4hfJMmzSJBhC83QIAGGVYr6uBAAA"

	// Decode the base64 string
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		panic(err)
	}

	// Create a gzip reader
	gzipReader, err := gzip.NewReader(bytes.NewReader(decodedBytes))
	if err != nil {
		panic(err)
	}
	defer gzipReader.Close()

	// Copy the decompressed data to stdout
	if _, err := io.Copy(os.Stdout, gzipReader); err != nil {
		panic(err)
	}
}