//ff:func feature=token type=helper control=iteration dimension=1
//ff:what lineWords — the words of a line of streamLines (the candidates joined with " "), parted at the spaces outside the brackets of the names (a name such as 보조 용언 holds a space)
package token

func lineWords(line string) []string {
	var out []string
	depth, start := 0, 0
	for i, c := range line {
		switch {
		case c == '[':
			depth++
		case c == ']':
			depth--
		case c == ' ' && depth == 0:
			out, start = append(out, line[start:i]), i+1
		}
	}
	return append(out, line[start:])
}
