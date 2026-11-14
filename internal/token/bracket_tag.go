//ff:func feature=token type=helper control=sequence
//ff:what tag — put a tag on a line of the bracket; a line that has one already keeps it and gets the new one after it, two spaces apart
package token

func (b *bracket) tag(line int, text string) {
	if b.tags[line] != "" {
		text = b.tags[line] + "  " + text
	}
	b.tags[line] = text
}
