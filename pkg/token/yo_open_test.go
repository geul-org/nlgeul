//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestYoOpen — before 요 a piece is told it stands where its letters are a closing written wrong when it stands on a stem under a syllable with no final consonant (결과에요), or right after the letter 이 that stands under a syllable with a final consonant (영숙이에요); not when 요 does not follow, not on a stem under a final consonant (집에요), not after 이 under a syllable with none when it is not asked on a stem (학교이에요), not at the very start
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestYoOpen(t *testing.T) {
	for _, c := range []struct {
		text   string
		start  int
		prev   rune
		next   int
		onStem bool
		want   bool
	}{
		{"결과에요", 2, '과', rule.LinkYo, true, true},
		{"결과에요", 2, '과', rule.LinkYo, false, false},
		{"결과에요", 2, '과', -1, true, false},
		{"집에요", 1, '집', rule.LinkYo, true, false},
		{"영숙이에요", 3, '이', rule.LinkYo, false, true},
		{"학교이에요", 3, '이', rule.LinkYo, false, false},
		{"학교이에요", 3, '이', rule.LinkYo, true, true},
		{"이에요", 1, '이', rule.LinkYo, false, false},
		{"영숙이에요", 3, '이', -1, false, false},
		{"에요", 0, 0, rule.LinkYo, true, false},
	} {
		sc := scanOf(c.text)
		if got := sc.yoOpen(rule.Piece{Prev: c.prev, Next: c.next}, c.start, c.onStem); got != c.want {
			t.Errorf("%s %d %v: %v", c.text, c.start, c.onStem, got)
		}
	}
}
