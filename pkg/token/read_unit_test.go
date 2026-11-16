//ff:func feature=token type=reader control=sequence
//ff:what TestReadUnit — a unit's contests as lines (contestLines: a line for each branch with its verdict and preference); with a trace the judgments first; the full stop that closes the unit as a last line; a unit of signs only comes back as the contests of its signs (...[줄임표]; ->, signs no candidate stands for, as their letters with no name); every candidate the search keeps is printed (아버지가방에들어가신다: a contest of every candidate no graph rejects, more than any fixed number kept before, the whole unit as one new word among them); a unit that begins with a sign of a pair keeps the sign as its stem, once for each way it is read; the signs before and after the letters are branches of their own; the lines of the graphs that judged the head come first with a trace
package token

import (
	"slices"
	"strings"
	"testing"
)

func TestReadUnit(t *testing.T) {
	ix, _ := newIndex(Options{})
	lines := func(ix *index, u unit, trace bool) []string {
		contests, notes := readUnit(ix, u, trace)
		return append(notes, contestLines(contests)...)
	}
	plain := lines(ix, unit{text: "아버지가"}, false)
	if len(plain) < 2 || !strings.HasPrefix(plain[0], "아버지+가[주격:가]\t") {
		t.Fatal(plain)
	}
	traced := lines(ix, unit{text: "아버지가."}, true)
	if len(traced) < 2 || !strings.HasPrefix(traced[0], "  · ") || !strings.HasPrefix(traced[len(traced)-1], ".[마침표]\t") || traced[len(traced)-2] != plain[len(plain)-1] {
		t.Fatal(traced)
	}
	if got := lines(ix, unit{text: "..."}, false); len(got) != 1 || !strings.HasPrefix(got[0], "...[줄임표]\t") {
		t.Fatal(got)
	}
	if got := lines(ix, unit{text: "->"}, false); len(got) != 1 || !strings.HasPrefix(got[0], "->\t") {
		t.Fatal("signs with no name", got)
	}
	if got := lines(ix, unit{text: "#태그"}, false); len(got) < 3 || got[0] != "#[샵]\t성립 1.00 선호 0.00" || got[1] != "#[태그]\t성립 1.00 선호 0.00" {
		t.Fatal("the signs before the letters", got)
	}
	if got := lines(ix, unit{text: "뭐?!"}, false); len(got) < 2 || !strings.HasPrefix(got[len(got)-2], "?[물음표]\t") || !strings.HasPrefix(got[len(got)-1], "![느낌표]\t") {
		t.Fatal("the signs after the letters", got)
	}
	if got := lines(ix, unit{text: "(", head: 1, name: pairName, candidates: []pairing{{openName, 1}}, notes: []string{"  · (\t열기:(\t1.00\t통과"}}, true); len(got) != 2 || got[0] != "  · (\t열기:(\t1.00\t통과" || !strings.HasPrefix(got[1], "([열기:1]\t") {
		t.Fatal("the lines of the graphs that judged the head come first with a trace", got)
	}
	if got := lines(ix, unit{text: "(", head: 1, name: pairName, candidates: []pairing{{openName, 2}}, bare: true}, false); len(got) != 2 || got[0] != "([열기:2]\t성립 1.00 선호 -5.05" || got[1] != "(\t성립 1.00 선호 -5.05" {
		t.Fatal("a head read as a pair at depth 2 and as its letter", got)
	}
	long := "아버지가방에들어가신다"
	if got := lines(ix, unit{text: "(", head: 1, name: pairName, candidates: []pairing{{openName, 1}}}, false); len(got) != 1 || !strings.HasPrefix(got[0], "([열기:1]\t") {
		t.Fatal("an opening sign alone", got)
	}
	if got := lines(ix, unit{text: long}, false); len(got) <= 16 || !slices.ContainsFunc(got, func(l string) bool {
		return strings.HasPrefix(l, "아버지+가[주격:가] 방에+들[접미:들] 어가신+다[맺음:다]\t")
	}) || !slices.ContainsFunc(got, func(l string) bool { return strings.HasPrefix(l, long+"\t") }) {
		t.Fatal("every candidate the search keeps", got)
	}
}
