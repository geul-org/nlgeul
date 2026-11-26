//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestBareUnits — 당 right after a unit sign with no number before it is a tail whose stem is the unit, named as the unit table names it (㎡당); nothing for a suffix that needs a number (짜리), a unit written in Hangul (개당, 마리당 with the suffix 짜리's entry under the same letter), a unit after a run of digits or a Hangul numeral (3㎡당, 삼㎡당), and a suffix whose graph turns the piece down; of the units kept under one letter only the one whose letters fit is read (g, not kg, in g당)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestBareUnits(t *testing.T) {
	sc := scanOf("㎡당")
	per := sc.ix.after['당'][0]
	if per.name != "접미:당" || per.slot != placeOrder {
		t.Fatal(per)
	}
	got := sc.bareUnits(per, 1, 2)
	if len(got) != 1 || got[0].start != 1 || got[0].only != "㎡" || got[0].as != "단위:제곱미터" || len(got[0].pieces) != 1 || got[0].pieces[0].names != "접미:당" || got[0].pieces[0].verdict <= 0 || got[0].weight != sc.ix.afterPrefs[false] {
		t.Fatal(got)
	}
	for _, c := range []struct {
		text  string
		start int
	}{{"개당", 1}, {"마리당", 2}, {"3㎡당", 2}, {"삼㎡당", 2}} {
		sc := scanOf(c.text)
		if got := sc.bareUnits(per, c.start, c.start+1); len(got) != 0 {
			t.Fatal(c.text, got)
		}
	}
	sc = scanOf("g당")
	if got := sc.bareUnits(per, 1, 2); len(got) != 1 || got[0].only != "g" || got[0].as != "단위:그램" {
		t.Fatal("of the units under one letter only the one whose letters fit", got)
	}
	sc = scanOf("㎡짜리")
	for _, e := range sc.ix.after['리'] {
		if got := sc.bareUnits(e, 1, 3); len(got) != 0 {
			t.Fatal(e.name, got)
		}
	}
	g := toulmin.NewGraph("접미:당")
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: "접미:당", Concludes: "표지"})
	except(g, rule.PieceIs, whyNoNumber, fact(rule.PieceNum)).Attacks(claim)
	sc = scanOf("㎡당")
	if got := sc.bareUnits(&entry{name: "접미:당", form: "당", slot: placeOrder, g: g}, 1, 2); len(got) != 0 {
		t.Fatal("a piece its graph turns down", got)
	}
}
