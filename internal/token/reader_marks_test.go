//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderMarks — a line for each unit, its letters, a tab and the marks of its candidates: with the rules on 아버지가방에들어가신다 is cut as 아버지 가방에 들어가신다 with 2/49/50 at 3 and 2/41 at 6, as 아버지가 방에 … with 2/41 at 4, and as … 들어 가신다 with 2/47 at 8; with none on every mark is "-"; two units are two lines
package token

import (
	"slices"
	"strings"
	"testing"
)

func TestReaderMarks(t *testing.T) {
	on := spacingReader(t, "words-spacing.tsv", SpacingArticles...)
	lines := on.Marks("아버지가방에들어가신다")
	if len(lines) != 1 || !strings.HasPrefix(lines[0], "아버지가방에들어가신다\t") {
		t.Fatal(lines)
	}
	cands := strings.Split(strings.TrimPrefix(lines[0], "아버지가방에들어가신다\t"), ";")
	for _, want := range []string{"3:2/49/50,6:2/41,11", "4:2/41,6:2/41,11", "3:2/49/50,6:2/41,8:2/47,11"} {
		if !slices.Contains(cands, want) {
			t.Fatal(want, cands)
		}
	}
	off := spacingReader(t, "words-spacing.tsv")
	if lines := off.Marks("아버지가방에들어가신다"); len(lines) != 1 || !strings.Contains(lines[0], "\t3:-,6:-,11") || strings.Contains(lines[0], "/") {
		t.Fatal(lines)
	}
	if lines := on.Marks("아버지가 방에"); len(lines) != 2 || !strings.HasPrefix(lines[1], "방에\t") {
		t.Fatal(lines)
	}
}
