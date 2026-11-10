//ff:func feature=token type=reader control=sequence
//ff:what TestVariants — a tail is read as it is and once for each stem the letters are a changed form of: 들 before 었 is 듣 too, before 고 not; a fused ㄴ under 사 is 살's; a tail that asks for a change is read only so; a stem that does not take the tail, a tail that shows its own stem, and one that begins with nothing give no more
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestVariants(t *testing.T) {
	sc := wordsScan("들었다", layerCost)
	past := []piece{{text: "었", tail: rule.TailPred}, {text: "다", tail: rule.TailPred}}
	got := sc.variants(0, tail{start: 1, pieces: past, head: headWhole})
	if len(got) != 2 || got[0].shown != "" || got[1].shown != "듣" {
		t.Fatal(got)
	}
	if got = sc.variants(0, tail{start: 1, pieces: past, head: headN}); len(got) != 1 {
		t.Fatal("ㄷ does not change before ㄴ", got)
	}
	live := wordsScan("산다", layerCost).variants(0, tail{start: 1, base: '사', pieces: []piece{{text: "(ㄴ)다", tail: rule.TailVerb}}, head: headN})
	if len(live) != 2 || live[1].shown != "살" {
		t.Fatal(live)
	}
	only := wordsScan("지은", layerCost).variants(0, tail{start: 1, pieces: []piece{{text: "은", tail: rule.TailPred}}, head: headEuForm, irr: irrS})
	if len(only) != 1 || only[0].shown != "짓" || only[0].irr != 0 {
		t.Fatal(only)
	}
	if got = sc.variants(0, tail{start: 1, pieces: []piece{{text: "이", tail: rule.TailCase}}, head: headWhole}); len(got) != 1 {
		t.Fatal("a verb's stem does not take a particle", got)
	}
	if got = sc.variants(0, tail{start: 1, pieces: past, head: headWhole, shown: "듣"}); len(got) != 1 {
		t.Fatal("a tail that shows its stem", got)
	}
	if got = sc.variants(0, tail{start: 1, pieces: past}); len(got) != 1 {
		t.Fatal("a tail that begins with nothing", got)
	}
	if got = sc.variants(0, tail{start: 1, pieces: past, head: headWhole, irr: irrS}); len(got) != 0 {
		t.Fatal("a tail that asks for another change", got)
	}
}
