//ff:func feature=parse type=helper control=sequence
//ff:what bareName — the name of a candidate of a sign read between letters: the candidate's own name, or, when it has none (bareApart: =), the first name the tables of special letters give the sign (=[기호:같음]); none when they give none either
package parse

func bareName(name string, sign rune) string {
	if names := formNames(sign); name == "" && len(names) > 0 {
		return names[0]
	}
	return name
}
