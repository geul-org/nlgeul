//ff:func feature=parse type=rule control=sequence
//ff:what TestIndexAddSentenceCountedShort — the index of a reader holds the rule of a sentence ruleCountedShort after the rule of a noun right after a modifier (nounModifiedRule) and the rule of a demonstrative standing alone (determinerAloneRule), in this order (other rules may come after it)
package parse

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceCountedShort(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	modified, det, counted := at(ruleNounModified), at(ruleDeterminerAlone), at(ruleCountedShort)
	if modified < 0 || det <= modified || counted <= det {
		t.Fatal(modified, det, counted, len(ix.sentence))
	}
}
