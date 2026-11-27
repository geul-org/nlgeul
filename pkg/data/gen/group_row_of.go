//ff:func feature=data type=helper control=sequence
//ff:what groupRowOf — the group table line of a code point (tab-separated, no line end), or false when its category is not a plain punctuation or symbol one or the form table already has a line for it
package main

import (
	"slices"
	"strings"
)

func groupRowOf(t *ucdTable, jamo map[rune]rune, r rune, block string) (string, bool) {
	if !slices.Contains(groupCategories, categoryOf(t, r)) {
		return "", false
	}
	if _, ok := rowOf(t, jamo, r, block); ok {
		return "", false
	}
	group, name := groupOf(t, r, block)
	return strings.Join([]string{string(r), group, name, string(r), "-"}, "\t"), true
}
