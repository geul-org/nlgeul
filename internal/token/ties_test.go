//ff:func feature=token type=reader control=sequence
//ff:what TestTies — a full stop ties any two runs; a sign of numbers only two runs of digits; another sign never
package token

import "testing"

func TestTies(t *testing.T) {
	s := []rune("a.b 1:2 a-1 1-a 1=2")
	if !ties(s, [2]int{0, 1}, [2]int{2, 3}) || !ties(s, [2]int{4, 5}, [2]int{6, 7}) || ties(s, [2]int{8, 9}, [2]int{10, 11}) || ties(s, [2]int{12, 13}, [2]int{14, 15}) || ties(s, [2]int{16, 17}, [2]int{18, 19}) {
		t.Fatal("ties")
	}
}
