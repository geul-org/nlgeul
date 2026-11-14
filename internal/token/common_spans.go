//ff:func feature=token type=helper control=iteration dimension=2
//ff:what commonSpans — the spans every line begins with, in order; none for no line
package token

func commonSpans(lines [][]Span) []Span {
	if len(lines) == 0 {
		return nil
	}
	out := lines[0]
	for _, ln := range lines[1:] {
		k := 0
		for k < len(out) && k < len(ln) && out[k] == ln[k] {
			k++
		}
		out = out[:k]
	}
	return out
}
