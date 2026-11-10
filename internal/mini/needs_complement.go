//ff:func feature=mini type=helper control=sequence
//ff:what needsComplement — is the sentence's predicate 되다 or 아니다: its stem under the longest tail is exactly 되 or 아니
package mini

import "slices"

// The two predicates are named by the grammar (morphemes.md sections 2
// and 4), not looked up. 사용되다 has the stem 사용되, not 되.
func needsComplement(s *sentence) bool {
	return slices.Contains(s.stems, "되") || slices.Contains(s.stems, "아니")
}
