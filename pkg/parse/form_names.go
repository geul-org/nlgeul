//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what formNames — the names the tables of special letters give a letter, one for each of its lines in their order (㈎: 괄호:한글); none for a letter the tables lack
package parse

func formNames(r rune) []string {
	var out []string
	for _, f := range forms[r] {
		out = append(out, f.name())
	}
	return out
}
