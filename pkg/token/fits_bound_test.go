//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestFitsBound — article 42: the word after the boundary is a nominal whose stem is a dependent noun of the lists (the closed set, after a nominal, after a modifying closing); not a nominal of another stem, nor a verb of the same letters
package token

import "testing"

func TestFitsBound(t *testing.T) {
	for _, stem := range []string{"거/의존명사:거", "때/열쇠:체언", "중/열쇠:체언", "듯/열쇠:체언", "뚜루뚜"} {
		want := stem != "뚜루뚜"
		if got := fitsBound(gap{right: spacingWord(stem)}); got != want {
			t.Fatal(stem, got)
		}
	}
	if fitsBound(gap{right: spacingWord("거/열쇠:동사")}) || fitsBound(gap{right: spacingWord("-거/주격:가")}) {
		t.Fatal("not a nominal")
	}
}
