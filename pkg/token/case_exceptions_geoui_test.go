//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestCaseExceptionsGeoui — the row of 거의 in the case exception table held in the program (graph 관형격:의, ground observed): end to end, over the small word table with the adverb 거의 and the noun 과거 added, 거의 alone and after a modifying form is read first as the adverb and is no longer read as 거 with the adnominal particle 의; a unit whose letters are other than 거의 keeps the particle (과거의)
package token

import (
	"slices"
	"strings"
	"testing"
)

func TestCaseExceptionsGeoui(t *testing.T) {
	rows := caseExceptions["관형격:의"]
	if len(rows) != 1 || rows[0].Word != "거의" || rows[0].Reading != "거의[열쇠:부사]" || rows[0].GroundKind != groundSeen {
		t.Fatal(rows)
	}
	r := streamReader(t, "words.tsv")
	r.ix.lex["거의"] = word{kinds: kindAdverb, count: 9000, by: [6]int{0, 0, 0, 9000}}
	r.ix.lex["과거"] = word{kinds: kindNominal, count: 9000, by: [6]int{9000}}
	for _, s := range []string{"거의", "먹는 거의 모든 것", "과거의 일"} {
		var words []string
		for _, b := range bestNodes(r, s) {
			words = append(words, b.String())
		}
		got := " │ " + strings.Join(words, " │ ") + " │ "
		want := " │ 거의[열쇠:부사] │ "
		if s == "과거의 일" {
			want = " │ 과거[열쇠:체언]+의[관형격:의] │ "
		}
		if !strings.Contains(got, want) {
			t.Errorf("%s: want %s\n%s", s, want, got)
		}
		if s != "과거의 일" && netHas(r, s, "거[의존명사:거]+의[관형격:의]") {
			t.Errorf("%s: 거+의 is still a candidate", s)
		}
	}
	if !slices.ContainsFunc(caseExceptions["맺음:소"], func(c caseException) bool { return c.Word == "담소" }) {
		t.Fatal("the rows before it stay")
	}
}
