//ff:func feature=parse type=helper control=sequence
//ff:what adnominal — is a word's candidate an adnominal (관형어): its last part is a modifying closing (꾸밈:는, 꾸밈:ㄴ), or it is a determiner's stem standing alone with no marker (이[관형사:이], 각[열쇠:관형사])
package parse

import "strings"

func adnominal(parts []Part) bool {
	if len(parts) == 0 {
		return false
	}
	if len(parts) == 1 {
		return parts[0].Stem && determiner(parts[0].Name)
	}
	return strings.HasPrefix(parts[len(parts)-1].Name.String(), adnPrefix)
}
