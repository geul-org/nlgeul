//ff:func feature=token type=rule control=sequence
//ff:what TestVerdictOf — a marker with a stem stands (1); without one it is defeated with the reason; a graph that fails or claims nothing gives -1
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestVerdictOf(t *testing.T) {
	ix := newIndex()
	e := ix.byKey['은'][0]
	if v, why := verdictOf(e.g, rule.Piece{HasStem: true, PrevEnd: rule.EndHard, Next: -1}, "것은"); v != 1 || why != nil {
		t.Fatal(v, why)
	}
	v, why := verdictOf(e.g, rule.Piece{PrevEnd: rule.EndVowel, Next: -1}, "은")
	if v != -1 || len(why) != 2 || why[0] != "앞에 줄기가 없다" {
		t.Fatal(v, why)
	}
	bad := toulmin.NewGraph("bad")
	bad.Rule(func(toulmin.Context, toulmin.Specs) (bool, any) { panic("x") })
	if v, _ := verdictOf(bad, rule.Piece{}, ""); v != -1 {
		t.Fatal(v)
	}
	if v, _ := verdictOf(toulmin.NewGraph("empty"), rule.Piece{}, ""); v != -1 {
		t.Fatal(v)
	}
}
