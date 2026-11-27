//ff:func feature=data type=helper control=sequence
//ff:what rowOf — the table line of a code point (tab-separated, no line end), or false when the character is not special, its form is not kept or its value is blank
package main

import (
	"slices"
	"strings"
)

func rowOf(t *ucdTable, jamo map[rune]rune, r rune, block string) (string, bool) {
	if skippedCategory(categoryOf(t, r)) || !special(t, r) {
		return "", false
	}
	row, ok := derive(t, jamo, r, block)
	if !ok || !slices.Contains(formsKept, row.Form) || strings.TrimSpace(row.Value) == "" {
		return "", false
	}
	return strings.Join([]string{string(r), row.Form, row.Kind, row.Value, formatValue(t, r)}, "\t"), true
}
