//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what numberLetter — is r a letter that writes a number in a shape of its own, a number of one letter (①, ⓪, ⑴, ⒈, ❶, ½, ²): a letter the tables of special letters give the content 수 that is no sign (signRune) and no letter of a run of Latin letters and digits (a full-width digit and a Roman numeral letter are read as runs are)
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func numberLetter(r rune) bool {
	for _, f := range forms[r] {
		if f.content == numName {
			return !signRune(r) && !rule.RunRune(r)
		}
	}
	return false
}
