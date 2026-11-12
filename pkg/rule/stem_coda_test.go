//ff:func feature=rule type=rule control=sequence
//ff:what TestStemCoda — applies when markers follow a stem whose last syllable ends in the criterion's final consonant (ㅆ); not with no tail or a bare tail, another final consonant, no final consonant, an empty form or a last letter that is not Hangul; the form is given back
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestStemCoda(t *testing.T) {
	specs := toulmin.Specs{&Bits{Of: "jong", Mask: JongSS}}
	yes, form := StemCoda(pieceCtx(Piece{Form: "나갔", Tail: TailPred}), specs)
	one, _ := StemCoda(pieceCtx(Piece{Form: "갔", Tail: TailCase}), specs)
	both, _ := StemCoda(pieceCtx(Piece{Form: "나갔", Tail: TailBare | TailPred}), specs)
	if yes != 1 || one != 1 || both != 1 || form != "나갔" {
		t.Fatal(yes, one, both, form)
	}
	noTail, _ := StemCoda(pieceCtx(Piece{Form: "나갔"}), specs)
	bare, _ := StemCoda(pieceCtx(Piece{Form: "나갔", Tail: TailBare}), specs)
	if noTail != -1 || bare != -1 {
		t.Fatal("no markers after the stem", noTail, bare)
	}
	other, _ := StemCoda(pieceCtx(Piece{Form: "나간", Tail: TailPred}), specs)
	open, _ := StemCoda(pieceCtx(Piece{Form: "나가", Tail: TailPred}), specs)
	first, _ := StemCoda(pieceCtx(Piece{Form: "갔나", Tail: TailPred}), specs)
	if other != -1 || open != -1 || first != -1 {
		t.Fatal("the last syllable does not end in ㅆ", other, open, first)
	}
	empty, ef := StemCoda(pieceCtx(Piece{Tail: TailPred}), specs)
	latin, lf := StemCoda(pieceCtx(Piece{Form: "갔x", Tail: TailPred}), specs)
	if empty != -1 || latin != -1 || ef != "" || lf != "갔x" {
		t.Fatal(empty, latin, ef, lf)
	}
	n, _ := StemCoda(pieceCtx(Piece{Form: "나간", Tail: TailPred}), toulmin.Specs{&Bits{Of: "jong", Mask: JongN}})
	if n != 1 {
		t.Fatal("the criterion names the final consonant", n)
	}
}
