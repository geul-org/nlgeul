//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAddClosings — the closings are indexed by their last token or their lead's key; those that attach only to the copula stem are there too, and the names of those it does not take are kept; each closing of moreClosings is there under its name with the count closingCounts gives it (어야지 with its fused form, 라면서 with its form with 으), and among the forms of the tables; the closings the table already had keep their names and counts; the graph of an auxiliary particle after a linking closing takes 어서, 아서 and 면서 as well as 고 and 지, and not 다; end to end 먹되, 먹어서는, 먹고는 and 먹으면서도 read so
package token

import (
	"slices"
	"strings"
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestAddClosings(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addClosings()
	if len(ix.byKey['만']) == 0 || len(ix.byKey['ㄴ']) == 0 || len(ix.byKey[keyVowel]) == 0 {
		t.Fatal("missing keys")
	}
	var needs, holds bool
	for _, e := range ix.byKey['요'] {
		needs, holds = needs || (e.form == "에요" && e.needs), holds || (e.form == "예요" && e.holds)
	}
	if !needs || !holds || !slices.Contains(ix.barred, "꾸밈:는") || slices.Contains(ix.barred, "맺음:다") {
		t.Fatal(needs, holds, ix.barred)
	}
	for _, c := range []struct {
		key  rune
		name string
	}{
		{'되', "맺음:되"}, {'지', "맺음:어야지"}, {'지', "맺음:아야지"}, {'지', "맺음:어야지(합침)"},
		{'서', "맺음:ㄴ다면서"}, {'서', "맺음:는다면서"}, {'서', "맺음:라면서"}, {'서', "맺음:으라면서"},
		{'래', "맺음:길래"}, {'가', "맺음:든가"}, {'들', "맺음:은들"}, {'들', "맺음:ㄴ들"}, {'도', "맺음:ㄹ지라도"},
		{'고', "맺음:으라고"}, {'다', "맺음:답니다"}, {'다', "맺음:다"}, {'서', "맺음:면서"},
	} {
		if e := ix.named(c.key, c.name); e.form == "" || e.class != classClosing || e.count != closingCounts[c.name] {
			t.Errorf("%s: %+v", c.name, e)
		}
	}
	if ix.named('되', "맺음:되").count == 0 || ix.named('래', "맺음:길래").count == 0 {
		t.Fatal("the counts of the closings of this layer's own list")
	}
	if !ix.closingSet["되"] || !ix.closingSet["ㄴ다면서"] || !ix.closingSet["어야지"] || !ix.closingSet["으라고"] {
		t.Fatal("the forms of moreClosings among the closings")
	}
	for _, form := range []string{"맺음:어서", "맺음:아서", "맺음:면서", "맺음:고", "맺음:지"} {
		if v, _ := verdictOf(ix.linkAux.g, rule.Piece{Form: form, After: "보조사:는", HasStem: true, Next: -1}, ""); v <= 0 {
			t.Errorf("%s takes 는", form)
		}
	}
	if v, _ := verdictOf(ix.linkAux.g, rule.Piece{Form: "맺음:다", After: "보조사:는", HasStem: true, Next: -1}, ""); v > 0 {
		t.Fatal("다 takes no auxiliary particle")
	}
	r, _ := New(Options{Words: "testdata/words.tsv", Asleep: true})
	for in, want := range map[string]string{
		"먹되":    "먹[열쇠:동사]+되[맺음:되]",
		"먹어서는":  "먹[열쇠:동사]+어서[맺음:어서]+는[보조사:는]",
		"먹고는":   "먹[열쇠:동사]+고[맺음:고]+는[보조사:는]",
		"먹으면서도": "먹[열쇠:동사]+으면서도[맺음:으면서도]",
	} {
		if got := strings.Join(r.Lines(in, false), "\n"); !strings.Contains(got, want) {
			t.Errorf("%s: no %s\n%s", in, want, got)
		}
	}
	if got := strings.Join(r.Lines("먹으면서도", false), "\n"); strings.Contains(got, "+으면서[맺음:으면서]+도") {
		t.Fatal("으면서도 read as 으면서 and 도\n" + got)
	}
}
