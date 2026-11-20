//ff:func feature=parse type=helper control=sequence
//ff:what predicateStem — is the stem of a word as one of its candidates reads it a verb's or an adjective's of the word table (열쇠:동사, 열쇠:형용사) or an auxiliary verb of the closed set (보조 용언:있): not a nominal with the copula, nor a new word
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func predicateStem(n rule.Near) bool {
	return n.Kind == keyKind+":"+kindLabel[kindVerb] || n.Kind == keyKind+":"+kindLabel[kindAdj] || labelOf(n.Kind) == auxSet
}
