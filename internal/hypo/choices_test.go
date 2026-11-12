//ff:func feature=hypo type=helper control=sequence
//ff:what TestChoices — the bundles of readings asleep are left out and the rest keep their order; no bundle, and bundles all asleep, give none
package hypo

import (
	"testing"

	"github.com/geul-org/nlgeul/internal/token"
)

func TestChoices(t *testing.T) {
	bundles := []token.Bundle{
		{Unit: 0, Text: "가", Branches: []token.Branch{{Verdict: 0.9}, {Verdict: 0.5}}},
		{Unit: 0, Text: "가나", Branches: []token.Branch{{Verdict: 0.2, Asleep: true}, {Verdict: 0.1, Asleep: true}}},
		{Unit: 1, Text: "다", Branches: []token.Branch{{Verdict: 1}}},
		{Unit: 2, Text: "라", Branches: []token.Branch{{Verdict: 0.7}, {Verdict: 0.6}, {Verdict: 0.3}}},
		{Unit: 2, Text: "라마", Branches: []token.Branch{{Verdict: 0.1, Asleep: true}}},
		{Unit: 2, Text: ".", Branches: []token.Branch{{Verdict: 1}}},
	}
	got := Choices(bundles)
	if len(got) != 4 {
		t.Fatalf("%d bundles, want 4: %+v", len(got), got)
	}
	if got[0].Text != "가" || got[1].Text != "다" || got[2].Text != "라" || got[3].Text != "." {
		t.Fatalf("texts %q %q %q %q, want 가 다 라 .", got[0].Text, got[1].Text, got[2].Text, got[3].Text)
	}
	if got[0].Unit != 0 || got[1].Unit != 1 || got[2].Unit != 2 || got[3].Unit != 2 {
		t.Fatalf("units %d %d %d %d, want 0 1 2 2", got[0].Unit, got[1].Unit, got[2].Unit, got[3].Unit)
	}
	if len(got[0].Branches) != 2 || len(got[1].Branches) != 1 || len(got[2].Branches) != 3 || got[2].Branches[2].Verdict != 0.3 {
		t.Fatalf("branches changed: %+v", got)
	}

	if got := Choices(nil); len(got) != 0 {
		t.Fatalf("no bundle gives %+v", got)
	}
	if got := Choices([]token.Bundle{bundles[1], bundles[4]}); len(got) != 0 {
		t.Fatalf("bundles all asleep give %+v", got)
	}
}
