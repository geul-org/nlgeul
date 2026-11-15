//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAuxClosingsQuote — an auxiliary particle (도, 는, 만, 야) right after a quoting closing 다고, ㄴ다고, 는다고, 라고, 자고, 냐고 or 느냐고 is one more piece of the same word (좋다고만, 가자고는, 먹느냐고도, 떠난다고는); 학생이라고는 is still read with the particle 이라고
package token

import (
	"strings"
	"testing"
)

func TestAuxClosingsQuote(t *testing.T) {
	r, _ := New(Options{Words: "testdata/words.tsv"})
	r.ix.lex["좋"] = word{kinds: kindAdj, count: 50, by: [6]int{0, 0, 50}}
	r.ix.lex["떠나"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}}
	for in, want := range map[string]string{
		"좋다고만":  "좋[열쇠:형용사]+다고[맺음:다고]+만[보조사:만]",
		"가자고는":  "가[열쇠:동사]+자고[맺음:자고]+는[보조사:는]",
		"먹느냐고도": "먹[열쇠:동사]+느냐고[맺음:느냐고]+도[보조사:도]",
		"떠난다고는": "떠나[열쇠:동사]+(ㄴ)다고[맺음:ㄴ다고]+는[보조사:는]",
		"먹는다고도": "먹[열쇠:동사]+는다고[맺음:는다고]+도[보조사:도]",
	} {
		if got := strings.Join(r.Lines(in, false), "\n"); !strings.Contains(got, want) {
			t.Errorf("%s: no %s\n%s", in, want, got)
		}
	}
	r.ix.lex["학생"] = word{kinds: kindNominal, count: 50, by: [6]int{50}}
	if got := strings.Join(r.Lines("학생이라고는", false), "\n"); !strings.Contains(got, "학생[열쇠:체언]+이라고[부사격:이라고]+는[보조사:는]") {
		t.Fatal(got)
	}
}
