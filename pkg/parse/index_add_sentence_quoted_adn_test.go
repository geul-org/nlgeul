//ff:func feature=parse type=rule control=sequence
//ff:what TestIndexAddSentenceQuotedAdn — the index of a reader holds the rule of a sentence ruleQuotedAdn right after the rule of the topic particle after a closing sign (shutTopicRule)
package parse

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceQuotedAdn(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	topic, adn := at(ruleShutTopic), at(ruleQuotedAdn)
	if topic < 0 || adn != topic+1 {
		t.Fatal(topic, adn, len(ix.sentence))
	}
}
