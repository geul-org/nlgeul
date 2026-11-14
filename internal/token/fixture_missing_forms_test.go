//ff:func feature=token type=helper control=iteration dimension=1
//ff:what missingForms — the ways of cutting that are not given exactly once by the segments (keys, as segmentKeys counts them)
package token

func missingForms(forms map[string]bool, keys map[string]int) []string {
	var bad []string
	for key := range forms {
		if keys[key] != 1 {
			bad = append(bad, "form "+key)
		}
	}
	return bad
}
