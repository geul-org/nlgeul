//ff:func feature=token type=helper control=sequence
//ff:what String — a part printed: what the letters do not show in brackets, its letters, and its readings in square brackets (with the form it is written in when that is wrong); a stem no table knows is printed bare
package token

func (p Part) String() string {
	text := p.Text
	if p.Hidden != "" {
		text = "(" + p.Hidden + ")" + text
	}
	if p.Stem && len(p.Names) == 0 {
		return text
	}
	return text + "[" + corrected(namesText(p.Names), p.Written) + "]"
}
