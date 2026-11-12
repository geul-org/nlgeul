//ff:func feature=token type=helper control=iteration dimension=1
//ff:what namesForm — does a tail make its word a modifying form (its last piece is a modifying closing: 작은) or a naming form (a naming closing is among its pieces: 먹기는): such a word is no predicate where it stands
package token

import "strings"

func namesForm(pieces []piece) bool {
	for i, p := range pieces {
		if strings.HasPrefix(p.names, namingPrefix) || (i == len(pieces)-1 && strings.HasPrefix(p.names, adnPrefix)) {
			return true
		}
	}
	return false
}
