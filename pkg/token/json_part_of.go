//ff:func feature=token type=helper control=sequence
//ff:what jsonPartOf — a part as it is written in JSON: the name as it is printed (null when the part has none, a stem no table knows or a sign), known unless the part is a new word (Known), restored when it leaves no trace in the letters
package token

func jsonPartOf(p Part) JSONPart {
	out := JSONPart{Text: p.Text, Known: p.Known(), Restored: p.Unseen, Hidden: p.Hidden, Stem: p.Stem}
	if p.Name != (Name{}) {
		s := corrected(p.Name.String(), p.Written)
		out.Name = &s
	}
	return out
}
