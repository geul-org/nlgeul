//ff:func feature=token type=reader control=sequence
//ff:what TestLoneHonors — 가심 reads the noun form under the honorific, 가세 the fused 어 when asked for; another letter, the start, and no word table give nothing
package token

import "testing"

func TestLoneHonors(t *testing.T) {
	got := wordsScan("가심을", layerCost).loneHonors(2, false)
	if len(got) != 1 || got[0].pieces[1].text != "(ㅁ)" || got[0].close == nil || got[0].close.name != nounName || got[0].gate == nil {
		t.Fatal(got)
	}
	got = wordsScan("가세요", layerCost).loneHonors(2, true)
	if len(got) != 1 || got[0].pieces[1].names != vowelName || got[0].start != 1 {
		t.Fatal(got)
	}
	sc := wordsScan("가세요", layerCost)
	if len(sc.loneHonors(2, false)) != 0 || len(sc.loneHonors(0, true)) != 0 || len(scanOf("가심").loneHonors(2, false)) != 0 {
		t.Fatal("read where it should not be")
	}
}
