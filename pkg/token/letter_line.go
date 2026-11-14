//ff:func feature=token type=helper control=sequence
//ff:what letterLine — the letters of a stretch, as they are (the letters of the sentence without its spaces, from Start up to End, kept inside the sentence), for a stretch nothing is known of
package token

func letterLine(span Span, letters []rune) string {
	start := min(max(span.Start, 0), len(letters))
	end := min(max(span.End, start), len(letters))
	return string(letters[start:end])
}
