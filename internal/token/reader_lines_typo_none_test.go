//ff:func feature=token type=reader control=sequence
//ff:what TestReaderLinesTypoNone — with no word table a closing written wrong is read as the right closing's candidate with the letters it is written with added (할께 as 할게's 하+(ㄹ)게 candidate, the stem as the table-less candidate gives it); no noun is read as written wrong, and no candidate of a form written right has a mark
package token

import (
	"strings"
	"testing"
)

func TestReaderLinesTypoNone(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	got := strings.Join(r.Lines("할께", false), "\n")
	if want := "  하+(ㄹ)게[맺음:ㄹ게,오타정정:ㄹ께]\t"; !strings.Contains(got, want) {
		t.Fatalf("missing %q in\n%s", want, got)
	}
	if right := strings.Join(r.Lines("할게", false), "\n"); !strings.Contains(right, "  하+(ㄹ)게[맺음:ㄹ게]\t") || strings.Contains(right, "오타정정") {
		t.Fatal(right)
	}
	if got := strings.Join(r.Lines("역활", false), "\n"); strings.Contains(got, "오타정정") {
		t.Fatal(got)
	}
	if got := strings.Join(r.Lines("갈려고", false), "\n"); !strings.Contains(got, "  가+려고[맺음:려고,오타정정:ㄹ려고]\t") || !strings.Contains(got, "  갈+려고[맺음:려고]\t") {
		t.Fatal(got)
	}
}
