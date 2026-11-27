//ff:func feature=data type=helper control=sequence
//ff:what TestCompatJamo — a lone initial becomes a compatibility jamo, an initial and a vowel become a syllable, other characters stay
package main

import "testing"

func TestCompatJamo(t *testing.T) {
	jamo := jamoMap(loadSample(t))
	assertEq(t, compatJamo(jamo, "ᄀ"), "ㄱ")
	assertEq(t, compatJamo(jamo, "가"), "가")
	assertEq(t, compatJamo(jamo, "A"), "A")
}
