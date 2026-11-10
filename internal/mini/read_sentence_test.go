//ff:func feature=mini type=reader control=sequence
//ff:what TestReadSentence — the last word is marked, 이/가 words before it are counted, and the last word's stem is kept; an empty sentence has no words
package mini

import "testing"

func TestReadSentence(t *testing.T) {
	words := readSentence("물이 얼음이 빨리 되었다")
	if len(words) != 4 || !words[3].last || words[0].last {
		t.Fatal(len(words))
	}
	s := words[0].sent
	if s.nominative != 2 || len(s.stems) != 1 || s.stems[0] != "되" || words[3].sent != s {
		t.Fatalf("%+v", s)
	}
	if words = readSentence("곰이"); words[0].sent.nominative != 0 {
		t.Fatal("the last word counted")
	}
	if readSentence("  ") != nil {
		t.Fatal("empty")
	}
}
