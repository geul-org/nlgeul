//ff:func feature=data type=helper control=sequence
//ff:what TestDirectionAt — a pair takes two words, a single word takes one, and an unknown word reads as empty
package main

import "testing"

func TestDirectionAt(t *testing.T) {
	w := []string{"LEFT", "RIGHT", "ARROW", "UP"}
	if d, n := directionAt(w, 0); d != "양쪽" || n != 2 {
		t.Fatal(d, n)
	}
	if d, n := directionAt(w, 1); d != "오른쪽" || n != 1 {
		t.Fatal(d, n)
	}
	if d, n := directionAt(w, 2); d != "" || n != 1 {
		t.Fatal(d, n)
	}
	if d, n := directionAt(w, 3); d != "위" || n != 1 {
		t.Fatal(d, n)
	}
}
