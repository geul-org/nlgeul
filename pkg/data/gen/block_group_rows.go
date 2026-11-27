//ff:func feature=data type=helper control=iteration dimension=1
//ff:what blockGroupRows — the group table lines of one block in code order; private use and surrogate blocks give none
package main

import "strings"

func blockGroupRows(t *ucdTable, jamo map[rune]rune, b blockRange) []string {
	if strings.Contains(b.Name, "Private Use") || strings.Contains(b.Name, "Surrogates") {
		return nil
	}
	var out []string
	for r := b.Lo; r <= b.Hi; r++ {
		if line, ok := groupRowOf(t, jamo, r, b.Name); ok {
			out = append(out, line)
		}
	}
	return out
}
