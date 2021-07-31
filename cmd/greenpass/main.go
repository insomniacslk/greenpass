package main

import (
	"fmt"
	"log"
	"os"

	"github.com/insomniacslk/greenpass"
)

func main() {
	// open and decode image file
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file name>\n", os.Args[0])
		os.Exit(1)
	}
	fname := os.Args[1]
	fd, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Failed to open %s: %v", fname, err)
	}
	if err := greenpass.Read(fd); err != nil {
		log.Fatal(err)
	}
}
