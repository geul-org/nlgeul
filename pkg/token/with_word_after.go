//ff:func feature=token type=helper control=iteration dimension=1
//ff:what withWordAfter — the spans the lines have in common, without the last of them when a line ends with it: a cut at the end of such a span leaves no word after it for that line, so that nothing could be read from the cut on
package token

func withWordAfter(common []Span, lines [][]Span) []Span {
	for _, ln := range lines {
		if len(common) >= len(ln) {
			common = common[:max(len(ln)-1, 0)]
		}
	}
	return common
}
