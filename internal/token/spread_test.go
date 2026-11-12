//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestSpread — a unit headed by a sign of a pair gives each candidate that names the head once for each way the sign is read (열기:1, 열기:2, …, with its depth), and once more with the head as its letter alone when the sign is also read so; a candidate that does not name the head stays as it is; a unit of another kind, and no contest, are given back as they are
package token

import (
	"reflect"
	"testing"
)

func TestSpread(t *testing.T) {
	headedBranch := Branch{Words: []Eojeol{{Parts: []Part{{Text: "(", Name: Name{Kind: pairName}, Stem: true}}, End: 1}}, Verdict: 1, Pref: -5}
	whole := Branch{Words: []Eojeol{{Parts: []Part{{Text: "(a", Stem: true}}, End: 2}}, Verdict: 0.33, Pref: -9}
	u := unit{text: "(", head: 1, name: pairName, candidates: []pairing{{openName, 1}, {openName, 2}}, bare: true}
	got := spread([]Contest{{Text: "(", Branches: []Branch{headedBranch, whole}}}, u)
	if len(got) != 1 || len(got[0].Branches) != 4 {
		t.Fatal(got)
	}
	br := got[0].Branches
	for i, depth := range []int{1, 2} {
		p := br[i].Words[0].Parts[0]
		if p.Name != (Name{Kind: openName, Form: string(rune('0' + depth))}) || p.Depth != depth || br[i].Verdict != 1 || br[i].Pref != -5 {
			t.Errorf("candidate %d: %+v", i, br[i])
		}
	}
	if p := br[2].Words[0].Parts[0]; p.Name != (Name{}) || p.Depth != 0 || p.Text != "(" || br[2].Pref != -5 {
		t.Fatal("the head as its letter", br[2])
	}
	if !reflect.DeepEqual(br[3], whole) {
		t.Fatal("a candidate of the whole unit stays", br[3])
	}
	plain := []Contest{{Text: "곰", Branches: []Branch{headedBranch}}}
	if got = spread(plain, unit{text: "곰"}); !reflect.DeepEqual(got, plain) {
		t.Fatal("a plain unit", got)
	}
	if got = spread(nil, u); got != nil {
		t.Fatal("no contest", got)
	}
	single := spread([]Contest{{Text: ")", Branches: []Branch{headedBranch}}}, unit{text: ")", head: 1, name: pairName, candidates: []pairing{{shutName, 3}}})
	if p := single[0].Branches[0].Words[0].Parts[0]; len(single[0].Branches) != 1 || p.Name.Kind != shutName || p.Name.Form != "3" || p.Depth != 3 {
		t.Fatal("one way, not also a letter", single)
	}
}
