//ff:func feature=token type=helper control=sequence
//ff:what determiner — is a word's reading that of a determiner: its name is a determiner of the closed set (관형사:이) or the word table's kind 관형사 (열쇠:관형사); a word with no name (a new word) or one that is a nominal is not
package token

func determiner(n Name) bool {
	return n.Kind == determinerLabel || (n.Kind == keyKind && n.Form == determinerLabel)
}
