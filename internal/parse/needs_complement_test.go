//ff:func feature=parse type=helper control=sequence
//ff:what TestNeedsComplement — 되 and 아니 as the whole stem; 사용되 and 먹 are not
package parse

import "testing"

func TestNeedsComplement(t *testing.T) {
	if !needsComplement(&sentence{stems: []string{"되"}}) || !needsComplement(&sentence{stems: []string{"아니"}}) || needsComplement(&sentence{stems: []string{"사용되"}}) || needsComplement(&sentence{stems: []string{"먹"}}) || needsComplement(&sentence{}) {
		t.Fatal("needsComplement")
	}
}
