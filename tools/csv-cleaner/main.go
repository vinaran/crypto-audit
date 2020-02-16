package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 2 {
		printUsage()
		os.Exit(1)
	}

	r := os.Stdin
	if len(args) == 2 {
		fileName := args[1]
		file, err := os.Open(fileName)
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}

		r = file
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	contents := buf.String()

	fmt.Print(contents)
}

func printUsage() {
	msg := `csv cleaner.

Usage:
    csv-cleaner [file]`

	fmt.Println(msg)
}
