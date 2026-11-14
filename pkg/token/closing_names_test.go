//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestClosingNames — no rows give no names; a row of an 어/아 closing gives the four names of its entries (어, 아, 여 and the fused one), one that takes 으 two (면서, 으면서), the gated naming closing 이름:기; connNames, mustNames, auxClosings and nounHeads are so made, in the order of their rows, and each of their names is that of a closing entry of the index (with a word table) under the key of its form
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestClosingNames(t *testing.T) {
	if got := closingNames(); got != nil {
		t.Fatal("no rows", got)
	}
	vowel := []string{"맺음:어", "맺음:아", "맺음:여", "맺음:어(합침)"}
	if !slices.Equal(connNames, vowel) || !slices.Equal(closingNames(parse.Marker{Lead: rule.LeadVowel}), vowel) {
		t.Fatal(connNames)
	}
	if want := []string{"맺음:어야", "맺음:아야", "맺음:여야", "맺음:어야(합침)"}; !slices.Equal(mustNames, want) {
		t.Fatal(mustNames)
	}
	want := slices.Concat([]string{"맺음:지", "맺음:고", "맺음:게"}, vowel, []string{"맺음:어서", "맺음:아서", "맺음:여서", "맺음:어서(합침)", "맺음:면서", "맺음:으면서"})
	if !slices.Equal(auxClosings, want) {
		t.Fatal(auxClosings)
	}
	if !slices.Equal(nounHeads, []string{"이름:기"}) {
		t.Fatal(nounHeads)
	}
	rows := []parse.Marker{{Lead: rule.LeadVowel}, {Form: "야", Lead: rule.LeadVowel}, {Form: "지"}, {Form: "고"}, {Form: "게"}, {Lead: rule.LeadVowel}, {Form: "서", Lead: rule.LeadVowel}, {Form: "면서", Eu: true}}
	var entries []*entry
	for _, m := range append(rows, gatedClosings...) {
		entries = append(entries, closingEntries(m)...)
	}
	names := slices.Concat(connNames, mustNames, auxClosings, nounHeads)
	if len(entries) != len(names) {
		t.Fatal(len(entries), names)
	}
	ix := wordsScan("x", layerCost).ix
	for i, e := range entries {
		if got := ix.named(entryKey(e), names[i]); got.g == nil || got.class != classClosing || got.form != e.form {
			t.Errorf("%s: %+v", names[i], got)
		}
	}
}
