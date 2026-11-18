//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestCaseExceptionsNaga — the rows of 나가 in the case exception table held in the program (graphs 주격:가 and 보격:가, ground the norm: the pronoun 나 is 내 before 가): end to end with a word table of testdata, 나가 is read first as 나가다 with 어 and no longer as the pronoun 나 with the nominative 가, which comes first without the rows; 내가 keeps the pronoun with 가 first
package token

import (
	"slices"
	"testing"
)

func TestCaseExceptionsNaga(t *testing.T) {
	for _, graph := range []string{"주격:가", "보격:가"} {
		i := slices.IndexFunc(caseExceptions[graph], func(c caseException) bool { return c.Word == "나가" })
		if i < 0 || caseExceptions[graph][i].Reading != "나가[열쇠:동사]+(아)[맺음:어(합침)]" || caseExceptions[graph][i].GroundKind != groundNorm {
			t.Fatal(graph, caseExceptions[graph])
		}
	}
	verb, pron := "나가[열쇠:동사]+[맺음:어(합침)]", "나[열쇠:체언]+가[주격:가]"
	w, ok := jsonWordAt(streamReader(t, "words-naga.tsv").JSONSentence("밖으로 나가 보았다", 0), "나가")
	if got := jsonWordShown(w); !ok || got[0] != verb || slices.Contains(got, pron) {
		t.Errorf("나가: %v", got)
	}
	w, ok = jsonWordAt(streamReader(t, "words-naga.tsv").JSONSentence("내가 가", 0), "내가")
	if got := jsonWordShown(w); !ok || got[0] != "내[열쇠:체언]+가[주격:가]" {
		t.Errorf("내가: %v", got)
	}
	old := caseExceptions
	t.Cleanup(func() { caseExceptions = old })
	caseExceptions = map[string][]caseException{}
	w, _ = jsonWordAt(streamReader(t, "words-naga.tsv").JSONSentence("밖으로 나가 보았다", 0), "나가")
	if got := jsonWordShown(w); got[0] != pron {
		t.Fatal("without the rows the pronoun with 가 comes first", got)
	}
}
