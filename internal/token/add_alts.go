//ff:func feature=token type=reader control=iteration dimension=2
//ff:what addAlts — the changed forms of every predicate stem of the word table, each with the stems it stands for (in the order of the stems); the stems the table calls regular that fuse 어 as ㅎ stems do; and for the changes that leave the stem's last syllable looking unlike what a tail asks for (ㄷ, ㅅ), from the changed last syllable to the stem's own
package token

import "sort"

func (ix *index) addAlts() {
	ix.alts, ix.irrEnds = map[string][]alt{}, map[rune][]alt{}
	keys := make([]string, 0, len(ix.lex))
	for key, w := range ix.lex {
		if w.kinds&(kindVerb|kindAdj) != 0 {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	for _, key := range keys {
		for _, c := range changed(key, ix.lex[key].irr) {
			ix.addAlt(c.stem, key, c.kind)
		}
		if form, ok := eoStems[key]; ok {
			ix.addAlt(form, key, irrHE)
		}
	}
}
