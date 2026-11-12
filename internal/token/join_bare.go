//ff:func feature=token type=helper control=sequence
//ff:what joinBare — some bundles of signs with one more: a bundle of letters with no name right after another such bundle is joined to it (-> is one bundle), any other is added
package token

func joinBare(out []Bundle, b Bundle) []Bundle {
	if len(out) == 0 || !bareSign(b) || !bareSign(out[len(out)-1]) {
		return append(out, b)
	}
	last := &out[len(out)-1]
	last.Text += b.Text
	last.Branches[0].Words[0].Parts[0].Text = last.Text
	return out
}
