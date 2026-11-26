//ff:func feature=parse type=helper control=sequence
//ff:what TestGenitiveFormsOf — the forms of the contracts whose particle is named 관형격:의, in the order given (none from none); from the contracts table the form is 내 (genitiveForms)
package parse

import (
	"slices"
	"testing"
)

func TestGenitiveFormsOf(t *testing.T) {
	got := genitiveFormsOf([]contract{{form: "내", stem: "나", name: "관형격:의"}, {form: "난", stem: "나", name: "보조사:는"}, {form: "네", stem: "너", name: "관형격:의"}})
	if !slices.Equal(got, []string{"내", "네"}) {
		t.Fatal(got)
	}
	if got := genitiveFormsOf(nil); got != nil {
		t.Fatal(got)
	}
	if !slices.Equal(genitiveForms, []string{"내"}) {
		t.Fatal(genitiveForms)
	}
}
