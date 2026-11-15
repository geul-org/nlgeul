//ff:func feature=token type=reader control=sequence
//ff:what TestReaderStreamTypoNone — with no word table a closing written wrong is read as the right closing's candidate with the letters it is written with added (할께 as 할게's 하+(ㄹ)게 candidate, the stem as the table-less candidate gives it); no noun is read as written wrong, and no candidate of a form written right has a mark
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamTypoNone(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	got := streamText(r, "할께")
	if want := "\n하+(ㄹ)게[맺음:ㄹ게,오타정정:ㄹ께]\n"; !strings.Contains(got, want) {
		t.Fatalf("missing %q in\n%s", want, got)
	}
	if right := streamText(r, "할게"); !strings.Contains(right, "\n하+(ㄹ)게[맺음:ㄹ게]\n") || strings.Contains(right, "오타정정") {
		t.Fatal(right)
	}
	if got := streamText(r, "역활"); strings.Contains(got, "오타정정") {
		t.Fatal(got)
	}
	if got := streamText(r, "갈려고"); !strings.Contains(got, "\n가+려고[맺음:려고,오타정정:ㄹ려고]\n") || !strings.Contains(got, "\n갈+려고[맺음:려고]\n") {
		t.Fatal(got)
	}
}
