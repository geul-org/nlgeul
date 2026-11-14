//ff:func feature=token type=helper control=sequence
//ff:what text — what is put after the box line at a line: a space, the inner and the outer letters, and a space with the tag when the line has one
package token

func (b *bracket) text(line int) string {
	s := " " + string(b.inner[line]) + string(b.outer[line])
	if b.tags[line] != "" {
		s += " " + b.tags[line]
	}
	return s
}
