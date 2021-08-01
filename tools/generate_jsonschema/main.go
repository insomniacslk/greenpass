package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	flagOutputFile = flag.String("o", "DCC.combined-schema.json", "output file")
	flagVersion    = flag.String("V", "1.3.0", "Schema version")
)

func fetch(u string) ([]byte, error) {
	resp, err := http.Get(u)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	return body, nil
}

func writeFile(fileName string, data []byte) error {
	fd, err := os.Create(*flagOutputFile)
	if err != nil {
		return err
	}
	defer func() {
		if err := fd.Close(); err != nil {
			log.Printf("Failed to close file '%s': %v", fileName, err)
		}
	}()
	_, err = fd.Write(data)
	return err
}

func main() {
	flag.Parse()

	u := fmt.Sprintf("https://raw.githubusercontent.com/ehn-dcc-development/ehn-dcc-schema/release/%s/DCC.combined-schema.json", *flagVersion)
	jsonData, err := fetch(u)
	if err != nil {
		log.Fatalf("Failed to fetch JSON schema: %v", err)
	}
	if err := writeFile(*flagOutputFile, jsonData); err != nil {
		log.Fatalf("Failed to write to file '%s': %v", *flagOutputFile, err)
	}
	log.Printf("Written JSON schema to file '%s'", *flagOutputFile)
}
