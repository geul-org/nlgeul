//ff:func feature=token type=reader control=sequence
//ff:what TestOnStemYo — 에 right after the suffix 이 (영숙이에요) stands rebutted as it is given, not only when asked on a stem; before 요 the particle readings right on a stem are asked once more and told so: 에 after a syllable with no final consonant (고양이에요) stands rebutted, after one with a final consonant (집에요) or after nothing (a unit 에요 of its own, read as attached to the unit before) as it stood; with no 요 after it nothing is asked again
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestOnStemYo(t *testing.T) {
	verdictAt := func(text string, j, next int, attach bool) (float64, float64) {
		sc := scanOf(text)
		sc.attach = attach
		pieces, from, _ := sc.readGroup(classParticle, j, next, false)
		ends, _ := sc.onStem(j, next, pieces, from)
		return pieceNamed(pieces, "부사격:에").verdict, pieceNamed(ends, "부사격:에").verdict
	}
	if given, ends := verdictAt("결과에요", 3, rule.LinkYo, false); given != 1 || ends >= 0.34 || ends <= 0 {
		t.Fatal("after a vowel", given, ends)
	}
	if given, ends := verdictAt("영숙이에요", 4, rule.LinkYo, false); given >= 0.34 || given <= 0 || ends != given {
		t.Fatal("right after the suffix 이 on a stem", given, ends)
	}
	if given, ends := verdictAt("학교이에요", 4, rule.LinkYo, false); given != 1 || ends >= 0.34 {
		t.Fatal("after 이 under a syllable with no final consonant is no suffix, but on a stem it is asked as one", given, ends)
	}
	if given, ends := verdictAt("집에요", 2, rule.LinkYo, false); given != 1 || ends != 1 {
		t.Fatal("after a final consonant", given, ends)
	}
	if given, ends := verdictAt("에요", 1, rule.LinkYo, true); given != 1 || ends != 1 {
		t.Fatal("with no stem before it", given, ends)
	}
	if given, ends := verdictAt("고양이에", 4, -1, false); given != 1 || ends != 1 {
		t.Fatal("with no 요 after it", given, ends)
	}
}
