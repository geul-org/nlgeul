//ff:func feature=token type=reader control=sequence
//ff:what boundAfterAdn — is the stem from b the dependent noun 것 or its spoken form 거 (adnOnto) written right after a modifying closing: the stem does not begin the unit and some candidate has a word that ends in a modifying closing end at b (adnEnds)
package token

import "slices"

// The fact is read here; what it lifts is the kinds' graphs' to judge.
// The word before is asked as afterAdn asks it: whether any candidate
// has such a word end there.
func (sc *scan) boundAfterAdn(b int, stem string) bool {
	return b > 0 && sc.adnEnds[b] && slices.Contains(adnOnto, stem)
}
