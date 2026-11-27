//ff:func feature=data type=helper control=sequence
//ff:what TestAllDigits — only a non-empty string of digits counts
package main

import "testing"

func TestAllDigits(t *testing.T) {
	if !allDigits("12") || allDigits("") || allDigits("1a") {
		t.Fatal("digits")
	}
}
