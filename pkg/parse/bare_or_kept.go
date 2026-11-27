//ff:func feature=parse type=helper control=sequence
//ff:what bareOrKept — is a part a stem that is a sign as one with no name was before the tables of special letters named it: a stem with no name, or a stem named by the hand-written table or the table of kinds (keptName: ○○[가림:원], **[기호:별표]); a part that is no stem, and one named by any other table, is not
package parse

func bareOrKept(p Part) bool {
	return p.Stem && (p.Name == (Name{}) || keptName(firstRune(p.Text), p.Name.String()))
}
