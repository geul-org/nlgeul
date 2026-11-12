//ff:func feature=token type=helper control=sequence
//ff:what TestWithHead — a copy of a reading whose first part has the names and the depth given; the reading itself, its other parts and its other words are not changed
package token

import (
	"reflect"
	"testing"
)

func TestWithHead(t *testing.T) {
	br := Branch{Words: []Eojeol{
		{Parts: []Part{{Text: "(", Names: []Name{{Kind: pairName}}}, {Text: "를"}}, End: 2},
		{Parts: []Part{{Text: "나"}}, End: 3},
	}, Verdict: 1, Pref: -2}
	got := withHead(br, []Name{{Kind: shutName, Form: "2"}}, 2)
	first := got.Words[0].Parts[0]
	if !reflect.DeepEqual(first.Names, []Name{{Kind: shutName, Form: "2"}}) || first.Depth != 2 || first.Text != "(" {
		t.Fatal(first)
	}
	if got.Words[0].Parts[1].Text != "를" || got.Words[1].Parts[0].Text != "나" || got.Words[0].End != 2 || got.Verdict != 1 || got.Pref != -2 {
		t.Fatal("the rest is kept", got)
	}
	if br.Words[0].Parts[0].Names[0].Kind != pairName || br.Words[0].Parts[0].Depth != 0 {
		t.Fatal("the reading is changed", br)
	}
	if bare := withHead(br, nil, 0); len(bare.Words[0].Parts[0].Names) != 0 || bare.Words[0].Parts[0].Depth != 0 {
		t.Fatal("the head as its letter", bare)
	}
}
