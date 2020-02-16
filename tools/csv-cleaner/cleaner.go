package main

import (
	"encoding/csv"
	"io"
	"strings"
)

func clean(r io.Reader, w io.Writer) error {
	csvr := csv.NewReader(r)

	for {
		rec, err := csvr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		s := strings.Join(rec, ",") + "\n"
		w.Write([]byte(s))
	}

	return nil
}
