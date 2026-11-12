//ff:func feature=rule type=rule control=iteration dimension=1
//ff:what TestListedNot — with Not it applies when the fact is not in the list (a piece after that takes no 으, a form out of the list, a fused piece under another syllable, an extension or a top-level domain not listed), and not when it is; it still gives the fact
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestListedNot(t *testing.T) {
	takers := &List{Name: "takers", Of: ListAfter, Items: []string{"맺음:니"}}
	forms := &List{Name: "forms", Of: ListForm, Items: []string{"맺음:지"}}
	bases := &List{Name: "bases", Of: ListBase, Items: []string{"하", "되"}}
	ext := &List{Name: "ext", Of: ListExt, Items: []string{"py", "md"}}
	tld := &List{Name: "tld", Of: ListLast, Items: []string{"com", "kr"}}
	cases := []struct {
		list *List
		ctx  toulmin.Context
		want float64
		fact string
	}{
		{takers, pieceCtx(Piece{After: "맺음:고"}), 1, "맺음:고"},
		{takers, pieceCtx(Piece{After: "맺음:니"}), -1, "맺음:니"},
		{forms, pieceCtx(Piece{Form: "맺음:다"}), 1, "맺음:다"},
		{forms, pieceCtx(Piece{Form: "맺음:지"}), -1, "맺음:지"},
		{bases, pieceCtx(Piece{Base: "크"}), 1, "크"},
		{bases, pieceCtx(Piece{Base: "하"}), -1, "하"},
		{ext, markCtx(Mark{Ext: "exe", Last: "exe"}), 1, "exe"},
		{ext, markCtx(Mark{Ext: "md", Last: "md"}), -1, "md"},
		{tld, markCtx(Mark{Ext: "py", Last: "py"}), 1, "py"},
		{tld, markCtx(Mark{Ext: "co", Last: "kr"}), -1, "kr"},
		{tld, toulmin.NewContext(), 1, ""},
	}
	for _, c := range cases {
		got, fact := Listed(c.ctx, toulmin.Specs{c.list, &Not{}, &Claim{Name: "a", Concludes: "b"}})
		if got != c.want || fact != c.fact {
			t.Errorf("Listed(%s, not) = %v, %v; want %v, %q", c.list.Name, got, fact, c.want, c.fact)
		}
	}
}
