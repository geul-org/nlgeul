//ff:func feature=token type=rule control=sequence
//ff:what TestListed — the criterion of a named list over one fact with its items; rule.Listed under it holds when that fact of the piece is among the items and not when it is not (turned round with not), and gives the fact it looked up; a list without a name or a fact does not validate
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestListed(t *testing.T) {
	items := []string{"맺음:지", "맺음:고"}
	l := listed("linking", rule.ListBefore, items)
	if l.Name != "linking" || l.Of != rule.ListBefore || !slices.Equal(l.Items, items) || l.SpecName() != "list:linking" || l.Validate() != nil {
		t.Fatal(l)
	}
	ctx := toulmin.NewContext()
	ctx.Set(rule.KeyPiece, rule.Piece{Form: "는", Before: "맺음:고"})
	if v, got := rule.Listed(ctx, toulmin.Specs{l}); v != 1 || got != "맺음:고" {
		t.Fatal("the piece before it is listed", v, got)
	}
	if v, got := rule.Listed(ctx, toulmin.Specs{listed("forms", rule.ListForm, items)}); v != -1 || got != "는" {
		t.Fatal("the piece's form is not listed", v, got)
	}
	if v, _ := rule.Listed(ctx, toulmin.Specs{l, not}); v != -1 {
		t.Fatal("turned round with not", v)
	}
	if listed("", rule.ListForm, items).Validate() == nil || listed("none", 0, items).Validate() == nil {
		t.Fatal("a list without a name or a fact")
	}
}
