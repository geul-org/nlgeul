//ff:func feature=token type=reader control=sequence
//ff:what TestReaderStreamApostrophe — end to end the quote of 곰이 '사자 호랑이'는 싫다 opens and closes with 는 on it, and it's has no quote
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamApostrophe(t *testing.T) {
	bare, _ := New(Options{})
	got := streamText(bare, "곰이 '사자 호랑이'는 싫다")
	if !strings.Contains(got, "'[열기:1]") || !strings.Contains(got, "\n'[닫기:1]\n") || !strings.Contains(got, "\n+는[보조사:는]\n") || !strings.HasPrefix(streamScore(bare, "곰이 '사자 호랑이'는 싫다", "+는[보조사:는]"), "1.00 ") {
		t.Fatal("a quote closed with a particle on it\n" + got)
	}
	if got := streamText(bare, "it's"); strings.Contains(got, "[열기") || strings.Contains(got, "[닫기") {
		t.Fatal("an apostrophe inside a Latin word\n" + got)
	}
}
