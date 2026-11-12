//ff:func feature=token type=helper control=sequence
//ff:what Label — what a part is called in a contest's head line: its readings (with the form it is written in when that is wrong), or 새 낱말 for a stem no table knows
package token

func (p Part) Label() string {
	if p.Stem && len(p.Names) == 0 {
		return newWord
	}
	return corrected(namesText(p.Names), p.Written)
}
