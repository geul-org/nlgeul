//ff:func feature=token type=helper control=sequence
//ff:what Label — what a part is called in a contest's head line: its readings (with the form it is written in when that is wrong), 새 낱말 for a stem no table knows, or its letters for signs sent up with no name
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func (p Part) Label() string {
	if p.Stem && len(p.Names) == 0 && len(rule.Trim(p.Text)) == 0 {
		return p.Text
	}
	if p.Stem && len(p.Names) == 0 {
		return newWord
	}
	return corrected(namesText(p.Names), p.Written)
}
