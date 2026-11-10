//ff:func feature=rule type=rule control=sequence
//ff:what TestNextUnlisted — applies when the piece that follows is not in the list
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestNextUnlisted(t *testing.T) {
	specs := toulmin.Specs{toulmin.NewStringListSpec("takers", []string{"맺음:니"})}
	yes, _ := NextUnlisted(pieceCtx(Piece{After: "맺음:고"}), specs)
	no, _ := NextUnlisted(pieceCtx(Piece{After: "맺음:니"}), specs)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
