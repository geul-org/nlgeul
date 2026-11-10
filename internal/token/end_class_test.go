//ff:func feature=token type=reader control=sequence
//ff:what TestEndClass — a vowel, ㄹ, another consonant, not Hangul
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestEndClass(t *testing.T) {
	if endClass('가') != rule.EndVowel || endClass('갈') != rule.EndL || endClass('각') != rule.EndHard || endClass('x') != rule.EndOther {
		t.Fatal("endClass")
	}
}
