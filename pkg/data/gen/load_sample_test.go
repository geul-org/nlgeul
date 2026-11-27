//ff:func feature=data type=helper control=sequence
//ff:what loadSample — the table made from the sample UnicodeData.txt lines
package main

import (
	"strings"
	"testing"
)

func loadSample(t *testing.T) *ucdTable {
	tbl, err := parseUnicodeData(strings.NewReader(sampleUCD))
	if err != nil {
		t.Fatal(err)
	}
	return tbl
}
