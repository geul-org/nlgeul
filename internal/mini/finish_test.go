//ff:func feature=mini type=reader control=sequence
//ff:what TestFinish — 먹었+다 with its stem, count and pieces; 방식이+다 reads the copula as a piece; no stem gives nothing; 에요 is read after 이 and 아니 only
package mini

import "testing"

func TestFinish(t *testing.T) {
	da := ending{form: "다", kind: kindFinal, count: 7}
	if got, ok := finish([]rune("먹었다"), 2, da, 2, "+다"); !ok || got != (tailFact{"다", kindFinal, "먹", 4, false, 7, "먹+었+다"}) {
		t.Fatal(got, ok)
	}
	if got, ok := finish([]rune("방식이다"), 3, da, 2, "+다"); !ok || got != (tailFact{"다", kindFinal, "방식이", 4, true, idaCount, "방식+이+다"}) {
		t.Fatal(got, ok)
	}
	if _, ok := finish([]rune("었다"), 1, da, 2, "+다"); ok {
		t.Fatal("no stem")
	}
	eyo := ending{form: "에요", flags: fNeedIda}
	if _, ok := finish([]rune("먹에요"), 1, eyo, 4, "+에요"); ok {
		t.Fatal("에요 after a plain stem")
	}
	if got, ok := finish([]rune("아니에요"), 2, eyo, 4, "+에요"); !ok || got.copula || got.stem != "아니" || got.marked != "아니+에요" {
		t.Fatal(got, ok)
	}
	if got, ok := finish([]rune("문제였다"), 3, da, 2, "+다"); !ok || !got.copula || got.marked != "문제+였+다" {
		t.Fatal(got, ok)
	}
	if got, ok := finish([]rune("학생이에요"), 3, eyo, 4, "+에요"); !ok || !got.copula || got.size != 6 || got.marked != "학생+이+에요" {
		t.Fatal(got, ok)
	}
}
