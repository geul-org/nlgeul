//ff:func feature=parse type=rule control=sequence
//ff:what TestAddWideSuffix — addWide puts the suffixes read after a nominal in the index with the graph the word before them must pass: 율 in the place of a suffix, with no count, stands after a vowel or ㄴ and not after another final consonant; 당 in the place of a suffix that ends its word stands at its word's end and not before a particle
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddWideSuffix(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addParticles()
	ix.addWide()
	yul := ix.named('율', "접미:율")
	if yul.gate != ix.suffixGate || yul.g == nil || yul.class != classParticle || yul.slot != rule.LinkSuffix || yul.after != rule.AfterOpenN || yul.count != 0 {
		t.Fatal("the suffix 율 after a nominal", yul)
	}
	if v, _ := verdictOf(yul.g, rule.Piece{Form: "율", HasStem: true, Prev: '가', Next: -1}, "증가율", false); v <= 0 {
		t.Fatal("율 after a vowel", v)
	}
	if v, _ := verdictOf(yul.g, rule.Piece{Form: "율", HasStem: true, Prev: '산', Next: -1}, "출산율", false); v <= 0 {
		t.Fatal("율 after ㄴ", v)
	}
	if v, _ := verdictOf(yul.g, rule.Piece{Form: "율", HasStem: true, Prev: '장', Next: -1}, "보장율", false); v > 0 {
		t.Fatal("율 after a final consonant other than ㄴ", v)
	}
	dang := ix.named('당', "접미:당")
	if dang.gate != ix.suffixGate || dang.g == nil || dang.class != classParticle || dang.slot != rule.LinkLast || dang.after != rule.AfterAny || dang.count != 0 {
		t.Fatal("the suffix 당 after a nominal", dang)
	}
	if v, _ := verdictOf(dang.g, rule.Piece{Form: "당", HasStem: true, Prev: '팀', Slot: rule.LinkLast, Next: -1}, "팀당", false); v <= 0 {
		t.Fatal("당 at its word's end", v)
	}
	if v, _ := verdictOf(dang.g, rule.Piece{Form: "당", HasStem: true, Prev: '팀', Slot: rule.LinkLast, Next: rule.LinkAdv}, "팀당", false); v > 0 {
		t.Fatal("당 before a particle", v)
	}
}
