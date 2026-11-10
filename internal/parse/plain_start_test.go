//ff:func feature=parse type=reader control=sequence
//ff:what TestPlainStart — 으면 after a consonant, 면 after a vowel or ㄹ, not 면 right after another consonant; 는다 needs a consonant
package parse

import "testing"

func TestPlainStart(t *testing.T) {
	myeon := ending{form: "면", flags: fEu}
	if start, ok := plainStart([]rune("먹으면"), 3, myeon); !ok || start != 1 {
		t.Fatal(start, ok)
	}
	if start, ok := plainStart([]rune("가면"), 2, myeon); !ok || start != 1 {
		t.Fatal(start, ok)
	}
	if _, ok := plainStart([]rune("측면"), 2, myeon); ok {
		t.Fatal("측면")
	}
	if _, ok := plainStart([]rune("면"), 1, myeon); ok {
		t.Fatal("no stem")
	}
	neunda := ending{form: "는다", after: afterBatchim}
	if _, ok := plainStart([]rune("먹는다"), 3, neunda); !ok {
		t.Fatal("먹는다")
	}
	if _, ok := plainStart([]rune("가는다"), 3, neunda); ok {
		t.Fatal("가는다")
	}
}
