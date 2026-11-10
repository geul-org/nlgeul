//ff:func feature=mini type=rule control=sequence
//ff:what TestClassify — a failing graph stops the sentence; an empty sentence has no lines; words are printed as typed, punctuation kept around the divided word
package mini

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestClassify(t *testing.T) {
	bad := toulmin.NewGraph("bad")
	bad.Rule(func(toulmin.Context, toulmin.Specs) (bool, any) { panic("x") })
	if _, err := classify(bad, "이것은 문장이다"); err == nil {
		t.Fatal("a panicking rule passed")
	}
	g, _ := newGraph()
	if out, err := classify(g, "  "); err != nil || out != nil {
		t.Fatal(out, err)
	}
	if out, _ := classify(g, "그것이 \"문제였다.\""); out[1] != "\"문제+였+다.\":서술어" {
		t.Fatal(out)
	}
	if out, _ := classify(g, "\"밥을 먹자.\""); len(out) != 2 || out[0] != "\"밥+을:목적어|관형어" || out[1] != "먹+자.\":서술어" {
		t.Fatal(out)
	}
}
