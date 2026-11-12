//ff:func feature=token type=rule control=sequence
//ff:what codaStem — is the stem from b before tail t a new word that ends in the final consonant ㅆ with markers after it (떨어졌 before 다, 졌 before 다): the new word's graph then does not let it stand (a stem that ends in ㅆ is 있 or a word made with it, which the word table lists; any other ㅆ is the slot 았/었/겠); the stem is judged as it was before a piece was fused into its last syllable; a word some table knows and a tail that names its own stem are not asked, nor is a whole new word with no marker; only with a word table (without one 맛있 is not told from such a stem)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func (sc *scan) codaStem(b int, t tail) bool {
	if !sc.ix.lexed || len(t.pieces) == 0 || t.shown != "" || t.as != "" {
		return false
	}
	stem := string(sc.s[b:t.start])
	if t.base != 0 {
		stem = string(sc.s[b:t.start-1]) + string(t.base)
	}
	if v, _ := verdictOf(sc.ix.stem, rule.Piece{Form: stem, Tail: tailClass(t.pieces)}, ""); v > 0 {
		return false
	}
	_, _, known := sc.stemOf(b, t)
	return known <= 0
}
