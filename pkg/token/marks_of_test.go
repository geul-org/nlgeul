//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestMarksOf — each candidate as its word ends joined with ",", the ends inside the unit followed by ":" and the rules that fit joined with "/" ("-" for none), the candidates joined with ";"; the last end has no mark, and no candidate is an empty text
package token

import (
	"slices"
	"strings"
	"testing"
)

func TestMarksOf(t *testing.T) {
	sc := spacingScan("아버지가방에들어가신다", SpacingArticles...)
	got := strings.Split(sc.marksOf(sc.search()), ";")
	if len(got) < 8 || !slices.Contains(got, "3:2/49/50,6:2/41,11") || !slices.Contains(got, "4:2/41,6:2/41,11") || !slices.Contains(got, "3:2/49/50,6:2/41,8:2/47,11") {
		t.Fatal(len(got), got[:4])
	}
	for _, c := range got {
		if !strings.HasSuffix(c, ",11") && c != "11" {
			t.Fatal("the last end", c)
		}
	}
	none := spacingScan("아버지가방에들어가신다")
	if first := strings.Split(none.marksOf(none.search()), ";")[0]; first != "3:-,6:-,11" {
		t.Fatal(first)
	}
	if got := none.marksOf(nil); got != "" {
		t.Fatal(got)
	}
}
