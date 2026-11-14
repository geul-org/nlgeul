//ff:func feature=token type=reader control=sequence
//ff:what TestSpacingOf — the rules that are on and that a boundary is a place of, in the order of SpacingArticles; a rule that is off is not named, and with none on there is none
package token

import (
	"slices"
	"testing"
)

func TestSpacingOf(t *testing.T) {
	g := gap{left: spacingWord("아버지/열쇠:체언"), leftLen: 3, right: spacingWord("가방/열쇠:체언", "에/부사격:에"), rightLen: 3}
	all, _ := newIndex(Options{Spacing: SpacingArticles})
	if got := all.spacingOf(g); !slices.Equal(got, []string{"2", "49", "50"}) {
		t.Fatal(got)
	}
	some, _ := newIndex(Options{Spacing: []string{"50", "41"}})
	if got := some.spacingOf(gap{left: spacingWord("방/열쇠:체언", "에/부사격:에"), right: spacingWord("놀/열쇠:동사")}); !slices.Equal(got, []string{"41"}) {
		t.Fatal(got)
	}
	none, _ := newIndex(Options{})
	if got := none.spacingOf(g); len(got) != 0 {
		t.Fatal(got)
	}
}
