//ff:func feature=data type=helper control=sequence
//ff:what TestTagOf — a tagged decomposition gives its tag, a canonical one or none gives nothing
package main

import "testing"

func TestTagOf(t *testing.T) {
	assertEq(t, tagOf("<circle> 0031"), "<circle>")
	assertEq(t, tagOf("0041 0301"), "")
	assertEq(t, tagOf(""), "")
}
