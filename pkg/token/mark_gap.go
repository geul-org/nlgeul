//ff:func feature=token type=reader control=sequence
//ff:what markGap — the boundary after word i of a candidate's words as the spacing rules see it: the two words that meet there and the letters each takes up
package token

func markGap(words []step, i int) gap {
	start := 0
	if i > 0 {
		start = words[i-1].end
	}
	return gap{left: words[i].word, right: words[i+1].word, leftLen: words[i].end - start, rightLen: words[i+1].end - words[i].end}
}
