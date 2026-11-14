//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestSmallEnd — where the longest run of Sino-Korean digits and 십, 백, 천 from a place ends; the place itself when there is none, and the end of the unit when the run goes to it
package token

import "testing"

func TestSmallEnd(t *testing.T) {
	sc := scanOf("오천만이십삼")
	for at, want := range map[int]int{0: 2, 2: 2, 3: 6, 6: 6} {
		if got := sc.smallEnd(at); got != want {
			t.Fatal(at, got)
		}
	}
}
