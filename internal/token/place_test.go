//ff:func feature=token type=reader control=sequence
//ff:what TestPlace — a whole form reports the end before it; a form at the unit's start has no stem; a piece under a consonant or fused into a vowel sits in the stem's last syllable; a form that is not there is not placed
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestPlace(t *testing.T) {
	sc := scanOf("다음으로간다해서")
	if start, p, ok := sc.place(&entry{form: "으로"}, 4, false); !ok || start != 2 || !p.HasStem || p.Prev != '음' {
		t.Fatal(start, p, ok)
	}
	if _, p, ok := sc.place(&entry{form: "다"}, 1, false); !ok || p.HasStem {
		t.Fatal(p, ok)
	}
	if start, p, ok := sc.place(&entry{form: "다", lead: jongN}, 6, false); !ok || start != 5 || p.Base != "가" {
		t.Fatal(start, p, ok)
	}
	if _, _, ok := sc.place(&entry{form: "다", lead: rule.JongL}, 6, false); ok {
		t.Fatal("wrong consonant")
	}
	if start, _, ok := sc.place(&entry{form: "서", lead: leadVowel}, 8, false); !ok || start != 7 {
		t.Fatal(start, ok)
	}
	if _, _, ok := sc.place(&entry{form: "로써"}, 4, false); ok {
		t.Fatal("form not there")
	}
	if _, _, ok := sc.place(&entry{form: "다음으로간"}, 3, false); ok {
		t.Fatal("longer than the text")
	}
	if _, _, ok := sc.place(&entry{lead: jongN}, 0, false); ok {
		t.Fatal("a fused piece at the start")
	}
}
