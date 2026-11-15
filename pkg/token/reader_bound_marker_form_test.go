//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderBoundMarkerForm — the dependent nouns of the spacing table the closed set now lists too whose letters the grammar tables also read as a particle, a closing or a pre-final ending (시, 듯, 만) are read under the set's name and under the word table's kind, the set's name first, with one verdict and preference (필요할 시에는, 올 듯 말 듯, 십 년 만에); without a word table they are read under the set's name
package token

import (
	"strings"
	"testing"
)

func TestReaderBoundMarkerForm(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-bound-marker.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	bare, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	for text, want := range map[string][]string{
		"필요할 시에는": {"시[의존명사:시]+에[부사격:에]+는[보조사:는]", "시[열쇠:체언]+에[부사격:에]+는[보조사:는]"},
		"올 듯 말 듯": {"듯[의존명사:듯]", "듯[열쇠:체언]"},
		"십 년 만에":  {"만[의존명사:만]+에[부사격:에]", "만[열쇠:체언]+에[부사격:에]"},
	} {
		out := strings.Join(r.Lines(text, false), "\n")
		if bad := orderedLines(out, want); len(bad) > 0 {
			t.Errorf("%q: %q is missing or out of order in\n%s", text, bad, out)
		}
		if first := lineScore(out, want[0]); first == "" || first != lineScore(out, want[1]) {
			t.Errorf("%q: the two names have one score in\n%s", text, out)
		}
		if out := strings.Join(bare.Lines(text, false), "\n"); !strings.Contains(out, want[0]) {
			t.Errorf("%q: no dependent noun without a word table in\n%s", text, out)
		}
	}
}
