//ff:func feature=token type=reader control=sequence
//ff:what TestAddAlts — the changed forms of the table's predicates are kept with the stems they stand for: 사 for 살, 들 for 듣, 도우 and 도오 for 돕, 그래 for 그렇 and for 그러; a noun's letters are not changed; ㄷ and ㅅ keep the stem's own last syllable under the changed one
package token

import (
	"fmt"
	"testing"
)

func TestAddAlts(t *testing.T) {
	ix := wordsScan("x", layerCost).ix
	if fmt.Sprint(ix.alts["사"]) != fmt.Sprint([]alt{{"살", irrL}}) || fmt.Sprint(ix.alts["들"]) != fmt.Sprint([]alt{{"듣", irrD}}) {
		t.Fatal(ix.alts["사"], ix.alts["들"])
	}
	if fmt.Sprint(ix.alts["도우"]) != fmt.Sprint([]alt{{"돕", irrB}}) || fmt.Sprint(ix.alts["도오"]) != fmt.Sprint([]alt{{"돕", irrBO}}) {
		t.Fatal(ix.alts["도우"], ix.alts["도오"])
	}
	if fmt.Sprint(ix.alts["그래"]) != fmt.Sprint([]alt{{"그러", irrHE}, {"그렇", irrHE}}) {
		t.Fatal(ix.alts["그래"])
	}
	if _, ok := ix.alts["마다"]; ok || len(ix.alts["무"]) != 1 {
		t.Fatal("only a predicate's stem changes (물 the verb, not 마당 the noun)", ix.alts["무"])
	}
	if fmt.Sprint(ix.irrEnds['들']) != fmt.Sprint([]alt{{"듣", irrD}}) || fmt.Sprint(ix.irrEnds['지']) != fmt.Sprint([]alt{{"짓", irrS}}) || len(ix.irrEnds['사']) != 0 {
		t.Fatal(ix.irrEnds)
	}
}
