//ff:func feature=token type=reader control=sequence
//ff:what TestOwnNames — the names of a closed-set word are those of the sets whose kind took what follows, each apart; none for a word of no set or when none of its sets' kinds took it
package token

import (
	"reflect"
	"testing"
)

func TestOwnNames(t *testing.T) {
	two := &entry{name: "관형사:두|보조 용언:두"}
	if got := ownNames(two, kindDeterminer|kindVerb); !reflect.DeepEqual(got, []string{"관형사:두", "보조 용언:두"}) {
		t.Fatal(got)
	}
	if got := ownNames(two, kindVerb); !reflect.DeepEqual(got, []string{"보조 용언:두"}) {
		t.Fatal(got)
	}
	if got := ownNames(two, kindNominal); got != nil {
		t.Fatal("none of its sets", got)
	}
	if ownNames(nil, kindNominal) != nil {
		t.Fatal("a word of no set")
	}
}
