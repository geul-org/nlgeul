//ff:func feature=token type=helper control=sequence
//ff:what joinBare — some contests of signs with one more: a contest of letters with no name right after another such contest is joined to it (-> is one contest), any other is added
package token

func joinBare(out []Contest, b Contest) []Contest {
	if len(out) == 0 || !bareSign(b) || !bareSign(out[len(out)-1]) {
		return append(out, b)
	}
	last := &out[len(out)-1]
	last.Text += b.Text
	last.Branches[0].Words[0].Parts[0].Text = last.Text
	return out
}
