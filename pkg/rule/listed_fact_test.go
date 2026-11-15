//ff:func feature=rule type=helper control=iteration dimension=1
//ff:what TestListedFact — each List… fact is read from its own place (the piece's form, the names before and after it, its base syllable, the sign's run right of the dot and last run, the word, the tail's closing, and of the candidate judged in its sentence its stem, the stem's name and its last piece's name, the last piece of the nearest word before it and the stem of the nearest word after it); a fact out of range falls to the closing, and a context without the place gives "" (a candidate with no word before or after it gives "" for those)
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestListedFact(t *testing.T) {
	ctx := ctxFor(&Word{Text: "먹었다"}, Tail{Closing: "다", Size: 2})
	ctx.Set(KeyPiece, Piece{Form: "었", Before: "줄기:먹", After: "맺음:다", Base: "하"})
	ctx.Set(KeyMark, Mark{Ext: "b", Last: "kr", Parts: []string{"a", "b"}})
	ctx.Set(KeyAround, Around{
		Self:   Near{Stem: "학생", Kind: "열쇠:체언", Last: "보격:이"},
		Before: []Near{{Stem: "밥", Last: "목적격:을"}, {Stem: "먼", Last: "주격:이"}},
		After:  []Near{{Stem: "되", Last: "맺음:다", Pred: true}, {Stem: "먼"}},
	})
	want := map[int]string{
		ListForm: "었", ListBefore: "줄기:먹", ListAfter: "맺음:다", ListBase: "하",
		ListExt: "b", ListLast: "kr", ListWord: "먹었다", ListClosing: "다", 0: "다",
		ListSelfStem: "학생", ListSelfKind: "열쇠:체언", ListSelfLast: "보격:이", ListPrevLast: "목적격:을", ListNextStem: "되",
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
	alone := toulmin.NewContext()
	alone.Set(KeyAround, Around{Self: Near{Stem: "학생"}})
	if listedFact(alone, ListPrevLast) != "" || listedFact(alone, ListNextStem) != "" || listedFact(alone, ListSelfStem) != "학생" {
		t.Error("a candidate alone in its sentence")
	}
}
