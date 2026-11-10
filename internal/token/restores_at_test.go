//ff:func feature=token type=reader control=sequence
//ff:what TestRestoresAt — with a word table ㅕ comes from ㅣ+어 and from the stem's own ㅕ, under ㅆ too; 마 is also 말+아 with a word table only (end to end: 하지 마, 마라 read with 말 when the table has it, never without a table); without one, and for another vowel, only as restores gives
package token

import (
	"strings"
	"testing"
)

func TestRestoresAt(t *testing.T) {
	ws := wordsScan("켜켰가", layerCost)
	if got := ws.restoresAt(1, leadVowel); len(got) != 2 || got[0].stem != '키' || got[1].stem != '켜' || got[1].piece != "어" {
		t.Fatal(got)
	}
	if got := ws.restoresAt(2, jongSS); len(got) != 2 || got[1].stem != '켜' || got[1].piece != "었" {
		t.Fatal(got)
	}
	if got := ws.restoresAt(3, leadVowel); len(got) != 2 || got[0].stem != '가' {
		t.Fatal("another vowel", got)
	}
	if got := scanOf("켜").restoresAt(1, leadVowel); len(got) != 1 {
		t.Fatal("no word table", got)
	}
	if got := wordsScan("마", layerCost).restoresAt(1, leadVowel); len(got) != 3 || got[0].stem != '마' || got[2].stem != '말' || got[2].piece != "아" {
		t.Fatal("마 = 말+아 with a word table", got)
	}
	if got := scanOf("마").restoresAt(1, leadVowel); len(got) != 2 || got[0].stem != '마' || got[1].stem != '므' {
		t.Fatal("마 without a word table", got)
	}
	if got := wordsScan("만", layerCost).restoresAt(1, jongN); len(got) != 1 || got[0].stem != '마' {
		t.Fatal("a lone ㄴ under 마 is no 말", got)
	}
	r, _ := New(Options{Words: "testdata/words.tsv", Asleep: true})
	r.ix.lex["말"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}}
	if got := strings.Join(r.Lines("하지 마", false), "\n"); !strings.Contains(got, "말[열쇠:동사]+(아)[맺음:어(합침)]") {
		t.Fatal("하지 마\n" + got)
	}
	if got := strings.Join(r.Lines("마라", false), "\n"); !strings.Contains(got, "말[열쇠:동사]+(아)라[맺음:어라(합침)]") {
		t.Fatal("마라\n" + got)
	}
	if got := strings.Join(Lines("하지 마 마라", false), "\n"); strings.Contains(got, "말[") || strings.Contains(got, "말+") {
		t.Fatal("no 말 without a word table\n" + got)
	}
}
