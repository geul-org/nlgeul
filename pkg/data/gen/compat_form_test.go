//ff:func feature=data type=helper control=sequence
//ff:what TestCompatForm — brackets, tortoise-shell brackets, a dot or a comma after digits, the Number Forms and jamo blocks, and the rest each give their form
package main

import "testing"

func TestCompatForm(t *testing.T) {
	form, v := compatForm([]rune("(1)"), "")
	assertEq(t, form+string(v), "괄호1")
	form, v = compatForm([]rune("〔本〕"), "")
	assertEq(t, form+string(v), "거북등 괄호本")
	form, v = compatForm([]rune("10."), "")
	assertEq(t, form+string(v), "점 붙음10")
	form, v = compatForm([]rune("10,"), "")
	assertEq(t, form+string(v), "쉼표 붙음10")
	form, _ = compatForm([]rune("a."), "")
	assertEq(t, form, "그 밖의 호환")
	form, _ = compatForm([]rune("I"), "Number Forms")
	assertEq(t, form, "로마 숫자")
	form, _ = compatForm([]rune("ᄀ"), "Hangul Compatibility Jamo")
	assertEq(t, form, "호환 자모")
	form, _ = compatForm([]rune("x"), "Other")
	assertEq(t, form, "그 밖의 호환")
}
