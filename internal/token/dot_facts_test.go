//ff:func feature=token type=reader control=sequence
//ff:what TestDotFacts — the sides, the label to the right and the last label of a dotted name, those two in lower case; v1.2.3 is numbers with a leading v, task.py is not numbers, v alone is no number
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestDotFacts(t *testing.T) {
	m := dotFacts([]string{"v1", "2", "3"}, 0)
	if !m.Numeric || m.Ext != "2" || m.Last != "3" || m.Left != rule.SideAlnum || m.Right != rule.SideAlnum|rule.SideDigit {
		t.Fatalf("%+v", m)
	}
	if m = dotFacts([]string{"task", "py"}, 0); m.Numeric || m.Ext != "py" || m.Run != 1 {
		t.Fatalf("%+v", m)
	}
	if m = dotFacts([]string{"v", "2"}, 0); m.Numeric {
		t.Fatalf("%+v", m)
	}
	if m = dotFacts([]string{"Notes", "TXT", "Bak"}, 0); m.Ext != "txt" || m.Last != "bak" {
		t.Fatalf("%+v", m)
	}
}
