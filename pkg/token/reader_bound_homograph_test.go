//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderBoundHomograph — the dependent noun of the spacing table the closed set now lists too that the dictionary has another part of speech of the same form for (간) is read under the set's name and under the word table's kind, the set's name first, with one verdict and preference (자식 간에, 형제 간의); without a word table it is read under the set's name (자식 간에)
package token

import (
	"strings"
	"testing"
)

func TestReaderBoundHomograph(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-bound-homograph.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for text, want := range map[string][]string{
		"자식 간에": {"간[의존명사:간]+에[부사격:에]", "간[열쇠:체언]+에[부사격:에]"},
		"형제 간의": {"간[의존명사:간]+의[관형격:의]", "간[열쇠:체언]+의[관형격:의]"},
	} {
		out := strings.Join(r.Lines(text, false), "\n")
		if bad := orderedLines(out, want); len(bad) > 0 {
			t.Errorf("%q: %q is missing or out of order in\n%s", text, bad, out)
		}
		if first := lineScore(out, want[0]); first == "" || first != lineScore(out, want[1]) {
			t.Errorf("%q: the two names have one score in\n%s", text, out)
		}
	}
	bare, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	if out := strings.Join(bare.Lines("자식 간에", false), "\n"); lineScore(out, "간[의존명사:간]+에[부사격:에]") == "" {
		t.Errorf("자식 간에: no dependent noun without a word table in\n%s", out)
	}
}
