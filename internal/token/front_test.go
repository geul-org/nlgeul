//ff:func feature=token type=reader control=sequence
//ff:what TestFront — a minus or a money sign right before a number is taken into it; nothing before, another sign, and a sign after a letter are not
package token

import "testing"

func TestFront(t *testing.T) {
	sc := scanOf("-3 $40 3 =3 a-3 가-3")
	if at, names := sc.front(1, numName); at != 0 || names != "부호|수" {
		t.Fatal(at, names)
	}
	if at, names := sc.front(4, numName); at != 3 || names != "돈|수" {
		t.Fatal(at, names)
	}
	if at, _ := sc.front(0, numName); at != 0 {
		t.Fatal("at the start")
	}
	if at, _ := sc.front(10, numName); at != 10 {
		t.Fatal("another sign")
	}
	if at, _ := sc.front(14, numName); at != 14 {
		t.Fatal("after a letter")
	}
	if at, _ := sc.front(18, numName); at != 18 {
		t.Fatal("after a syllable")
	}
}
