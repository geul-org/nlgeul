//ff:func feature=parse type=helper control=sequence
//ff:what romanSure — is a text a Roman numeral of I, V and X (romanWritten) that its letters alone do not leave open: two letters or more, and not X alone repeated (XX and XXX stand for what is not told: XX동, XXX 씨); one letter is a letter as well (I, V, X)
package parse

import "strings"

func romanSure(text string) bool {
	return romanWritten(text) && len(text) > 1 && strings.Trim(text, romanHidden) != ""
}
