//ff:func feature=parse type=helper control=sequence
//ff:what TestWordOver — with a word table a run that begins before the place given and ends at the end given is found when it is a word of the table (하나 over 하 and 나), not a run that begins at that place, nor a run that is no word, and with no word table none is
package parse

import "testing"

func TestWordOver(t *testing.T) {
	sc := wordsScan("큰하나", layerCost)
	if !sc.wordOver(2, 3) {
		t.Fatal("하나 ends there and begins before the place")
	}
	if sc.wordOver(1, 3) || sc.wordOver(0, 3) || sc.wordOver(1, 2) {
		t.Fatal("no word of the table begins before the place and ends there")
	}
	if scanOf("하나").wordOver(1, 2) {
		t.Fatal("with no word table")
	}
}
