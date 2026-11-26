//ff:func feature=parse type=helper control=sequence
//ff:what TestLowered — a candidate's verdict is brought down to the verdict given, in the branch and in a copy of the score of its first word; a candidate whose verdict is already no higher is given back as it is, and one with no scores kept has only its own verdict lowered
package parse

import (
	"reflect"
	"testing"
)

func TestLowered(t *testing.T) {
	scores := []wordScore{{verdict: 1, pref: -5, born: 3}, {verdict: 0.9, pref: -1, born: 4}}
	br := Branch{Words: []Eojeol{{End: 1}, {End: 2}}, Verdict: 1, Pref: -6, scores: scores}
	got := lowered(br, 0.25)
	if got.Verdict != 0.25 || got.Pref != -6 || got.scores[0] != (wordScore{verdict: 0.25, pref: -5, born: 3}) || got.scores[1] != scores[1] {
		t.Fatal(got)
	}
	if scores[0].verdict != 1 || br.Verdict != 1 {
		t.Fatal("the candidate given is changed", br, scores)
	}
	if same := lowered(br, 1); !reflect.DeepEqual(same, br) {
		t.Fatal("no higher already", same)
	}
	weak := Branch{Verdict: 0.5, scores: []wordScore{{verdict: 0.5}}}
	if same := lowered(weak, 0.75); !reflect.DeepEqual(same, weak) {
		t.Fatal("lower already", same)
	}
	if bare := lowered(Branch{Verdict: 1}, 0.5); bare.Verdict != 0.5 || bare.scores != nil {
		t.Fatal("no scores kept", bare)
	}
}
