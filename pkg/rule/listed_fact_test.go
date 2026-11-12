//ff:func feature=rule type=helper control=iteration dimension=1
//ff:what TestListedFact — each List… fact is read from its own place (the piece's form, the names before and after it, its base syllable, the sign's run right of the dot and last run, the word, the tail's closing); a fact out of range falls to the closing, and a context without the place gives ""
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestListedFact(t *testing.T) {
	ctx := ctxFor(&Word{Text: "먹었다"}, Tail{Closing: "다", Size: 2})
	ctx.Set(KeyPiece, Piece{Form: "었", Before: "줄기:먹", After: "맺음:다", Base: "하"})
	ctx.Set(KeyMark, Mark{Ext: "b", Last: "kr", Parts: []string{"a", "b"}})
	want := map[int]string{
		ListForm: "었", ListBefore: "줄기:먹", ListAfter: "맺음:다", ListBase: "하",
		ListExt: "b", ListLast: "kr", ListWord: "먹었다", ListClosing: "다", 0: "다",
	}
	empty := toulmin.NewContext()
	for of, fact := range want {
		if got := listedFact(ctx, of); got != fact {
			t.Errorf("listedFact(%d) = %q, want %q", of, got, fact)
		}
		if got := listedFact(empty, of); got != "" {
			t.Errorf("listedFact(empty, %d) = %q", of, got)
		}
	}
}
