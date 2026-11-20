//ff:func feature=parse type=rule control=sequence
//ff:what TestTableBound — the words boundTableRule asks are the dependent nouns of the table that stand after a modifying closing and bear the word table's name (리, 바, 지, 터, 자), not those of the closed set (수, 것, 때) nor those that stand after a nominal (측, 중)
package parse

import (
	"slices"
	"testing"
)

func TestTableBound(t *testing.T) {
	in := func(w string) bool { return slices.Contains(tableBound, w) }
	if !in("리") || !in("바") || !in("지") || !in("터") || !in("자") {
		t.Errorf("missing a word of the table: %v", tableBound)
	}
	if in("수") || in("것") || in("때") || in("측") || in("중") {
		t.Errorf("a word of the closed set or after a nominal: %v", tableBound)
	}
}
