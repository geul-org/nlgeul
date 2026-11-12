//ff:func feature=token type=reader control=sequence
//ff:what Lines — read a token stream and print what is sent up, every reading of every contest (LinesHead with no limit)
package token

// A unit is read without looking outside it, so the same unit always
// gives the same lines. The bundles are those Bundles gives.
func (r *Reader) Lines(sentence string, trace bool) []string {
	return r.LinesHead(sentence, trace, 0)
}
