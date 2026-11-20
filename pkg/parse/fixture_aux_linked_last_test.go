//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what auxLinkedLast — the candidates of the last word of a sentence read with words-aux-linked.tsv, with the rules of the index or none (rules false) and with a trace or not, and the trace lines Notes gave after the Next of the stretch that holds it
package parse

import "testing"

func auxLinkedLast(t *testing.T, sentence string, rules, trace bool) ([]Branch, []string) {
	t.Helper()
	r := streamReader(t, "words-aux-linked.tsv")
	st := r.Stream(sentence)
	if trace {
		st = r.StreamTrace(sentence)
	}
	if !rules {
		st.rules = nil
	}
	var brs []Branch
	var notes []string
	for seg, ok := st.Next(); ok; seg, ok = st.Next() {
		brs, notes = seg.Nodes[len(seg.Nodes)-1].Contest.Branches, st.Notes()
	}
	return brs, notes
}
