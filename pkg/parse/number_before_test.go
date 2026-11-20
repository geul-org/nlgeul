//ff:func feature=parse type=helper control=sequence
//ff:what TestNumberBefore — a digit, or a word of the unit table right after a digit, ends a number before a place (36조│제, 3개월│제, 의2│제); a syllable, a unit with no digit before it, a unit at the very head and the head itself do not
package parse

import "testing"

func TestNumberBefore(t *testing.T) {
	if sc := scanOf("36조제1항"); !sc.numberBefore(3) || !sc.numberBefore(1) || sc.numberBefore(0) {
		t.Fatal("36조제1항")
	}
	if !scanOf("조의2제3").numberBefore(3) {
		t.Fatal("after a digit")
	}
	if !scanOf("3개월제1").numberBefore(3) {
		t.Fatal("after a unit of two letters")
	}
	if scanOf("문제1").numberBefore(1) {
		t.Fatal("after a syllable")
	}
	if scanOf("가조제1").numberBefore(2) {
		t.Fatal("a unit with no digit before it")
	}
	if scanOf("조제1").numberBefore(1) {
		t.Fatal("a unit at the head")
	}
}
