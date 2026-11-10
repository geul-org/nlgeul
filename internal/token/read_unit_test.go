//ff:func feature=token type=reader control=sequence
//ff:what TestReadUnit — a unit's bundles as lines; with a trace the judgments first; the full stop that closes the unit as a last line; a field of punctuation only comes back as it is
package token

import (
	"strings"
	"testing"
)

func TestReadUnit(t *testing.T) {
	ix, _ := newIndex(Options{})
	plain := readUnit(ix, "아버지가", false)
	if len(plain) < 2 || !strings.HasPrefix(plain[1], "  아버지+가[주격:가|보격:가]\t") {
		t.Fatal(plain)
	}
	traced := readUnit(ix, "아버지가.", true)
	if len(traced) < 2 || !strings.HasPrefix(traced[0], "  · ") || traced[len(traced)-1] != ".[마침표:.]" || traced[len(traced)-2] != plain[len(plain)-1] {
		t.Fatal(traced)
	}
	if got := readUnit(ix, "...", false); len(got) != 1 || got[0] != "..." {
		t.Fatal(got)
	}
}
