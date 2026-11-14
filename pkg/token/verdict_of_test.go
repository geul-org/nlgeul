//ff:func feature=token type=rule control=sequence
//ff:what TestVerdictOf — a marker with a stem stands (1); without one it is defeated, with the reasons when a trace is kept and with none when not
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestVerdictOf(t *testing.T) {
	ix, _ := newIndex(Options{})
	e := ix.byKey['은'][0]
	if v, why := verdictOf(e.g, rule.Piece{HasStem: true, Prev: '것', Next: -1}, "것은", true); v != 1 || why != nil {
		t.Fatal(v, why)
	}
	v, why := verdictOf(e.g, rule.Piece{Prev: '나', Next: -1}, "은", true)
	if v != -1 || len(why) != 2 || why[0] != whyNoStem {
		t.Fatal(v, why)
	}
	if v, why = verdictOf(e.g, rule.Piece{Prev: '나', Next: -1}, "은", false); v != -1 || why != nil {
		t.Fatal("no trace kept", v, why)
	}
}
