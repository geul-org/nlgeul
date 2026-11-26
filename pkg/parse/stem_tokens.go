//ff:func feature=parse type=reader control=sequence
//ff:what stemTokens — how many tokens the stem from b to end holds for preference when it stands as the candidate given: a path read with the Hangul after its last slash as its name is one token, as a path is (namedTail); any other stem as tokens counts it
package parse

func (sc *scan) stemTokens(b, end int, r stemCandidate) int {
	if r.v > 0 && r.name == pathName && len(sc.namedTail(b, end)) > 0 {
		return 1
	}
	return sc.tokens(b, end)
}
