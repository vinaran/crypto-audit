package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var update = flag.Bool("update", false, "update golden file")

func TestCleaner(t *testing.T) {
	flag.Parse()

	tests := []struct {
		name     string
		fileName string
	}{
		{
			name:     "basic csv file",
			fileName: "./testdata/statement-20131231.csv",
		},
	}

	tmpExt := ".golden.tmp"
	gldExt := ".golden"
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			in, err := openFile(tc.fileName)
			if err != nil {
				log.Fatal(err)
			}

			tmp := tc.fileName + tmpExt
			out, err := os.Create(tmp)
			if err != nil {
				t.Fatal(err)
			}

			err = clean(in, out)
			if err != nil {
				t.Fatal(err)
			}

			gld := tc.fileName + gldExt
			if *update {
				fmt.Printf("updating golden test file: %s\n", gld)

				b, err := openAndReadAll(tmp)
				if err != nil {
					t.Fatal(err)
				}

				err = ioutil.WriteFile(gld, b, 0644)
				if err != nil {
					t.Fatal(err)
				}

				return
			}

			res, err := compareGoldenFiles(tmp, gld)
			if err != nil {
				t.Fatal(err)
			}

			if !res {
				t.Fatalf("Output did not match golden test file. Compare files using the diff command:\n`diff %s %s`", tmp, gld)
			}
		})
	}
}

func compareGoldenFiles(tmpFileName, gldFileName string) (bool, error) {
	tmpb, err := openAndReadAll(tmpFileName)
	if err != nil {
		return false, err
	}

	gldb, err := openAndReadAll(gldFileName)
	if err != nil {
		return false, err
	}

	return bytes.Equal(tmpb, gldb), nil
}

func openAndReadAll(fileName string) ([]byte, error) {
	file, err := openFile(fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
