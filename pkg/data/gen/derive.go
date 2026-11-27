//ff:func feature=data type=helper control=sequence
//ff:what derive — the row of a character (form, content, value) from its decomposition, its name or its number; false when none of them gives one
package main

func derive(t *ucdTable, jamo map[rune]rune, r rune, block string) (formRow, bool) {
	tag := tagOf(t.Chars[r].Decomp)
	if tag != "" && !isPresentationTag(tag) {
		return deriveTagged(t, jamo, r, block), true
	}
	if row, found, ok := deriveByName(t, r, block); found {
		return row, ok
	}
	return deriveNumeric(t, r, block)
}
