//ff:func feature=rule type=helper control=sequence
//ff:what TestHarmony — 아 after ㅏ ㅑ ㅗ, 어 after the others and after ㅆ, 여 only after 하; a letter that is no Hangul syllable, and a piece that begins with none of these, fit anything
package rule

import "testing"

func TestHarmony(t *testing.T) {
	if !Harmony('잡', '아') || Harmony('잡', '어') || !Harmony('먹', '어') || Harmony('먹', '아') || !Harmony('보', '았') || !Harmony('얇', '아') {
		t.Fatal("ㅏ ㅑ ㅗ take 아, the others 어")
	}
	if !Harmony('갔', '어') || Harmony('갔', '아') || !Harmony('겠', '어') {
		t.Fatal("after ㅆ always 어")
	}
	if !Harmony('하', '여') || !Harmony('하', '였') || Harmony('하', '아') || Harmony('되', '여') || Harmony('가', '였') {
		t.Fatal("여 only after 하")
	}
	if !Harmony('x', '아') || !Harmony('먹', '고') {
		t.Fatal("no Hangul syllable before, or no such piece")
	}
}
