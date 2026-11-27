//ff:func feature=parse type=helper control=sequence
//ff:what TestMidPart — a named candidate of a sign between letters is a part of that name and no stem (a middle dot); a candidate with no name is a stem named by the tables of special letters when they name the sign (=[기호:같음]), and a stem with no name when they do not
package parse

import "testing"

func TestMidPart(t *testing.T) {
	if p := midPart("·", "가운뎃점", '·'); p.Name != (Name{Kind: "가운뎃점"}) || p.Stem || p.Text != "·" {
		t.Fatal(p)
	}
	if p := midPart("=", "", '='); p.Name != (Name{"기호", "같음"}) || !p.Stem {
		t.Fatal(p)
	}
	if p := midPart("가", "", '가'); p.Name != (Name{}) || !p.Stem {
		t.Fatal(p)
	}
}
