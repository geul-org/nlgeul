//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestStreamRedoAgain — the ways of a stretch a window cut off come one after another as the stretch itself gave them; going to another such stretch and back passes over the ways already sent; once the ways are used up none is given, and only the one stretch read again last is held. A stretch whose candidate the rules of a sentence raised (honorRule: 들어가신+다 by 30) is read again with the same lift: Next sends it raised, and Again gives the ways in the raised order with the raised lambdas, the ways the stretch itself gives after it was judged
package token

import "testing"

func TestStreamRedoAgain(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	want := redoWays(t, r, 5)
	st := r.Stream(redoText)
	st.Next()
	st.Next()
	a, _ := st.Next()
	st.Next()
	st.Next()
	b, _ := st.Next()
	for k := 1; k < 3; k++ {
		if seg, ok := st.Again(a.Region); !ok || segmentLine(seg) != want[k] {
			t.Fatal("way", k, segmentLine(seg), want[k])
		}
	}
	held := st.redone
	if _, ok := st.Again(b.Region); !ok || st.redone == held || st.redoOf == nil {
		t.Fatal("another stretch is read again")
	}
	if seg, ok := st.Again(a.Region); !ok || segmentLine(seg) != want[3] || seg.Rank != 4 {
		t.Fatal("back to the first, past the ways sent", segmentLine(seg))
	}
	for ok := true; ok; _, ok = st.Again(a.Region) {
	}
	if _, ok := st.Again(a.Region); ok {
		t.Fatal("used up")
	}

	raised := redoWaysRuled(t, r, []sentenceRule{honorRule()}, 3)
	if len(raised) != 3 || raised[0].Nodes[0].Contest.Branches[0].String() != "들어가신+다[맺음:다]" || segmentLine(raised[0]) != want[0] {
		t.Fatal("the lift puts 들어가신+다 first", raised[0].Nodes[0].Contest.Branches[0].String())
	}
	ruled := r.Stream(redoText)
	ruled.rules = []sentenceRule{honorRule()}
	ruled.Next()
	ruled.Next()
	sent, _ := ruled.Next()
	if sent.Lambda != raised[0].Lambda || sent.Nodes[0].Contest.Branches[0].String() != "들어가신+다[맺음:다]" || len(ruled.sent[2].lifts) == 0 {
		t.Fatal("Next sends it raised", sent.Lambda, raised[0].Lambda)
	}
	for k := 1; k < 3; k++ {
		seg, ok := ruled.Again(sent.Region)
		if !ok || segmentLine(seg) != segmentLine(raised[k]) || seg.Lambda != raised[k].Lambda || seg.Rank != k+1 || seg.Nodes[0].Contest.Branches[0].String() != raised[k].Nodes[0].Contest.Branches[0].String() {
			t.Fatal("way", k, segmentLine(seg), seg.Lambda, raised[k].Lambda)
		}
	}
	if re := ruled.redone; len(re.lifts) != len(ruled.sent[2].lifts) || re.out[0].Nodes[0].Contest.Branches[0].String() != "들어가신+다[맺음:다]" || re.out[0].Lambda != sent.Lambda {
		t.Fatal("the stretch read again holds the lift", re.lifts)
	}
}
