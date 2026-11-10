//ff:func feature=token type=reader control=sequence
//ff:what TestOwnName — a word of two sets is named by the set whose kind took what follows; by both when both did; by none when it stands as another kind, and a word of no set has no such name
package token

import "testing"

func TestOwnName(t *testing.T) {
	two := &entry{name: "관형사:두|보조 용언:두"}
	if ownName(two, kindVerb) != "보조 용언:두" || ownName(two, kindDeterminer|kindVerb) != two.name || ownName(two, kindNominal) != "" || ownName(nil, kindNominal) != "" {
		t.Fatal(ownName(two, kindVerb), ownName(two, kindNominal))
	}
}
