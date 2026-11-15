//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestRyeoneunClosing — the modifying closings 려는 and 려던 (with 으 after a stem with a final consonant, which 짓 keeps as its own) are read after a verb stem as one word, and a noun that ends in 려 and takes 는 keeps its reading as a noun
package token

import (
	"strings"
	"testing"
)

func TestRyeoneunClosing(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-ryeoneun.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct{ sentence, word, want string }{
		{"그는 새 집을 지으려는 계획을 세웠다.", "지으려는", "짓[열쇠:동사]+으려는[꾸밈:으려는]"},
		{"서류를 받으려는 사람은 신청서를 낸다.", "받으려는", "받[열쇠:동사]+으려는[꾸밈:으려는]"},
		{"내가 하려던 말은 그것이다.", "하려던", "하[열쇠:동사]+려던[꾸밈:려던]"},
		{"떠나려는 배를 붙잡았다.", "떠나려는", "떠나[열쇠:동사]+려는[꾸밈:려는]"},
		{"그 우려는 지나쳤다.", "우려는", "우려[열쇠:체언]+는[보조사:는]"},
		{"배려는 고마운 일이다.", "배려는", "배려[열쇠:체언]+는[보조사:는]"},
	} {
		got := "\n" + strings.Join(r.Lines(c.sentence, false), "\n")
		_, rest, ok := strings.Cut(got, "\n"+c.word+"\t경합")
		if !ok {
			t.Errorf("%s: no contest for %s", c.sentence, c.word)
			continue
		}
		_, first, _ := strings.Cut(rest, "\n")
		if !strings.HasPrefix(strings.TrimSpace(first), c.want+"\t") {
			t.Errorf("%s: first is %q, want %s", c.sentence, first, c.want)
		}
	}
}
