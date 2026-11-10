//ff:func feature=mini type=rule control=sequence
//ff:what TestMark — a standing particle reading divides before every particle of the chain; else the standing tail reading's division; else the word is whole
package mini

import "testing"

func TestMark(t *testing.T) {
	w := readWord("그것만으로도", false, &sentence{})
	if got := mark(w, &ruling{chain: true, tail: "x"}); got != "그것+만+으로+도" {
		t.Fatal(got)
	}
	if got := mark(w, &ruling{tail: "문제+였+다"}); got != "문제+였+다" {
		t.Fatal(got)
	}
	if got := mark(w, &ruling{}); got != "그것만으로도" {
		t.Fatal(got)
	}
}
