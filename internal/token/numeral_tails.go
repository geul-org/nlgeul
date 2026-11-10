//ff:func feature=token type=reader control=iteration dimension=1
//ff:what numeralTails — the tails of a piece after each number that is no run of digits: the number names the stem; after a lone Hangul numeral the piece's graph is asked again, with whether the two together are a word and whether the numeral is a lone Sino-Korean digit read with no word table
package token

import "github.com/geul-org/koparse/pkg/rule"

func (sc *scan) numeralTails(e *entry, start, j int, facts rule.Piece, pc piece, words []numeral) []tail {
	var out []tail
	for _, w := range words {
		q := pc
		facts.Whole, facts.Lone = sc.ix.wholeWord(w.text, e.form), sc.ix.loneDigit(w.text)
		if facts.Whole || facts.Lone {
			q.verdict = sc.judge(e, start, j, facts)
		}
		if q.verdict > 0 {
			out = append(out, tail{start: start, pieces: []piece{q}, only: w.text, as: numName, weight: w.pref})
		}
	}
	return out
}
