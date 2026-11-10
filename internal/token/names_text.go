//ff:func feature=token type=helper control=iteration dimension=1
//ff:what namesText — some readings printed: kind:form, joined with "|"; a kind alone when the form is empty; the kinds of a word table's stem under one 열쇠 (열쇠:체언|부사)
package token

import "strings"

func namesText(names []Name) string {
	out := make([]string, 0, len(names))
	for i, n := range names {
		one := n.Kind
		if n.Form != "" {
			one += ":" + n.Form
		}
		if i > 0 && n.Kind == keyKind && names[i-1].Kind == keyKind {
			one = n.Form
		}
		out = append(out, one)
	}
	return strings.Join(out, "|")
}
