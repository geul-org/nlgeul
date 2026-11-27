//ff:func feature=parse type=reader control=sequence
//ff:what TestBehindPref — the preference of a candidate that stands behind the words of the word table is the least the table's kinds give its letters (the kind seen least), a new word's value when the table lacks the letters or gives them no count, and never above a new word's
package parse

import (
	"math"
	"testing"
)

func TestBehindPref(t *testing.T) {
	ix := wordsScan("", layerCost).ix
	fresh := math.Log10(1.0 / syllables)
	if got := ix.behindPref("뚜루", 2); got != 2*fresh {
		t.Fatal("letters the table lacks", got)
	}
	ix.lexTotal = 1e9
	ix.lex["자"] = word{kinds: kindNominal | kindAlone, count: 9200, by: [6]int{9000, 0, 0, 0, 0, 200}}
	if got := ix.behindPref("자", 1); !almost(got, math.Log10(200/ix.lexTotal)) || got >= math.Log10(9000/ix.lexTotal) {
		t.Fatal("the kind seen least", got)
	}
	ix.lex["곰"] = word{kinds: kindNominal, count: 5e8, by: [6]int{5e8}}
	if got := ix.behindPref("곰", 1); got != fresh {
		t.Fatal("never above a new word's", got)
	}
	ix.lex["솔"] = word{kinds: kindNominal}
	if got := ix.behindPref("솔", 1); got != fresh {
		t.Fatal("a kind with no count", got)
	}
}
