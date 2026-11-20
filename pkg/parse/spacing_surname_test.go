//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestSpacingSurname — article 48 with every rule on, over a small table with the counts of the big one (words-surname.tsv): a family name written onto a title is read as the family name and the title apart, first on the best way (오과장과의 회의, 김군은 왔다, 고씨는 말했다, 송씨에게 말했다, 유교수가 왔다, 오기자가 왔다, 오과장이다, 김선생님이 왔다); 남궁씨가 왔다 is cut 남궁 │ 씨+가 and 남궁[성:남궁] is a candidate of 남궁 (the table's noun 남궁 stands first)
package parse

import (
	"slices"
	"testing"
)

func TestSpacingSurname(t *testing.T) {
	r := spacingReader(t, "words-surname.tsv", SpacingArticles...)
	for _, c := range []struct {
		text  string
		first []string
	}{
		{"오과장과의 회의", []string{"{0 1} 오[성:오]", "{1 5} 과장[열쇠:체언]+과[부사격:과]+의[관형격:의]"}},
		{"김군은 왔다", []string{"{0 1} 김[성:김]", "{1 3} 군[열쇠:체언]+은[보조사:은]"}},
		{"고씨는 말했다", []string{"{0 1} 고[성:고]", "{1 3} 씨[열쇠:체언]+는[보조사:는]"}},
		{"송씨에게 말했다", []string{"{0 1} 송[성:송]", "{1 4} 씨[열쇠:체언]+에게[부사격:에게]"}},
		{"유교수가 왔다", []string{"{0 1} 유[성:유]", "{1 4} 교수[열쇠:체언]+가[주격:가]"}},
		{"오기자가 왔다", []string{"{0 1} 오[성:오]", "{1 4} 기자[열쇠:체언]+가[주격:가]"}},
		{"오과장이다", []string{"{0 1} 오[성:오]", "{1 5} 과장[열쇠:체언]+이[서술격:이]+다[맺음:다]"}},
		{"김선생님이 왔다", []string{"{0 1} 김[성:김]", "{1 5} 선생님[열쇠:체언]+이[주격:이]"}},
	} {
		if got := flatNodes(r, c.text, true); len(got) < 2 || !slices.Equal(got[:2], c.first) {
			t.Error(c.text, got)
		}
	}
	got := flatNodes(r, "남궁씨가 왔다", true)
	if len(got) < 2 || got[0] != "{0 2} 남궁[열쇠:체언]" || got[1] != "{2 4} 씨[열쇠:체언]+가[주격:가]" || !netHas(r, "남궁씨가 왔다", "남궁[성:남궁]") {
		t.Error("남궁씨가", got)
	}
}
