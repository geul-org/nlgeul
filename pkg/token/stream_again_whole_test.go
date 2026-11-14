//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestStreamAgainWhole — in the ways of 아버지가방에들어가신다 there is the whole stretch as one word [0-11]: its node holds the readings of the whole as one word (a new word, one candidate or more), each of one word, and that way comes after the ways that cut it into known words
package token

import "testing"

func TestStreamAgainWhole(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	st := r.Stream("아버지가방에들어가신다")
	first, _ := st.Next()
	ways := streamWays(st, first)
	at := -1
	for i, w := range ways {
		if len(w.Nodes) == 1 {
			at = i
		}
	}
	if at <= 0 || ways[at].Nodes[0].Span != (Span{0, 11}) || len(ways[at].Nodes[0].Contest.Branches) < 2 || ways[at].Rank != at+1 {
		t.Fatal(at)
	}
	whole, ok := wholeValues(st.sent[0])
	if !ok || len(whole) != len(ways[at].Nodes[0].Contest.Branches) {
		t.Fatal(whole)
	}
}
