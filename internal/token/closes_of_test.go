//ff:func feature=token type=helper control=sequence
//ff:what TestClosesOf — a tail with no closing entry gives 0; a piece that is not its closing (a particle, a slot) gives 0; the closing piece gives what its closing is marked with: final, link, both, or none
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestClosesOf(t *testing.T) {
	slot := piece{text: "었", names: "선어말:었"}
	closing := piece{text: "다", names: "맺음:다"}
	if got := closesOf(tail{start: 1, pieces: []piece{slot, closing}}, closing); got != 0 {
		t.Fatalf("a tail with no closing entry gives %d", got)
	}

	final := tail{start: 1, pieces: []piece{slot, closing}, close: &entry{name: "맺음:다", form: "다", ends: rule.KindFinal}}
	if got := closesOf(final, closing); got != rule.KindFinal {
		t.Fatalf("a final closing gives %d, want %d", got, rule.KindFinal)
	}
	if got := closesOf(final, slot); got != 0 {
		t.Fatalf("a slot gives %d", got)
	}
	if got := closesOf(final, piece{text: "가", names: "주격:가|보격:가"}); got != 0 {
		t.Fatalf("a particle gives %d", got)
	}

	linking := piece{text: "고", names: "맺음:고"}
	link := tail{start: 1, pieces: []piece{linking}, close: &entry{name: "맺음:고", form: "고", ends: rule.KindLink}}
	if got := closesOf(link, linking); got != rule.KindLink {
		t.Fatalf("a linking closing gives %d, want %d", got, rule.KindLink)
	}
	either := piece{text: "어", names: "맺음:어"}
	both := tail{start: 1, pieces: []piece{either}, close: &entry{name: "맺음:어", form: "어", ends: rule.KindFinal | rule.KindLink}}
	if got := closesOf(both, either); got != 3 {
		t.Fatalf("a closing that ends and links gives %d, want 3", got)
	}
	modifying := piece{text: "는", names: "꾸밈:는"}
	unmarked := piece{text: "네", names: "맺음:네"}
	none := tail{start: 1, pieces: []piece{unmarked}, close: &entry{name: "맺음:네", form: "네"}}
	if got := closesOf(none, unmarked); got != 0 {
		t.Fatalf("a closing the tables do not mark gives %d", got)
	}
	if got := closesOf(tail{start: 1, pieces: []piece{modifying}, close: &entry{name: "꾸밈:는", form: "는", ends: rule.KindLink}}, modifying); got != 0 {
		t.Fatalf("a piece not named 맺음 gives %d", got)
	}
}
