//ff:func feature=token type=reader control=sequence
//ff:what TestCaseNames — one name for each case; 보조사 for none; 접미 for the suffix-like
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestCaseNames(t *testing.T) {
	two := caseNames(rule.CaseSubj|rule.CaseComp, 3)
	if len(two) != 2 || two[0] != "주격" || two[1] != "보격" || caseNames(0, 4)[0] != "보조사" || caseNames(0, 0)[0] != "접미" {
		t.Fatal(two)
	}
}
