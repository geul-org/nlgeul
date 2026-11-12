//ff:func feature=token type=rule control=sequence
//ff:what TestMisspelt — the right form of a misspelt stem is named by the kinds both lines give and judged by those kinds' graphs, told that it is a misspelt form: it stands, lower than the right form written right and above 0, where its kind takes the tail; it does not where the kind takes no such tail or the kinds share none; in the first layer it stands whole, and known is as before
package token

import (
	"math"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMisspelt(t *testing.T) {
	ix := typoScan("x", layerCost).ix
	_, right := ix.known("맡기", rule.TailPred, false)
	name, v := ix.misspelt("맡기", kindVerb, rule.TailPred, false)
	if name != "열쇠:동사" || v <= 0 || v >= right || math.Abs(v-2/(1+(2/(1+right)-1)+forceTypo)+1) > 1e-9 {
		t.Fatal(name, v, right)
	}
	if name, v = ix.misspelt("맡기", kindVerb, rule.TailCase, false); name != "" || v != 0 {
		t.Fatal("a verb does not take a particle", name, v)
	}
	if name, v = ix.misspelt("맡기", kindNominal, rule.TailPred, false); name != "" || v != 0 {
		t.Fatal("the misspelt line and the word share no kind", name, v)
	}
	if name, v = ix.misspelt("고양이", kindVerb, rule.TailPred, false); name != "" || v != 0 {
		t.Fatal("a word of another kind", name, v)
	}
	if name, v = ix.misspelt("물", kindNominal|kindVerb, rule.TailCase, false); name != "열쇠:체언" || v <= 0 {
		t.Fatal("only the kinds that take the tail are named", name, v)
	}
	if name, v = ix.misspelt("물", kindNominal|kindVerb, rule.TailBare, true); name != "열쇠:체언" || v <= 0 {
		t.Fatal(name, v)
	}
	one := typoScan("x", layerKnown).ix
	if name, v = one.misspelt("맡기", kindVerb, rule.TailCase, false); name != "열쇠:동사" || v != 1 {
		t.Fatal("in the first layer every kind stands", name, v)
	}
	if name, v = ix.known("맡기", rule.TailPred, false); name != "열쇠:동사" || v != right {
		t.Fatal("known is as before", name, v)
	}
}
