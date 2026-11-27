//ff:func feature=data type=helper control=sequence
//ff:what deriveNumeric — the row of a number character with no decomposition that stands in a number block: a Roman numeral or another number form
package main

import (
	"slices"
	"strings"
)

func deriveNumeric(t *ucdTable, r rune, block string) (formRow, bool) {
	info := t.Chars[r]
	n, ok := parseNumeric(info.Numeric)
	if !ok || (info.Category != "No" && info.Category != "Nl") || !slices.Contains(numberBlocks, block) {
		return formRow{}, false
	}
	form := "수 꼴"
	if strings.Contains(info.Name, "ROMAN") {
		form = "로마 숫자"
	}
	return formRow{Form: form, Kind: "수", Value: formatNumber(n)}, true
}
