//ff:func feature=rule type=rule control=sequence
//ff:what TestNextListed — applies when the piece that follows is in the list
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestNextListed(t *testing.T) {
	specs := toulmin.Specs{toulmin.NewStringListSpec("barred", []string{"꾸밈:는"})}
	in, _ := NextListed(pieceCtx(Piece{After: "꾸밈:는"}), specs)
	out, _ := NextListed(pieceCtx(Piece{After: "맺음:다"}), specs)
	if in != 1 || out != -1 {
		t.Fatal(in, out)
	}
}
