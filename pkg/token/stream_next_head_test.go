//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestStreamNextHead — for the sentences the stream tests read, with the small tables and with none, and a head of 0 to 3: every stretch and every way the stream gives with NextHead and AgainHead is the same as with Next and Again but for the candidates of each node, which are exactly the first head of them (ties as they stand), and the stream ends at the same place; the head does cut (아버지가방에 has a node of more candidates than one) and 0 cuts nothing
package token

import "testing"

func TestStreamNextHead(t *testing.T) {
	check := func(r *Reader, s string) []string {
		var bad []string
		for head := 0; head <= 3; head++ {
			bad = append(bad, headWalk(r.Stream(s), r.Stream(s), head)...)
		}
		return bad
	}
	for _, words := range []string{"", "words-stream.tsv", "words.tsv"} {
		if bad := eachSentence(streamReader(t, words), check); len(bad) != 0 {
			t.Fatalf("%q: %q", words, bad)
		}
	}
	r := streamReader(t, "words-stream.tsv")
	full, _ := r.Stream("아버지가방에").Next()
	one, _ := r.Stream("아버지가방에").NextHead(1)
	all, _ := r.Stream("아버지가방에").NextHead(0)
	if len(full.Nodes[0].Contest.Branches) < 2 || len(one.Nodes[0].Contest.Branches) != 1 || len(all.Nodes[0].Contest.Branches) != len(full.Nodes[0].Contest.Branches) {
		t.Fatal(len(full.Nodes[0].Contest.Branches), len(one.Nodes[0].Contest.Branches))
	}
	if _, ok := r.Stream("").NextHead(1); ok {
		t.Fatal("no stretch in an empty sentence")
	}
}
