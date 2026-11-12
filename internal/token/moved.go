//ff:func feature=token type=reader control=iteration dimension=1
//ff:what moved — the readings of a unit whose head is a sign of a pair, whose first word has the sign as its first part, carried over to the letters after the sign (movedOne); the readings that took the sign into a word's letters are left
package token

func (sc *scan) moved(paths []*path, head int) []*path {
	var out []*path
	for _, p := range paths {
		words := wordsOf(p)
		if len(words) == 0 || len(words[0].word.Parts) == 0 {
			continue
		}
		if words[0].word.Parts[0].Name.Kind != pairName {
			continue
		}
		if np := sc.movedOne(words, head); np != nil {
			out = append(out, np)
		}
	}
	return out
}
