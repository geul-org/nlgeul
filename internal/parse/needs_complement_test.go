//ff:func feature=parse type=helper control=sequence
//ff:what TestNeedsComplement — 되 and 아니 as the whole stem; 사용되 and 먹 are not
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNeedsComplement(t *testing.T) {
	if !needsComplement(&rule.Sentence{Stems: []string{"되"}}) || !needsComplement(&rule.Sentence{Stems: []string{"아니"}}) || needsComplement(&rule.Sentence{Stems: []string{"사용되"}}) || needsComplement(&rule.Sentence{Stems: []string{"먹"}}) || needsComplement(&rule.Sentence{}) {
		t.Fatal("needsComplement")
	}
}
