//ff:func feature=token type=helper control=sequence
//ff:what afterCopula — the name the copula's graph is told stands after its stem: the auxiliary written onto the tail's 어/아 form when it is one the copula's stem does not take before it (copulaBars: 하 of 어하다), else what the tail begins with (firstName)
package token

import "slices"

func afterCopula(t tail) string {
	if len(t.pieces) > 1 && slices.Contains(copulaBars, t.pieces[1].names) {
		return t.pieces[1].names
	}
	return firstName(t)
}
