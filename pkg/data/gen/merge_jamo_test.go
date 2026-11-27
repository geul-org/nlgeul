//ff:func feature=data type=helper control=sequence
//ff:what TestMergeJamo — an initial and a vowel join into a syllable, a final joins a syllable with none, other characters are appended
package main

import "testing"

func TestMergeJamo(t *testing.T) {
	out := mergeJamo(nil, 0x1100)
	assertEq(t, string(out), "ᄀ")
	out = mergeJamo(out, 0x1161)
	assertEq(t, string(out), "가")
	out = mergeJamo(out, 0x11A8)
	assertEq(t, string(out), "각")
	out = mergeJamo(out, 0x11A8)
	assertEq(t, string(out), "각ᆨ")
	assertEq(t, string(mergeJamo([]rune{'a'}, 0x1161)), "aᅡ")
}
