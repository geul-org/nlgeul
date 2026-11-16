//ff:func feature=token type=helper control=sequence
//ff:what TestDativeStemsOf — the words of the contracts whose particle is named 부사격:에게, in the order given (none from none); from the contracts table they are 내, 네 and 제 (dativeStems)
package token

import (
	"slices"
	"testing"
)

func TestDativeStemsOf(t *testing.T) {
	got := dativeStemsOf([]contract{{form: "내게", stem: "내", name: "부사격:에게"}, {form: "이게", stem: "이것", name: "주격:이"}, {form: "제게", stem: "제", name: "부사격:에게"}})
	if !slices.Equal(got, []string{"내", "제"}) {
		t.Fatal(got)
	}
	if got := dativeStemsOf(nil); got != nil {
		t.Fatal(got)
	}
	if !slices.Equal(dativeStems, []string{"내", "네", "제"}) {
		t.Fatal(dativeStems)
	}
}
