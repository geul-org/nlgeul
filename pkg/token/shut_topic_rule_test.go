//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestShutTopicRule — the rule is named as the trace prints it, adds bonusShutTopic (more than bonusSentence), waits for nothing, and speaks only of a candidate that has the particle 이란 or 란 among its markers (not of 란 read as the modifying form of the copula, nor of the noun 이란, nor of a word with no marker)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestShutTopicRule(t *testing.T) {
	r := shutTopicRule()
	if r.name != ruleShutTopic || r.bonus != bonusShutTopic || bonusShutTopic <= bonusSentence || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Last: "보조사:이란", Marks: "보조사:이란"}, true},
		{rule.Near{Last: "보조사:란", Marks: "보조사:란"}, true},
		{rule.Near{Stem: "사랑", Kind: "열쇠:체언", Last: "보조사:이란", Marks: "보조사:이란"}, true},
		{rule.Near{Last: "꾸밈:란(라는)", Marks: "서술격:이+꾸밈:란(라는)"}, false},
		{rule.Near{Stem: "이란", Kind: "열쇠:체언"}, false},
		{rule.Near{Last: "부사격:라고", Marks: "부사격:라고"}, false},
	} {
		if got := r.about(c.near); got != c.want {
			t.Errorf("about(%+v) = %v", c.near, got)
		}
	}
}
