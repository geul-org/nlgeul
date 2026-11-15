//ff:func feature=token type=reader control=sequence
//ff:what TestOrderFront — 제 right before a number is taken into it, named 차례 before the names it had, at a word's head, after a sign and right after a number's end (a unit after a digit, a digit); after a syllable or a Latin letter it is not (the reason is traced), nor at the head, nor a letter that is not 제
package token

import (
	"strings"
	"testing"
)

func TestOrderFront(t *testing.T) {
	num := []string{numName}
	if at, names := scanOf("제1항").orderFront(1, num); at != 0 || strings.Join(names, ",") != "차례,수" {
		t.Fatal(at, names)
	}
	if at, _ := scanOf("(제3호").orderFront(2, num); at != 1 {
		t.Fatal("after a sign", at)
	}
	if at, names := scanOf("36조제1항").orderFront(4, num); at != 3 || strings.Join(names, ",") != "차례,수" {
		t.Fatal("after a unit", at, names)
	}
	if at, _ := scanOf("조의2제3항").orderFront(4, num); at != 3 {
		t.Fatal("after a digit", at)
	}
	sc := scanOf("문제1")
	if at, names := sc.orderFront(2, num); at != 2 || strings.Join(names, ",") != "수" || !strings.Contains(strings.Join(*sc.trace, "\n"), whyOrderLeft) {
		t.Fatal("after a syllable", at, names, *sc.trace)
	}
	if at, _ := scanOf("A제1").orderFront(2, num); at != 2 {
		t.Fatal("after a Latin letter", at)
	}
	if at, _ := scanOf("1").orderFront(0, num); at != 0 {
		t.Fatal("at the head", at)
	}
	if at, _ := scanOf("가1").orderFront(1, num); at != 1 {
		t.Fatal("not 제", at)
	}
}
