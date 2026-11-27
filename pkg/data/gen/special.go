//ff:func feature=data type=helper control=sequence
//ff:what special — whether a character counts as special: punctuation, symbols, number letters, and letters or numbers with a compatibility decomposition
package main

func special(t *ucdTable, r rune) bool {
	cat := categoryOf(t, r)
	if cat[0] == 'P' || cat[0] == 'S' || cat == "Nl" || cat == "No" {
		return true
	}
	if (cat[0] == 'L' || cat[0] == 'N') && r > 0x7F {
		tag := tagOf(t.Chars[r].Decomp)
		return tag != "" && !isPresentationTag(tag)
	}
	return false
}
