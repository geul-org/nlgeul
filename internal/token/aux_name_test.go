//ff:func feature=token type=helper control=sequence
//ff:what TestAuxName — the auxiliary reading among a form's readings is named; none gives ""
package token

import "testing"

func TestAuxName(t *testing.T) {
	if got := auxName([]*entry{{name: "호격:야"}, {name: "보조사:야"}}); got != "보조사:야" {
		t.Fatal(got)
	}
	if got := auxName([]*entry{{name: "주격:가"}}); got != "" {
		t.Fatal(got)
	}
}
