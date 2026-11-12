//ff:func feature=rule type=rule control=iteration dimension=1
//ff:what TestListed — applies when the fact the list asks for (a piece's form, the piece before or after it, the word, the tail's closing) is in the list, and gives that fact; not when it is not, when only a part of it is, or when it is empty
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestListed(t *testing.T) {
	forms := &List{Name: "forms", Of: ListForm, Items: []string{"듯하", "만하"}}
	before := &List{Name: "before", Of: ListBefore, Items: []string{"꾸밈:는"}}
	after := &List{Name: "barred", Of: ListAfter, Items: []string{"꾸밈:는"}}
	words := &List{Name: "words", Of: ListWord, Items: []string{"그리고", "이"}}
	closings := &List{Name: "adverbial", Of: ListClosing, Items: []string{"게"}}
	cases := []struct {
		list *List
		ctx  toulmin.Context
		want float64
		fact string
	}{
		{forms, pieceCtx(Piece{Form: "만하", Adn: true}), 1, "만하"},
		{forms, pieceCtx(Piece{Form: "가르", Adn: true}), -1, "가르"},
		{forms, pieceCtx(Piece{Form: "듯"}), -1, "듯"},
		{forms, pieceCtx(Piece{}), -1, ""},
		{before, pieceCtx(Piece{Before: "꾸밈:는", After: "맺음:다"}), 1, "꾸밈:는"},
		{before, pieceCtx(Piece{After: "꾸밈:는"}), -1, ""},
		{after, pieceCtx(Piece{After: "꾸밈:는"}), 1, "꾸밈:는"},
		{after, pieceCtx(Piece{After: "맺음:다"}), -1, "맺음:다"},
		{words, ctxFor(&Word{Text: "그리고"}, Tail{}), 1, "그리고"},
		{words, ctxFor(&Word{Text: "먹고"}, Tail{}), -1, "먹고"},
		{closings, ctxFor(&Word{}, Tail{Closing: "게", Size: 2}), 1, "게"},
		{closings, ctxFor(&Word{}, Tail{Closing: "고", Size: 2}), -1, "고"},
		{closings, toulmin.NewContext(), -1, ""},
	}
	for _, c := range cases {
		got, fact := Listed(c.ctx, toulmin.Specs{c.list})
		if got != c.want || fact != c.fact {
			t.Errorf("Listed(%s) = %v, %v; want %v, %q", c.list.Name, got, fact, c.want, c.fact)
		}
	}
}
