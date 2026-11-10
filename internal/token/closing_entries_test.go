//ff:func feature=token type=reader control=sequence
//ff:what TestClosingEntries — a plain closing one entry; one that takes 으 two; an 어/아 closing three whole forms and a fused one; a closing under a consonant one
package token

import (
	"testing"

	"github.com/geul-org/koparse/internal/parse"
	"github.com/geul-org/koparse/pkg/rule"
)

func TestClosingEntries(t *testing.T) {
	plain := closingEntries(parse.Marker{Form: "지만", After: afterAny, Kind: rule.KindLink})
	eu := closingEntries(parse.Marker{Form: "면", Eu: true, Kind: rule.KindLink})
	a := closingEntries(parse.Marker{Form: "서", Lead: -1, Kind: rule.KindLink})
	n := closingEntries(parse.Marker{Form: "다", Lead: jongN, Kind: rule.KindFinal, NoYo: true})
	if len(plain) != 1 || plain[0].name != "맺음:지만" || len(eu) != 2 || eu[1].form != "으면" || eu[1].after != afterBatchim || eu[0].after != afterOpen {
		t.Fatal(plain, eu)
	}
	if len(a) != 4 || a[0].form != "어서" || a[3].lead != leadVowel || a[3].form != "서" || len(n) != 1 || n[0].name != "맺음:ㄴ다" || !n[0].noYo {
		t.Fatal(a, n)
	}
}
