//ff:func feature=token type=helper control=iteration dimension=2
//ff:what Label — what a branch is called in a contest's head line: the labels of every part of its words, joined with " + "
package token

import "strings"

func (b Branch) Label() string {
	var labels []string
	for _, w := range b.Words {
		for _, p := range w.Parts {
			labels = append(labels, p.Label())
		}
	}
	return strings.Join(labels, " + ")
}
