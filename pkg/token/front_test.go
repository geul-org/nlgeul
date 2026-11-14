//ff:func feature=token type=reader control=sequence
//ff:what TestFront — a minus or a money sign right before a number is taken into it; nothing before, another sign, and a sign after a letter are not
package token

import (
	"strings"
	"testing"
)

func TestFront(t *testing.T) {
	sc := scanOf("-3 $40 3 =3 a-3 가-3")
	if at, names := sc.front(1, []string{numName}); at != 0 || strings.Join(names, ",") != "부호,수" {
		t.Fatal(at, names)
	}
	if at, names := sc.front(4, []string{numName}); at != 3 || strings.Join(names, ",") != "돈,수" {
		t.Fatal(at, names)
	}
	if at, _ := sc.front(0, []string{numName}); at != 0 {
		t.Fatal("at the start")
	}
	if at, _ := sc.front(10, []string{numName}); at != 10 {
		t.Fatal("another sign")
	}
	if at, _ := sc.front(14, []string{numName}); at != 14 {
		t.Fatal("after a letter")
	}
	if at, _ := sc.front(18, []string{numName}); at != 18 {
		t.Fatal("after a syllable")
	}
}
