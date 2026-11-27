//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestApproxPair — two runs of Sino-Korean numerals are an approximate number when the first ends in a digit and the second begins with the next one (육칠, 이삼십, 십오륙 with 륙 after 오); not two digits that are not neighbours, a first that ends in a greater number, 륙 after another digit, a native numeral or other letters
package parse

import "testing"

func TestApproxPair(t *testing.T) {
	for _, c := range []struct {
		first, second string
		want          bool
	}{
		{"육", "칠", true},
		{"이", "삼십", true},
		{"육", "칠만", true},
		{"십오", "륙", true},
		{"오", "육", true},
		{"일", "이", true},
		{"팔", "구", true},
		{"삼", "오", false},
		{"칠", "육", false},
		{"십", "일", false},
		{"사", "륙", false},
		{"구", "일", false},
		{"두", "세", false},
		{"육", "", false},
		{"곰", "칠", false},
	} {
		if got := approxPair(c.first, c.second); got != c.want {
			t.Errorf("approxPair(%q, %q) = %v", c.first, c.second, got)
		}
	}
}
