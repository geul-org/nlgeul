//ff:func feature=token type=helper control=sequence
//ff:what nameKind — the kind of a name as the rules of a sentence read it, the part before its colon (꾸밈 of 꾸밈:ㄹ, 의존명사 of 의존명사:수); the name itself when it has no colon
package token

import "strings"

func nameKind(name string) string {
	kind, _, _ := strings.Cut(name, ":")
	return kind
}
