//ff:func feature=data type=helper control=sequence
//ff:what TestComposeHangul — jamo in a row become syllables and other text is left as it is
package main

import "testing"

func TestComposeHangul(t *testing.T) {
	assertEq(t, composeHangul("(각)x"), "(각)x")
	assertEq(t, composeHangul(""), "")
}
