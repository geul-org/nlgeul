//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestTitleAfter — a title or an office of the list from the place given counts when the unit ends right after it (김과장), a particle follows it (김군은, 김선생님이: 선생님 and then 이) or a letter that is no Hangul syllable does (김씨,); not when other letters follow it (김양식장: 양 and then 식), when none stands there (김밥은) or the place is the unit's end
package token

import "testing"

func TestTitleAfter(t *testing.T) {
	for _, c := range []struct {
		text string
		want bool
	}{
		{"김과장", true},
		{"김군은", true},
		{"김선생님이", true},
		{"김씨,", true},
		{"김양식장", false},
		{"김밥은", false},
		{"김", false},
	} {
		if got := spacingScan(c.text).titleAfter(1); got != c.want {
			t.Fatal(c.text, got)
		}
	}
}
