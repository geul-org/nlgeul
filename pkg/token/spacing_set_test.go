//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestSpacingSet — the names of the spacing rules as a set: none gives an empty set, every name of SpacingArticles is taken, a name that is not one is an error that names it
package token

import (
	"strings"
	"testing"
)

func TestSpacingSet(t *testing.T) {
	if got, err := spacingSet(nil); err != nil || len(got) != 0 {
		t.Fatal(got, err)
	}
	got, err := spacingSet(SpacingArticles)
	if err != nil || len(got) != len(SpacingArticles) {
		t.Fatal(got, err)
	}
	for _, n := range SpacingArticles {
		if !got[n] {
			t.Fatal(n)
		}
	}
	if _, err := spacingSet([]string{"41", "99"}); err == nil || !strings.Contains(err.Error(), "99") {
		t.Fatal("an unknown name accepted", err)
	}
}
