//ff:func feature=token type=rule control=sequence
//ff:what TestVerdict — the verdict of one graph for one piece, the same value verdictOf gives with no word and no trace: a marker with a stem stands (1), one without is defeated (-1); the word under judgment is empty; a graph that fails or claims nothing gives -1
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestVerdict(t *testing.T) {
	ix, _ := newIndex(Options{})
	e := ix.byKey['은'][0]
	stem := rule.Piece{HasStem: true, Prev: '것', Next: -1}
	bare := rule.Piece{Prev: '나', Next: -1}
	if v, _ := verdictOf(e.g, stem, "", false); verdict(e.g, stem) != 1 || v != 1 {
		t.Fatal("a stem before it", verdict(e.g, stem), v)
	}
	if v, _ := verdictOf(e.g, bare, "", false); verdict(e.g, bare) != -1 || v != -1 {
		t.Fatal("no stem before it", verdict(e.g, bare), v)
	}
	word := toulmin.NewGraph("word")
	word.Rule(rule.Listed).With(listed("empty", rule.ListWord, []string{""}))
	if v := verdict(word, stem); v != 1 {
		t.Fatal("the word under judgment is empty", v)
	}
	bad := toulmin.NewGraph("bad")
	bad.Rule(func(toulmin.Context, toulmin.Specs) (bool, any) { panic("x") })
	if v := verdict(bad, stem); v != -1 {
		t.Fatal("a graph that fails", v)
	}
	if v := verdict(toulmin.NewGraph("empty"), stem); v != -1 {
		t.Fatal("a graph that claims nothing", v)
	}
}
