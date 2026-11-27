//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what formsByLetter — the lines of the tables of special letters by their letter, each letter's lines in the order they are written
package parse

func formsByLetter(list []form) map[rune][]form {
	out := map[rune][]form{}
	for _, f := range list {
		out[f.letter] = append(out[f.letter], f)
	}
	return out
}
