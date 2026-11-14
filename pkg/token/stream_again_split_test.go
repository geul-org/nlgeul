//ff:func feature=token type=reader control=sequence
//ff:what TestStreamAgainSplit — 돌아가셨다 is one stretch of three ways: the whole as one word [0-5] (with the readings of a known verb 돌아가+시+(었)+다 and of a new word, the two candidates), 돌아 | 가셨다 and 돌 | 아가셨다; a node holds no candidate of more than one word, whatever the ways
package token

import (
	"math/big"
	"testing"
)

func TestStreamAgainSplit(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	st := r.Stream("돌아가셨다")
	first, ok := st.Next()
	if !ok || first.Ways.Cmp(big.NewInt(3)) != 0 || first.Region != (Span{0, 5}) || len(first.Nodes) != 1 || len(first.Nodes[0].Contest.Branches) != 2 {
		t.Fatal(first, ok)
	}
	whole := first.Nodes[0].Contest.Branches
	if whole[0].String() != "돌아가[열쇠:동사]+시[선어말:시]+(었)[선어말:ㅆ]+다[맺음:다]" || whole[1].String() != "돌아가셨다" || whole[1].Verdict >= whole[0].Verdict {
		t.Fatal(whole)
	}
	ways := streamWays(st, first)
	if got := waysLines(ways); !sameStrings(got, []string{"돌아가셨다[0-5]", "돌아[0-2] 가셨다[2-5]", "돌[0-1] 아가셨다[1-5]"}) || !ranked(ways) {
		t.Fatal(got)
	}
	if n := len(ways[1].Nodes[1].Contest.Branches); n != 1 || ways[1].Nodes[0].Contest.Glued || !ways[1].Nodes[1].Contest.Glued {
		t.Fatal("Glued", n)
	}
}
