//ff:func feature=parse type=reader control=sequence
//ff:what TestPolite — the closings read at one place, each read once more after each polite slot that stands before it: 하오며 reads 하+오+며, 하오니 the tables' closing 오니 alone, 하옵니다 reads 옵 and 니다, 가온 reads 오 under the fused ㄴ; with no word table there is no such slot and nothing is read
package parse

import "testing"

func TestPolite(t *testing.T) {
	has := func(tails []tail, start int, names ...string) bool {
		for _, x := range tails {
			if x.start != start || len(x.pieces) != len(names) {
				continue
			}
			same := true
			for i, n := range names {
				same = same && x.pieces[i].names == n
			}
			if same {
				return true
			}
		}
		return false
	}
	sc := wordsScan("하오며", layerCost)
	if got := sc.polite(sc.closings(3, false, false)); !has(got, 1, "선어말:오", "맺음:며") {
		t.Fatal(got)
	}
	sc = wordsScan("하오니", layerCost)
	if got := sc.closings(3, false, false); has(got, 1, "선어말:오", "맺음:니") || !has(got, 1, "맺음:오니") {
		t.Fatal("오니 is the tables' closing", got)
	}
	sc = wordsScan("하옵니다", layerCost)
	if got := sc.closings(4, false, false); !has(got, 1, "선어말:옵", "맺음:니다") {
		t.Fatal("closings read the polite slots", got)
	}
	sc = wordsScan("가온", layerCost)
	if got := sc.closings(2, false, false); !has(got, 1, "선어말:오", "꾸밈:ㄴ") {
		t.Fatal(got)
	}
	sc = wordsScan("학교이옵니다", layerCost)
	if got := sc.closings(6, false, false); !has(got, 2, "서술격:이", "선어말:옵", "맺음:니다") || has(got, 3, "서술격:이(빠짐)", "선어말:옵", "맺음:니다") {
		t.Fatal("the copula stem only where it is written", got)
	}
	bare := scanOf("하오며")
	if got := bare.polite(bare.closings(3, false, false)); len(got) != 0 || len(bare.ix.polite) != 0 {
		t.Fatal("no word table, no polite slot", got)
	}
}
