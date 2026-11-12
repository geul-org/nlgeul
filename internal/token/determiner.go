//ff:func feature=token type=helper control=iteration dimension=1
//ff:what determiner — are a word's readings those of a determiner and nothing else: each name is a determiner of the closed set (관형사:이) or the word table's kind 관형사; a word with no name (a new word) or one that is a nominal too (열쇠:체언|관형사) is not
package token

func determiner(names []Name) bool {
	for _, n := range names {
		if n.Kind != determinerLabel && (n.Kind != keyKind || n.Form != determinerLabel) {
			return false
		}
	}
	return len(names) > 0
}
