//ff:func feature=token type=helper control=sequence
//ff:what Plural — does a contest hold more than one branch to choose from
package token

func (b Contest) Plural() bool {
	return len(b.Branches) > 1
}
