//ff:func feature=token type=reader control=sequence
//ff:what TestNewStream — a stream over a sentence has its units cut out and nothing read yet: no region, no trace line, the position at the start with a space taken as standing before it, no wait open and nothing judged; its rules of a sentence are the index's
package token

import "testing"

func TestNewStream(t *testing.T) {
	r := streamReader(t, "")
	st := newStream(r.ix, "아버지 가방에 들어간다", false)
	if len(st.units) != 3 || st.read != 0 || st.cursor != 0 || !st.spaced || len(st.queue) != 0 || len(st.sent) != 0 || len(st.notes) != 0 || st.trace || st.ix != r.ix {
		t.Fatal(st)
	}
	if traced := newStream(r.ix, "아버지", true); !traced.trace || len(traced.units) != 1 {
		t.Fatal(traced)
	}
	if st.held || st.behind != nil || st.pending != nil || len(st.rules) != len(r.ix.sentence) {
		t.Fatal(st.held, st.behind, st.pending)
	}
	r.ix.sentence = []sentenceRule{complementRule()}
	if ruled := newStream(r.ix, "학생이", false); len(ruled.rules) != 1 || ruled.rules[0].name != "보격 뒤 되" {
		t.Fatal(ruled.rules)
	}
}
