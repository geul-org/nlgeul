//ff:func feature=parse type=helper control=sequence
//ff:what genitiveContract — is a candidate a form the contracts table puts the adnominal particle 의 under (genitiveForms: 내, 네, 제), read as its pronoun of the word table with that particle alone (나, 너, 저 with 관형격:의); not the pronoun with 의 written out (나의), whose letters are not the form's
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func genitiveContract(n rule.Near) bool {
	return n.Kind == tableNoun && slices.Contains(genitiveNames, n.Marks) && slices.Contains(genitiveForms, n.Text)
}
