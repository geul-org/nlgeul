//ff:func feature=token type=reader control=sequence
//ff:what TestReadUnit — a unit's bundles as lines; with a trace the judgments first; the full stop that closes the unit as a last line; a field of punctuation only comes back as it is; the readings asleep are printed only when asked for
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
	long := "아버지가방에들어가신다"
	if strings.Contains(strings.Join(readUnit(ix, long, false), "\n"), "잠듦") {
		t.Fatal("readings asleep printed unasked")
	}
	all, _ := newIndex(Options{Asleep: true})
	if !strings.Contains(strings.Join(readUnit(all, long, false), "\n"), long+"\t잠듦(") {
		t.Fatal("readings asleep not printed when asked for")
	}
}
