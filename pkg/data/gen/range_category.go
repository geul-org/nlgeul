//ff:func feature=data type=helper control=iteration dimension=1
//ff:what rangeCategory — the category of the First/Last range a code point is in, or Cn when it is in none
package main

func rangeCategory(t *ucdTable, r rune) string {
	for _, rg := range t.Ranges {
		if r >= rg.Lo && r <= rg.Hi {
			return rg.Category
		}
	}
	return "Cn"
}
