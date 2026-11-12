//ff:func feature=token type=reader control=sequence
//ff:what TestEntryPiece — the piece of an entry has the letters as they stand, the entry's name, the verdict given, its kind of tail and the preference of its count and form; the piece of an entry that reads a closing written wrong has the right closing as its text, the letters it is written with, and the preference of the right closing's form (the length of the form when it has no count)
package token

import (
	"math"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestEntryPiece(t *testing.T) {
	sc := scanOf("x")
	e := &entry{name: "맺음:고", form: "고", class: classClosing, count: 500}
	got := sc.entryPiece(e, "고", 0.8)
	if got.text != "고" || got.names != "맺음:고" || got.verdict != 0.8 || got.tail != rule.TailPred || got.unseen || got.wrote != "" || got.pref != sc.ix.pref(500, "고") {
		t.Fatal(got)
	}
	wrong := &entry{name: "맺음:습니다", form: "읍니다", class: classClosing, count: 0, wrote: "읍니다", right: "습니다"}
	got = sc.entryPiece(wrong, "읍니다", 0.6)
	if got.text != "습니다" || got.wrote != "읍니다" || got.names != "맺음:습니다" || got.verdict != 0.6 || math.Abs(got.pref-3*math.Log10(1.0/syllables)) > 1e-9 {
		t.Fatal(got, 3*math.Log10(1.0/syllables))
	}
	jong := &entry{name: "맺음:ㄹ게", form: "께", class: classClosing, lead: rule.JongL, count: 900, wrote: "ㄹ께", right: "(ㄹ)게"}
	if got = sc.entryPiece(jong, "께", 1); got.text != "(ㄹ)게" || got.wrote != "ㄹ께" || got.pref != sc.ix.pref(900, "게") {
		t.Fatal(got)
	}
}
