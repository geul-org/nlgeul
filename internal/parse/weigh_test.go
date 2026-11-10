//ff:func feature=parse type=rule control=sequence
//ff:what TestWeigh — verdicts of one reading: rivals by their share of the counts, a claim that does not apply is absent, the higher verdict kept, the standing particle and tail readings noted, a failing graph returned
package parse

import (
	"math"
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestWeigh(t *testing.T) {
	g, _ := newGraph()
	w := readWord("이것은", false, &sentence{})
	r := &ruling{verdicts: map[string]float64{}}
	v := r.verdicts
	if err := weigh(g, w, w.tails[0], r); err != nil {
		t.Fatal(err)
	}
	// 조사 은 816만, 꾸밈 은 869만: 2/(1+869/1685)-1 and 2/(1+816/1685)-1.
	if math.Abs(v[labelSubj]-0.3195) > 1e-3 || math.Abs(v[labelAdn]-0.3474) > 1e-3 || v[labelObj] != v[labelSubj] || len(v) != 3 || !r.chain || r.tail != "이것+은" {
		t.Fatal(v)
	}
	v[labelAdn] = 0.9
	if weigh(g, w, w.tails[0], r); v[labelAdn] != 0.9 {
		t.Fatal("a lower verdict replaced a higher one")
	}
	bad := toulmin.NewGraph("bad")
	bad.Rule(func(toulmin.Context, toulmin.Specs) (bool, any) { panic("x") })
	if weigh(bad, w, tailFact{}, r) == nil {
		t.Fatal("a panicking rule passed")
	}
}
