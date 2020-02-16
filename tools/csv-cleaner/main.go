package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 2 {
		printUsage()
		os.Exit(1)
	}

	var r io.Reader = os.Stdin
	var w io.Writer = os.Stdout

	if len(args) == 2 {
		fileName := args[1]
		file, err := openFile(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		r = file
	}

	err := clean(r, w)
	if err != nil {
		log.Fatal(err)
	}

}

func openFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func printUsage() {
	msg := `csv cleaner.

Usage:
    csv-cleaner [file]`

	fmt.Println(msg)
}
