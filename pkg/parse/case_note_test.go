//ff:func feature=parse type=helper control=sequence
//ff:what TestCaseNote — the criteria of a case exception give the note of the word matched, whatever other criteria come before; criteria with no caseRows give none
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/data"
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestCaseNote(t *testing.T) {
	rows := &caseRows{Rows: []caseException{{Word: "담소", GroundKind: "관찰", Ground: "46/46", Row: data.Row{Line: 3}}}}
	specs := toulmin.Specs{listed("caseExceptions", rule.ListWord, []string{"담소"}), rows, &rule.Claim{Name: whyCase}}
	if got := caseNote(specs, "담소"); got != "(담소) — 예외 표 3: 관찰 46/46" {
		t.Fatal(got)
	}
	whole := toulmin.Specs{listed("wholeWords", rule.ListWord, []string{"그리고"}), &rule.Claim{Name: whyWhole}}
	if got := caseNote(whole, "그리고"); got != "" {
		t.Fatal(got)
	}
}
