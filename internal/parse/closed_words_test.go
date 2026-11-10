//ff:func feature=parse type=reader control=sequence
//ff:what TestClosedWords — the four closed sets come out under their names
package parse

import "testing"

func TestClosedWords(t *testing.T) {
	got := ClosedWords()
	if len(got) != 4 || len(got["의존명사"]) != 9 || got["접속 부사"][0] != "그리고" || len(got["보조 용언"]) != 7 || len(got["관형사"]) != 8 {
		t.Fatal(got)
	}
}
