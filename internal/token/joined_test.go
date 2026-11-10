//ff:func feature=token type=reader control=sequence
//ff:what TestJoined — each tail gets the piece on its right and keeps its start
package token

import "testing"

func TestJoined(t *testing.T) {
	got := joined([]tail{{start: 2, pieces: []piece{{text: "에서"}}}}, piece{text: "는"})
	if len(got) != 1 || got[0].start != 2 || len(got[0].pieces) != 2 || got[0].pieces[1].text != "는" || joined(nil, piece{}) != nil {
		t.Fatal(got)
	}
}
