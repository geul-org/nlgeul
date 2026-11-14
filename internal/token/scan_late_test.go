//ff:func feature=token type=reader control=sequence
//ff:what TestScanLate — the candidates kept at a place are asked for late only where a sign between letters or a cue stands and they are not worked out yet
package token

import "testing"

func TestScanLate(t *testing.T) {
	sc := &scan{kept: map[int][]*path{}, mid: map[int]midSign{}, shows: map[int]map[int]bool{}}
	if sc.late(3) {
		t.Fatal("no sign")
	}
	sc.mid[3] = midSign{}
	sc.shows[5] = map[int]bool{7: true}
	if !sc.late(3) || !sc.late(5) {
		t.Fatal("a sign or a cue")
	}
	sc.kept[3] = nil
	if sc.late(3) {
		t.Fatal("worked out")
	}
}
