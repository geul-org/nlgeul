//ff:func feature=rule type=rule control=sequence
//ff:what TestFormUnlisted — applies when the piece under judgment is not in the list
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestFormUnlisted(t *testing.T) {
	specs := toulmin.Specs{toulmin.NewStringListSpec("forms", []string{"맺음:지"})}
	yes, _ := FormUnlisted(pieceCtx(Piece{Form: "맺음:다"}), specs)
	no, _ := FormUnlisted(pieceCtx(Piece{Form: "맺음:지"}), specs)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
