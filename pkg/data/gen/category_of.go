//ff:func feature=data type=helper control=sequence
//ff:what categoryOf — the general category of a code point; one that UnicodeData.txt does not give is Cn
package main

func categoryOf(t *ucdTable, r rune) string {
	if info, ok := t.Chars[r]; ok {
		return info.Category
	}
	return rangeCategory(t, r)
}
