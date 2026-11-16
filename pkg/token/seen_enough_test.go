//ff:func feature=token type=rule control=sequence
//ff:what TestSeenEnough — the count seen enough is seenShare of the word table's total, so that a table twice as large asks twice as many, and at least one: with a total too small for the share to reach one count, or none, it is one
package token

import "testing"

func TestSeenEnough(t *testing.T) {
	big := (&index{lexTotal: 1e9}).seenEnough()
	if big != int(seenShare*1e9) || (&index{lexTotal: 2e9}).seenEnough() != 2*big {
		t.Fatal(big)
	}
	if (&index{lexTotal: 100}).seenEnough() != 1 || (&index{}).seenEnough() != 1 {
		t.Fatal("at least one")
	}
}
