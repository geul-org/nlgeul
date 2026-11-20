//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderBoundOtherKind — the dependent nouns of the spacing table the closed set now lists too that the word table lists as another kind as well (외, 적, 채) are read under the set's name and under the table's nominal, the set's name first, with one verdict and preference (그 외에, 가 본 적이, 앉은 채로); without a word table they are read under the set's name
package parse

import "testing"

func TestReaderBoundOtherKind(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-bound-other-kind.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	bare, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	for text, want := range map[string][]string{
		"그 외에":   {"외[의존명사:외]+에[부사격:에]", "외[열쇠:체언]+에[부사격:에]"},
		"가 본 적이": {"적[의존명사:적]+이[주격:이]", "적[열쇠:체언]+이[주격:이]"},
		"앉은 채로":  {"채[의존명사:채]+로[부사격:로]", "채[열쇠:체언]+로[부사격:로]"},
	} {
		out := streamText(r, text)
		if bad := orderedLines(out, want); len(bad) > 0 {
			t.Errorf("%q: %q is missing or out of order in\n%s", text, bad, out)
		}
		if first := streamScore(r, text, want[0]); first == "" || first != streamScore(r, text, want[1]) {
			t.Errorf("%q: the two names have one score in\n%s", text, out)
		}
		if out := streamText(bare, text); streamScore(bare, text, want[0]) == "" {
			t.Errorf("%q: no dependent noun without a word table in\n%s", text, out)
		}
	}
}
