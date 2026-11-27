//ff:func feature=data type=helper control=sequence
//ff:what deriveTagged — the row of a character whose decomposition has a tag: the form from the tag, the value from the characters it decomposes to
package main

import "strings"

func deriveTagged(t *ucdTable, jamo map[rune]rune, r rune, block string) formRow {
	d := t.Chars[r].Decomp
	tag, v := tagOf(d), []rune(decompValue(d))
	form, ok := tagForms[tag]
	if tag == "<compat>" {
		form, v = compatForm(v, block)
	} else if !ok {
		form = tag
	}
	val := strings.ReplaceAll(compatJamo(jamo, string(v)), "⁄", "/")
	kind := contentKind(t, val)
	if form == "로마 숫자" || form == "분수" {
		kind = "수"
	}
	return formRow{Form: form, Kind: kind, Value: val}
}
