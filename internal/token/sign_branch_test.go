//ff:func feature=token type=helper control=sequence
//ff:what TestSignBranch — a sign is a reading of its own: one word of the one part given, ending where it is told to, standing (verdict 1) and costing nothing
package token

import (
	"reflect"
	"testing"
)

func TestSignBranch(t *testing.T) {
	p := Part{Text: "?", Name: Name{Kind: "물음표"}}
	want := Branch{Words: []Eojeol{{Parts: []Part{p}, End: 3}}, Verdict: 1}
	if got := signBranch(p, 3); !reflect.DeepEqual(got, want) || got.Pref != 0 {
		t.Fatal(got)
	}
}
