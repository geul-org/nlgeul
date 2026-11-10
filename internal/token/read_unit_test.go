//ff:func feature=token type=reader control=sequence
//ff:what TestReadUnit — a unit's bundles as lines; with a trace the judgments first; the full stop that closes the unit as a last line; a field of punctuation only comes back as it is; the readings asleep are printed only when asked for; a unit that begins with a sign of a pair keeps the sign as its stem
package token

import (
	"strings"
	"testing"
)

func TestReadUnit(t *testing.T) {
	ix, _ := newIndex(Options{})
	lines := func(ix *index, u unit, trace bool) []string {
		bundles, notes := readUnit(ix, u, trace)
		for _, b := range bundles {
			if ix.asleep || !b.Asleep() {
				notes = append(notes, render(b)...)
			}
		}
		return notes
	}
	plain := lines(ix, unit{text: "아버지가"}, false)
	if len(plain) < 2 || !strings.HasPrefix(plain[1], "  아버지+가[주격:가|보격:가]\t") {
		t.Fatal(plain)
	}
	traced := lines(ix, unit{text: "아버지가."}, true)
	if len(traced) < 2 || !strings.HasPrefix(traced[0], "  · ") || traced[len(traced)-1] != ".[마침표:.]" || traced[len(traced)-2] != plain[len(plain)-1] {
		t.Fatal(traced)
	}
	if got := lines(ix, unit{text: "..."}, false); len(got) != 1 || got[0] != "..." {
		t.Fatal(got)
	}
	long := "아버지가방에들어가신다"
	if got := lines(ix, unit{text: ")를", head: 1, name: "닫는 묶음표"}, false); len(got) < 2 || !strings.HasPrefix(got[1], "  )[닫는 묶음표]+를[목적격:를]\t") {
		t.Fatal("a particle written onto a closing sign", got)
	}
	if got := lines(ix, unit{text: "(", head: 1, name: "여는 묶음표"}, false); len(got) != 1 || got[0] != "([여는 묶음표]" {
		t.Fatal("an opening sign alone", got)
	}
	if strings.Contains(strings.Join(lines(ix, unit{text: long}, false), "\n"), "잠듦") {
		t.Fatal("readings asleep printed unasked")
	}
	all, _ := newIndex(Options{Asleep: true})
	if !strings.Contains(strings.Join(lines(all, unit{text: long}, false), "\n"), long+"\t잠듦(") {
		t.Fatal("readings asleep not printed when asked for")
	}
}
