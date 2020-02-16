package main

import (
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
	w := os.Stdout

	if len(args) == 2 {
		fileName := args[1]
		file, err := os.Open(fileName)
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}

		r = file
	}

	err := clean(r, w)
	if err != nil {
		log.Fatal(err)
	}

}

func printUsage() {
	msg := `csv cleaner.

Usage:
    csv-cleaner [file]`

	fmt.Println(msg)
}
