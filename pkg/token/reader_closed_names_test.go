//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderClosedNames — a closed-set word the word table lists as a kind is read under the set's name and under the table's kind (감자 수가 많다: 수[의존명사:수] and 수[열쇠:체언]; 집에 있다: 있[보조 용언:있], 있[열쇠:동사] and 있[열쇠:형용사]; 영화를 보기에: 보[보조 용언:보] and 보[열쇠:동사]) with the same verdict and preference for one kind; a closed-set word the table does not list gets the set's name alone (놓다)
package token

import (
	"slices"
	"strings"
	"testing"
)

func TestReaderClosedNames(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-closed.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for text, want := range map[string][]string{
		"감자 수가 많다": {"수[의존명사:수]+가[주격:가]", "수[열쇠:체언]+가[주격:가]", "수[의존명사:수]+가[보격:가]", "수[열쇠:체언]+가[보격:가]"},
		"집에 있다":    {"있[보조 용언:있]+다[맺음:다]", "있[열쇠:동사]+다[맺음:다]", "있[열쇠:형용사]+다[맺음:다]"},
		"영화를 보기에":  {"보[보조 용언:보]+기[이름:기]+에[부사격:에]", "보[열쇠:동사]+기[이름:기]+에[부사격:에]"},
	} {
		lines := strings.Join(r.Lines(text, false), "\n")
		if bad := orderedLines(lines, want); len(bad) > 0 {
			t.Errorf("%q: %q is missing or out of order in\n%s", text, bad, lines)
		}
	}
	lines := r.Lines("집에 있다", false)
	pref := map[string]string{}
	for _, l := range lines {
		if name, score, ok := strings.Cut(strings.TrimSpace(l), "\t"); ok && strings.HasPrefix(name, "있[") {
			pref[name[:strings.Index(name, "]")+1]] = score
		}
	}
	if len(pref) != 3 || pref["있[보조 용언:있]"] != pref["있[열쇠:동사]"] || pref["있[열쇠:동사]"] == "" {
		t.Error("the two names of the verb have one score", pref)
	}
	if got := r.Lines("먹어 놓다", false); !slices.ContainsFunc(got, func(l string) bool { return strings.Contains(l, "놓[보조 용언:놓]") }) || slices.ContainsFunc(got, func(l string) bool { return strings.Contains(l, "놓[열쇠:") }) {
		t.Error("a closed-set word the table does not list is named by its set alone", got)
	}
}
