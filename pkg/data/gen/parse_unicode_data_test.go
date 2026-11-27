//ff:func feature=data type=helper control=sequence
//ff:what TestParseUnicodeData — characters are read by code point, a First/Last pair makes a range, a short line or a bad code is an error
package main

import (
	"strings"
	"testing"
)

func TestParseUnicodeData(t *testing.T) {
	tbl := loadSample(t)
	assertEq(t, tbl.Chars[0x2460].Decomp, "<circle> 0031")
	assertEq(t, tbl.Chars[0x00BD].Numeric, "1/2")
	assertEq(t, tbl.Chars[0x0009].Name, "")
	if len(tbl.Ranges) != 1 || tbl.Ranges[0].Lo != 0x4E00 || tbl.Ranges[0].Hi != 0x9FFF {
		t.Fatal(tbl.Ranges)
	}
	if _, err := parseUnicodeData(strings.NewReader("0041;A\n")); err == nil {
		t.Fatal("short line")
	}
	if _, err := parseUnicodeData(strings.NewReader("\n")); err != nil {
		t.Fatal(err)
	}
}
