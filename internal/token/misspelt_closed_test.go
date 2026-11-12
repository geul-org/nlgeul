//ff:func feature=token type=rule control=sequence
//ff:what TestMisspeltClosed — the right form of a misspelt stem that is a closed-set word (꺼 for the dependent noun 거) stands with its set's kind though the word table has no line for it, and is named by its set (의존명사:거) when it stands as that set's kind; the right form of a word the table alone knows is named by its kind as before
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMisspeltClosed(t *testing.T) {
	ix := typoScan("x", layerCost).ix
	name, v := ix.misspelt("거", kindNominal, rule.TailBare, true)
	if name != "의존명사:거" || v <= 0 {
		t.Fatal(name, v)
	}
	if _, right := ix.known("거", rule.TailBare, true); v >= right {
		t.Fatal("below the same word written right", v, right)
	}
	if name, v = ix.misspelt("거", kindNominal, rule.TailCase, false); name != "의존명사:거" || v <= 0 {
		t.Fatal("before a particle", name, v)
	}
	if name, v = ix.misspelt("거", kindVerb, rule.TailBare, true); name != "" || v != 0 {
		t.Fatal("a kind the closed-set word does not have", name, v)
	}
	if name, v = ix.misspelt("역할", kindNominal, rule.TailBare, true); name != "열쇠:체언" || v <= 0 {
		t.Fatal("a word only the table knows", name, v)
	}
}
