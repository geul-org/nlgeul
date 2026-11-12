//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestFusedNames — the names of the markers put into the vowel (어/아) or under the syllable (ㅆ) before them, each once and sorted, and no other: a small index gives exactly its two such names (not a whole form or a piece put under as ㄴ), an empty one none; the full index has 맺음:어(합침) and 선어말:ㅆ and not 맺음:다, and every marker whose lead is the vowel or ㅆ is in the list
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestFusedNames(t *testing.T) {
	small := &index{byKey: map[rune][]*entry{
		'다': {{name: "나", lead: rule.LeadVowel}, {name: "가", lead: rule.JongSS}},
		'라': {{name: "나", lead: rule.LeadVowel}, {name: "다"}, {name: "라", lead: rule.JongN}},
	}}
	if got := small.fusedNames(); !slices.Equal(got, []string{"가", "나"}) {
		t.Fatal("each name once, sorted, of the vowel and ㅆ only", got)
	}
	if got := (&index{}).fusedNames(); len(got) != 0 {
		t.Fatal("an empty index", got)
	}
	ix, _ := newIndex(Options{})
	got := ix.fusedNames()
	if len(got) == 0 || !slices.IsSorted(got) || len(slices.Compact(slices.Clone(got))) != len(got) {
		t.Fatal("sorted, each once", got)
	}
	if !slices.Contains(got, "맺음:어(합침)") || !slices.Contains(got, "선어말:ㅆ") || slices.Contains(got, "맺음:다") {
		t.Fatal("어 put into the vowel and ㅆ under it, not 다", got)
	}
	for key, entries := range ix.byKey {
		for _, e := range entries {
			if (e.lead == rule.LeadVowel || e.lead == rule.JongSS) && !slices.Contains(got, e.name) {
				t.Errorf("%c: %s is fused into the syllable before it", key, e.name)
			}
		}
	}
}
