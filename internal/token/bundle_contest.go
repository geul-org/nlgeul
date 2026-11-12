//ff:func feature=token type=helper control=sequence
//ff:what Contest — is a bundle a contest: more than one branch to choose from
package token

func (b Bundle) Contest() bool {
	return len(b.Branches) > 1
}
