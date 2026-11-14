//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestJudged — judged gives what verdictOf gives and remembers it: the same question is not judged again, another word's letters are a question of their own, and when judgedMax are remembered they are all dropped, the verdict staying the same
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestJudged(t *testing.T) {
	ix, _ := newIndex(Options{})
	e := ix.byKey['은'][0]
	stands, falls := rule.Piece{HasStem: true, Prev: '것', Next: -1}, rule.Piece{Prev: '나', Next: -1}
	want, _ := verdictOf(e.g, stands, "것은", false)
	if v := ix.judged(e.g, stands, "것은"); v != want || v != 1 || len(ix.memo) != 1 {
		t.Fatal(v, want, len(ix.memo))
	}
	ix.memo[judgeKey{e.g, stands, "것은"}] = 0.5
	if v := ix.judged(e.g, stands, "것은"); v != 0.5 || len(ix.memo) != 1 {
		t.Fatal("not answered from what is remembered", v)
	}
	if v := ix.judged(e.g, stands, "곳은"); v != 1 || len(ix.memo) != 2 {
		t.Fatal("another word is asked apart", v, len(ix.memo))
	}
	if v := ix.judged(e.g, falls, "은"); v != -1 || len(ix.memo) != 3 {
		t.Fatal(v, len(ix.memo))
	}
	for i := len(ix.memo); i < judgedMax; i++ {
		ix.memo[judgeKey{facts: rule.Piece{Slot: i}}] = 1
	}
	if v := ix.judged(e.g, falls, "은"); v != -1 || len(ix.memo) != judgedMax {
		t.Fatal("a remembered verdict drops nothing", v, len(ix.memo))
	}
	if v := ix.judged(e.g, stands, "것은"); v != 0.5 || len(ix.memo) != judgedMax {
		t.Fatal(v, len(ix.memo))
	}
	if v := ix.judged(e.g, stands, "터는"); v != 1 || len(ix.memo) != 1 {
		t.Fatal("full: all dropped, then this one kept", v, len(ix.memo))
	}
	if v := ix.judged(e.g, stands, "것은"); v != want || len(ix.memo) != 2 {
		t.Fatal("asked again after the drop, the graph's own verdict", v, want)
	}
}
