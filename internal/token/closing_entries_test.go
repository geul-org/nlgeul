//ff:func feature=token type=reader control=sequence
//ff:what TestClosingEntries — a plain closing one entry; one that takes 으 two, the 으 form after a consonant other than ㄹ, as 은, 을, 음 are; an 어/아 closing three whole forms and a fused one; a closing under a consonant one
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/internal/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestClosingEntries(t *testing.T) {
	plain := closingEntries(parse.Marker{Form: "지만", After: rule.AfterAny, Kind: rule.KindLink})
	eu := closingEntries(parse.Marker{Form: "면", Eu: true, Kind: rule.KindLink})
	a := closingEntries(parse.Marker{Form: "서", Lead: -1, Kind: rule.KindLink})
	n := closingEntries(parse.Marker{Form: "다", Lead: rule.JongN, Kind: rule.KindFinal, NoYo: true})
	if len(plain) != 1 || plain[0].name != "맺음:지만" || len(eu) != 2 || eu[1].form != "으면" || eu[1].after != rule.AfterHard || eu[0].after != rule.AfterOpen {
		t.Fatal(plain, eu)
	}
	adn := closingEntries(parse.Marker{Form: "을", After: rule.AfterBatchim, Kind: rule.KindAdn})
	polite := closingEntries(parse.Marker{Form: "습니다", After: rule.AfterBatchim, Kind: rule.KindFinal})
	if adn[0].after != rule.AfterHard || polite[0].after != rule.AfterBatchim {
		t.Fatal("을 attaches after a consonant other than ㄹ; 습니다 after any consonant", adn[0].after, polite[0].after)
	}
	if len(a) != 4 || a[0].form != "어서" || a[3].lead != rule.LeadVowel || a[3].form != "서" || len(n) != 1 || n[0].name != "맺음:ㄴ다" || !n[0].noYo {
		t.Fatal(a, n)
	}
}
