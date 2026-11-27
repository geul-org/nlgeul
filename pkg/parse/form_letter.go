//ff:func feature=parse type=helper control=sequence
//ff:what formLetter — is r a letter of the tables of special letters that is read as a letter of its own inside a unit, an atom (㊀, ㈠, ㆒): a letter of the tables that is no sign (signRune: a sign is a contest of its own at a unit's edge) and no letter of a run of Latin letters and digits (those are read as runs are)
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func formLetter(r rune) bool {
	return len(forms[r]) > 0 && !signRune(r) && !rule.RunRune(r)
}
