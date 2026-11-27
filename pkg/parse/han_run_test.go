//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestHanRun — a text of Han numerals alone is a run of them (十, 二三, 二〇二四, 百萬); a text with another Han character, a Hangul numeral or a digit, and an empty text, are not
package parse

import "testing"

func TestHanRun(t *testing.T) {
	for text, want := range map[string]bool{"十": true, "二三": true, "二〇二四": true, "百萬": true, "万": true, "第一": false, "三宅": false, "삼": false, "3": false, "": false} {
		if got := hanRun(text); got != want {
			t.Errorf("hanRun(%q) = %v", text, got)
		}
	}
}
