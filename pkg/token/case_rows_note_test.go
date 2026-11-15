//ff:func feature=token type=helper control=sequence
//ff:what TestCaseRowsNote — the note of a word a row has names the word, the row's line, the kind of its ground and the ground (the row of that word, not the first); a word no row has, and evidence that is no word, get none
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/data"
)

func TestCaseRowsNote(t *testing.T) {
	c := &caseRows{Rows: []caseException{
		{Word: "적소", GroundKind: "규범", Ground: "조항", Row: data.Row{Line: 2}},
		{Word: "담소", GroundKind: "관찰", Ground: "46/46, 29파일", Row: data.Row{Line: 4}},
	}}
	if got := c.note("담소"); got != "(담소) — 예외 표 4: 관찰 46/46, 29파일" {
		t.Fatal(got)
	}
	if got := c.note("감소"); got != "" {
		t.Fatal(got)
	}
	if got := c.note(nil); got != "" {
		t.Fatal(got)
	}
}
