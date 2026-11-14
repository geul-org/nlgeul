//ff:func feature=token type=reader control=sequence
//ff:what TestReaderBoxLines — a sentence drawn through the reader: 아버지가방에들어가신다 is one stretch of 60 ways, shown as many as head asks with the count of the rest and the brackets; every box line of a stretch ends at the same cell; with trace the lines of the unit come first; no word table gives boxes too; a sentence of spaces only gives no lines
package token

import (
	"strings"
	"testing"
)

func TestReaderBoxLines(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	s := "아버지가방에들어가신다"
	got := r.BoxLines(s, false, 3)
	if !strings.HasPrefix(got[0], "┌ [0-3] 아버지 ─ 후보경합 ") || !strings.HasSuffix(got[0], " ┬┐ ① λ 합 -18.88") {
		t.Fatal(got[0])
	}
	text := strings.Join(got, "\n")
	if strings.Count(text, "┬") != 3 || !strings.Contains(text, "(57개 생략됨)") || !strings.Contains(text, "어절경합 [0-11]") {
		t.Fatal(text)
	}
	if !strings.HasPrefix(got[len(got)-1], "1등 길: 아버지 │ 가방에 │ 들어가신다  (λ 합 ") {
		t.Fatal(got[len(got)-1])
	}
	traced := r.BoxLines(s, true, 3)
	if len(traced) <= len(got) || strings.HasPrefix(traced[0], "┌") || strings.Count(strings.Join(traced, "\n"), "┌") != strings.Count(text, "┌") {
		t.Fatal(traced[0])
	}
	plain, _ := New(Options{})
	if got := plain.BoxLines("아버지가", false, 0); len(got) < 3 || !strings.HasPrefix(got[0], "┌ [0-4] 아버지가 ─ 후보경합 ") {
		t.Fatal(got)
	}
	if got := r.BoxLines("  \t", false, 0); len(got) != 0 {
		t.Fatal(got)
	}
}
