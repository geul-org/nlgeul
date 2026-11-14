//ff:func feature=token type=rule control=iteration dimension=1
//ff:what addFarKey — note one key of the tables: the longest key so far, and for a key longer than a new word's stem may be (newStemMax) the key itself, the key less its last letter (the stem of a word whose last syllable a piece was fused into) and each beginning of it (the letters of a key as far as they are read)
package token

func (ix *index) addFarKey(key string) {
	letters := []rune(key)
	ix.keyMax = max(ix.keyMax, len(letters))
	if len(letters) <= newStemMax {
		return
	}
	ix.farKeys[key] = true
	ix.farBases[string(letters[:len(letters)-1])] = true
	for i := 1; i <= len(letters); i++ {
		ix.keyPrefixes[string(letters[:i])] = true
	}
}
