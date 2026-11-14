//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestStreamNext — Next gives the first way of each stretch in turn and reads a unit only when it needs one (after the first Next the second unit is not read, and only that one way of the stretch is made); what is asked in between by Again does not change what Next gives later (two streams of the same sentence, one asked for more, give the same stretches after); a stretch with no candidate at all is given as it is, with no nodes
package token

import (
	"math/big"
	"testing"
)

func TestStreamNext(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	s := "아버지가 방에 들어가신다"
	st := r.Stream(s)
	first, ok := st.Next()
	if !ok || first.Rank != 1 || first.Ways.Cmp(big.NewInt(2)) != 0 || first.Region != (Span{0, 4}) || segmentLine(first) != "아버지가[0-4]" {
		t.Fatal(first, ok)
	}
	if begun(st) != 1 || len(st.sent) != 1 || len(st.sent[0].out) != 1 || st.sent[0].sent != 1 || len(st.queue) != 0 {
		t.Fatal("read ahead", begun(st), len(st.sent[0].out))
	}
	if _, ok := st.Again(first.Region); !ok {
		t.Fatal("a second way")
	}
	rest := streamFirsts(st)
	plain := streamFirsts(r.Stream(s))
	if len(rest) != 2 || len(plain) != 3 || !sameStrings(waysLines(rest), waysLines(plain[1:])) || rest[0].Lambda != plain[1].Lambda || rest[1].Ways.Cmp(big.NewInt(8)) != 0 {
		t.Fatal(waysLines(rest), waysLines(plain))
	}
	for _, rg := range st.sent[1:] {
		if len(rg.out) != 1 {
			t.Fatal("only the first way is made", len(rg.out))
		}
	}
	if sg, ok := st.Next(); ok || sg.Nodes != nil {
		t.Fatal("end", sg, ok)
	}

	empty := &Stream{queue: []*region{newRegion(Contest{Text: "가나"}, 5, nil)}}
	sg, ok := empty.Next()
	if !ok || sg.Region != (Span{5, 7}) || len(sg.Nodes) != 0 || sg.Rank != 0 || sg.Ways.Cmp(big.NewInt(0)) != 0 {
		t.Fatal(sg, ok)
	}
	if _, ok := empty.Again(Span{5, 7}); ok {
		t.Fatal("no way for it")
	}

	far := streamReader(t, "words-far.tsv")
	cut := far.Stream(windowText(5))
	if sg, ok := cut.Next(); !ok || segmentLine(sg) != "고양이가[0-4]" || len(cut.sent) != 1 || cut.sent[0].nodes != nil || cut.sent[0].redo == nil || cut.sent[0].sent != 1 {
		t.Fatal("a stretch cut off by a window is kept as the letters it covers and how to read it again", cut.sent)
	}
	if cut.cur == nil || cut.cur.run.j >= cut.cur.run.n {
		t.Fatal("the rest of the unit is not read yet")
	}
}
