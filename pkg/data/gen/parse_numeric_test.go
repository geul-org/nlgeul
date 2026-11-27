//ff:func feature=data type=helper control=iteration dimension=1
//ff:what TestParseNumeric — integers, fractions and negative fractions are read; an empty, bad or zero-denominator column has no number
package main

import "testing"

func TestParseNumeric(t *testing.T) {
	if n, ok := parseNumeric("1/2"); !ok || n != 0.5 {
		t.Fatal(n, ok)
	}
	if n, ok := parseNumeric("-1/2"); !ok || n != -0.5 {
		t.Fatal(n, ok)
	}
	if n, ok := parseNumeric("100"); !ok || n != 100 {
		t.Fatal(n, ok)
	}
	for _, s := range []string{"", "x", "1/x", "1/0"} {
		if _, ok := parseNumeric(s); ok {
			t.Fatal(s)
		}
	}
}
