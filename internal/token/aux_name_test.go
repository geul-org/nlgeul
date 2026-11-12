//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestAuxName — the name of a particle reading that is an auxiliary particle read after a linking closing; "" for a reading that is not one (a vocative 야, a case particle)
package token

import "testing"

func TestAuxName(t *testing.T) {
	if got := auxName(&entry{name: "보조사:야"}); got != "보조사:야" {
		t.Fatal(got)
	}
	for _, name := range []string{"호격:야", "주격:가"} {
		if got := auxName(&entry{name: name}); got != "" {
			t.Fatal(name, got)
		}
	}
}
