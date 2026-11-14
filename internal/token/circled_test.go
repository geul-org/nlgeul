//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestCircled — 1 to 20 are ① to ⑳, 21 to 35 ㉑ to ㉟, 36 to 50 ㊱ to ㊿ (at both ends of each run), and any other number is in round brackets
package token

import "testing"

func TestCircled(t *testing.T) {
	for n, want := range map[int]string{
		1: "①", 2: "②", 20: "⑳", 21: "㉑", 35: "㉟", 36: "㊱", 50: "㊿", 51: "(51)", 0: "(0)", 120: "(120)",
	} {
		if got := circled(n); got != want {
			t.Fatalf("%d: %q, want %q", n, got, want)
		}
	}
}
