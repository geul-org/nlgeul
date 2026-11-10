//ff:func feature=token type=reader control=sequence
//ff:what TestAfterMask — each attachment condition gives the ends it allows; a non-Hangul end is always allowed
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestAfterMask(t *testing.T) {
	all := rule.EndVowel | rule.EndL | rule.EndHard | rule.EndOther
	if afterMask(afterAny) != all || afterMask(afterBatchim) != all&^rule.EndVowel || afterMask(afterHard) != rule.EndHard|rule.EndOther || afterMask(afterOpen) != all&^rule.EndHard || afterMask(afterVowel) != rule.EndVowel|rule.EndOther {
		t.Fatal("afterMask")
	}
}
