//ff:func feature=token type=helper control=sequence
//ff:what TestTakeWays — the ways asked of a real stream (아버지가방에들어가신다, 60 ways of one stretch): all of them with head 0 (the last asking finds none left), head of them with head, and no way is made beyond head; a stretch with one way is not asked for another, and one that gives fewer ways than it said stops the asking
package token

import (
	"math/big"
	"testing"
)

func TestTakeWays(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	s := "아버지가방에들어가신다"
	st := r.Stream(s)
	first, _ := st.Next()
	if got := takeWays(st, first, 3); len(got) != 3 || got[0].Rank != 1 || got[2].Rank != 3 || len(st.sent[0].out) != 3 {
		t.Fatal(len(got), len(st.sent[0].out))
	}
	all := r.Stream(s)
	first, _ = all.Next()
	got := takeWays(all, first, 0)
	if len(got) != 60 || got[59].Rank != 60 || got[59].Lambda > got[0].Lambda {
		t.Fatal(len(got))
	}
	exact := r.Stream(s)
	first, _ = exact.Next()
	if got := takeWays(exact, first, 60); len(got) != 60 || len(exact.sent[0].out) != 60 {
		t.Fatal(len(got))
	}
	larger := r.Stream(s)
	first, _ = larger.Next()
	if got := takeWays(larger, first, 500); len(got) != 60 {
		t.Fatal(len(got))
	}
	spaced := r.Stream("아버지가 방에 들어가신다")
	first, _ = spaced.Next()
	if first.Ways.Cmp(big.NewInt(2)) != 0 {
		t.Fatal(first.Ways)
	}
	if got := takeWays(spaced, first, 1); len(got) != 1 || len(spaced.sent[0].out) != 1 {
		t.Fatal(len(got))
	}
	second, _ := spaced.Next()
	if second.Ways.Cmp(big.NewInt(1)) != 0 {
		t.Fatal(second.Ways)
	}
	if got := takeWays(spaced, second, 5); len(got) != 1 || len(spaced.sent[1].out) != 1 {
		t.Fatal(len(got), len(spaced.sent[1].out))
	}
	second.Ways = big.NewInt(4)
	if got := takeWays(spaced, second, 0); len(got) != 1 {
		t.Fatal("a stretch that gives no more than it said stops the asking", len(got))
	}
}
