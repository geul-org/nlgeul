//ff:func feature=token type=helper control=sequence
//ff:what bareSign — is a bundle one reading of one part with no name: signs sent up as their letters
package token

func bareSign(b Bundle) bool {
	return len(b.Branches) == 1 && len(b.Branches[0].Words) == 1 && len(b.Branches[0].Words[0].Parts) == 1 && b.Branches[0].Words[0].Parts[0].Name == (Name{})
}
