package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func main() {

	// ─── PATHS ──────────────────────────────────────────────────────────────────────

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var (
		pathToParts = pwd + "/parts.txt"
		path        = pwd + "/something.png"
	)

	// ─── FILE THAT WERE GOING TO WRITE IN ───────────────────────────────────────────

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// ─── PATH OF WHERE THE PARTS ARE ────────────────────────────────────────────────

	caseFile, err := os.Open(pathToParts)
	if err != nil {
		log.Fatal(err)
	}
	defer caseFile.Close()

	// ─── READING EACH PART AND THEN WRITING IT INTO THE RESULT FILE ─────────────────

	scanner := bufio.NewScanner(caseFile)
	for scanner.Scan() {
		partBase64 := scanner.Text()

		dec, err := base64.StdEncoding.DecodeString(partBase64)
		if err != nil {
			panic(err)
		}

		if _, err := f.Write(dec); err != nil {
			panic(err)
		}
	}

}
