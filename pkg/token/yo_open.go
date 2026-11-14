//ff:func feature=token type=reader control=sequence
//ff:what yoOpen — whether a piece before the particle 요 stands where its letters are those of a closing written wrong: right on its stem after a syllable with no final consonant (onStem), or right after the suffix 이 that stands on a stem (a letter 이 at start-1 under a syllable with a final consonant, which is where that suffix stands)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The suffix 이 stands only on a stem that ends in a consonant and is
// the first piece of a chain, so what stands left of it is read from
// the letters. The ending of a word with that suffix is 이, which has no
// final consonant: the second condition of the case on a stem is true
// of it always.
func (sc *scan) yoOpen(p rule.Piece, start int, onStem bool) bool {
	if p.Next != rule.LinkYo {
		return false
	}
	if onStem && rule.JongOf(p.Prev) == 0 {
		return true
	}
	return start >= 2 && sc.s[start-1] == '이' && rule.JongOf(sc.s[start-2]) > 0
}
