//ff:func feature=token type=helper control=sequence
//ff:what modifiable — may a word ending in a modifying closing stand right before a word: its first part is a nominal stem (nominalStem), an order or a determiner, or its last part is a modifying or a naming closing (it modifies the nominal after it, or stands as a nominal itself)
package token

import "strings"

// An adnominal modifies the nominal after it; a word in between that
// is itself a determiner or a modifying form modifies the same
// nominal (새로운 그 길, 학생인 착한 아이). An adverb, a conjunction, a
// predicate that ends otherwise and a sign are no such word.
func modifiable(parts []Part) bool {
	if len(parts) == 0 {
		return false
	}
	first := parts[0]
	if nominalStem(first) || (first.Stem && (determiner(first.Name) || first.Name.Kind == nameOf(orderName).Kind)) {
		return true
	}
	last := parts[len(parts)-1].Name.String()
	return strings.HasPrefix(last, adnPrefix) || strings.HasPrefix(last, namingPrefix)
}
