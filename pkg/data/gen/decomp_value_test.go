//ff:func feature=data type=helper control=sequence
//ff:what TestDecompValue — the characters after the tag are joined, and an empty decomposition gives an empty string
package main

import "testing"

func TestDecompValue(t *testing.T) {
	assertEq(t, decompValue("<compat> 0028 0031 0029"), "(1)")
	assertEq(t, decompValue("0041"), "A")
	assertEq(t, decompValue(""), "")
}
