//ff:func feature=rule type=rule control=sequence
//ff:what TestFormListed — applies when the form of the piece under judgment is in the list, not when it is not or is empty, and gives the form back
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestFormListed(t *testing.T) {
	specs := toulmin.Specs{toulmin.NewStringListSpec("forms", []string{"듯하", "만하"})}
	yes, yf := FormListed(pieceCtx(Piece{Form: "만하", Adn: true}), specs)
	no, nf := FormListed(pieceCtx(Piece{Form: "가르", Adn: true}), specs)
	part, _ := FormListed(pieceCtx(Piece{Form: "듯"}), specs)
	empty, ef := FormListed(pieceCtx(Piece{}), specs)
	if yes != 1 || no != -1 || part != -1 || empty != -1 {
		t.Fatal(yes, no, part, empty)
	}
	if yf != "만하" || nf != "가르" || ef != "" {
		t.Fatal(yf, nf, ef)
	}
}
