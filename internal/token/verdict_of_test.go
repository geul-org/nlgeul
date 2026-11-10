//ff:func feature=token type=rule control=sequence
//ff:what TestVerdictOf — a marker with a stem stands (1); without one it is defeated with the reason
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestVerdictOf(t *testing.T) {
	ix, _ := newIndex(Options{})
	e := ix.byKey['은'][0]
	if v, why := verdictOf(e.g, rule.Piece{HasStem: true, Prev: '것', Next: -1}, "것은"); v != 1 || why != nil {
		t.Fatal(v, why)
	}
	v, why := verdictOf(e.g, rule.Piece{Prev: '나', Next: -1}, "은")
	if v != -1 || len(why) != 2 || why[0] != "앞에 줄기가 없다" {
		t.Fatal(v, why)
	}

}
