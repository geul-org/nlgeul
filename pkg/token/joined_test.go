//ff:func feature=token type=reader control=sequence
//ff:what TestJoined — each tail gets the piece on its right and keeps its start; several pieces are joined in order and the left tail's marks are kept
package token

import "testing"

func TestJoined(t *testing.T) {
	got := joined([]tail{{start: 2, base: '가', pieces: []piece{{text: "에서"}}}}, piece{text: "는"})
	if len(got) != 1 || got[0].start != 2 || got[0].base != '가' || len(got[0].pieces) != 2 || got[0].pieces[1].text != "는" || joined(nil, piece{}) != nil {
		t.Fatal(got)
	}
	many := joined([]tail{{start: 1, only: "아니", shown: "x"}}, piece{text: "a"}, piece{text: "b"})
	if len(many) != 1 || len(many[0].pieces) != 2 || many[0].only != "아니" || many[0].shown != "x" {
		t.Fatal(many)
	}
}
