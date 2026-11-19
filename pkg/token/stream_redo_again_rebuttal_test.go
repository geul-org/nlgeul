//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestStreamRedoAgainRebuttal — a stretch a window cut off whose candidate a rule of rebuttal lowered (enterRebuttalRule: 들어가+시+(ㄴ)다) is read again with the same rebuttal: its first way is no longer the one that reads 들어가신다 as one word, Next sends it so, and Again gives the next ways as the stretch gives them after it was judged, with the same lambdas; the stretch read again holds the lift of the rebuttal
package token

import "testing"

func TestStreamRedoAgainRebuttal(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	plain := redoWays(t, r, 1)
	down := redoWaysRuled(t, r, []sentenceRule{enterRebuttalRule()}, 3)
	if len(down) != 3 || segmentLine(down[0]) == plain[0] {
		t.Fatal("the rebuttal puts another way first", segmentLine(down[0]), plain[0])
	}
	ruled := r.Stream(redoText)
	ruled.rules = []sentenceRule{enterRebuttalRule()}
	ruled.Next()
	ruled.Next()
	sent, _ := ruled.Next()
	if segmentLine(sent) != segmentLine(down[0]) || sent.Lambda != down[0].Lambda || len(ruled.sent[2].lifts) == 0 || ruled.sent[2].lifts[0].force != 0.5 {
		t.Fatal("Next sends it rebutted", segmentLine(sent), sent.Lambda, down[0].Lambda)
	}
	for k := 1; k < 3; k++ {
		seg, ok := ruled.Again(sent.Region)
		if !ok || segmentLine(seg) != segmentLine(down[k]) || seg.Lambda != down[k].Lambda || seg.Rank != k+1 {
			t.Fatal("way", k, segmentLine(seg), seg.Lambda, down[k].Lambda)
		}
	}
	if re := ruled.redone; len(re.lifts) != len(ruled.sent[2].lifts) || re.out[0].Lambda != sent.Lambda {
		t.Fatal("the stretch read again holds the rebuttal", re.lifts)
	}
}
