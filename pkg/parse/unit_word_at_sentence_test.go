//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what TestUnitWordAtSentence — end to end over a small word table that has 년간, 월경, 일전 and 점차 as words: right after a number the number and its unit come first and the word written after the unit stands as a word of its own (5년 │ 간, 6월 │ 경, 2일 │ 전, 16점 │ 차), where the number and the known word came first before; 19세기 (no word written after a unit) keeps the known word, and 년간 and 월경 with no number before them stay as they were; the known word right after the number stays a candidate
package parse

import (
	"strings"
	"testing"
)

func TestUnitWordAtSentence(t *testing.T) {
	r := streamReader(t, "words-unit-word.tsv")
	for _, c := range []struct{ in, want string }{
		{"그는 5년간 일했다.", "5[수]+년[단위:년] │ 간[의존명사:간]"},
		{"6월경 왔다.", "6[수]+월[단위:월] │ 경[열쇠:체언]"},
		{"2일전까지 왔다.", "2[수]+일[단위:일] │ 전[열쇠:체언]+까지[보조사:까지]"},
		{"16점차까지", "16[수]+점[단위:점] │ 차[열쇠:체언]+까지[보조사:까지]"},
		{"19세기의 일", "19[수] │ 세기[열쇠:체언]+의[관형격:의]"},
		{"년간 계획", "년간[열쇠:체언]"},
		{"월경 주기", "월경[열쇠:체언]"},
	} {
		var words []string
		for _, b := range bestNodes(r, c.in) {
			words = append(words, b.String())
		}
		if got := " │ " + strings.Join(words, " │ ") + " │ "; !strings.Contains(got, " │ "+c.want+" │ ") {
			t.Errorf("%s: want %s first\n%s", c.in, c.want, got)
		}
	}
	if !netHas(r, "그는 5년간 일했다.", "년간[열쇠:체언]") {
		t.Fatal("년간 right after 5 is gone")
	}
}
