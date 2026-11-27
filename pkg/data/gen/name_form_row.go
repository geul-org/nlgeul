//ff:func feature=data type=helper control=sequence
//ff:what nameFormRow — the row of an enclosed character read from its name: its number, or its Latin letter; false when it holds neither
package main

func nameFormRow(t *ucdTable, r rune, form string) (formRow, bool) {
	info := t.Chars[r]
	if n, ok := parseNumeric(info.Numeric); ok {
		return formRow{Form: form, Kind: "수", Value: formatNumber(n)}, true
	}
	m := latinLetterName.FindStringSubmatch(info.Name)
	if m == nil {
		return formRow{}, false
	}
	letter := m[2]
	if m[1] == "SMALL" {
		letter = string(rune(letter[0]) + 32)
	}
	return formRow{Form: form, Kind: "로마자", Value: letter}, true
}
