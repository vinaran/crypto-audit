package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		printUsage()
		os.Exit(1)
	}

	fileName := args[1]
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	contents := buf.String()

	fmt.Print(contents)
}

func printUsage() {
	msg := `csv cleaner.

Usage:
    csv-cleaner file`

	fmt.Println(msg)
}
