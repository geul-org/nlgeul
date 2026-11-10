//ff:func feature=rule type=rule control=sequence
//ff:what TestShortKey — applies to a key of one syllable; not to a longer one, not to an empty one
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestShortKey(t *testing.T) {
	specs := toulmin.Specs{&Bits{Of: "syllables", Mask: 1}}
	short, _ := ShortKey(pieceCtx(Piece{Form: "마"}), specs)
	long, _ := ShortKey(pieceCtx(Piece{Form: "마시"}), specs)
	none, _ := ShortKey(pieceCtx(Piece{}), specs)
	if short != 1 || long != -1 || none != -1 {
		t.Fatal(short, long, none)
	}
}
