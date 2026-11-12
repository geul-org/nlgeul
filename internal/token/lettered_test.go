//ff:func feature=token type=helper control=sequence
//ff:what TestLettered — a place with a Latin letter or a digit on both sides (it's, 9'9, é'a); not at an edge ('90s) or next to a space, and not with a Hangul letter on either side (호랑이'는, a'는); end to end the quote of 곰이 '사자 호랑이'는 싫다 opens and closes with 는 on it, and it's has no quote
package token

import (
	"strings"
	"testing"
)

func TestLettered(t *testing.T) {
	r := []rune("a'b 'c")
	if !lettered(r, 1) || lettered(r, 4) || lettered(r, 0) || lettered(r, 5) {
		t.Fatal("lettered")
	}
	if !lettered([]rune("it's"), 2) || !lettered([]rune("9'9"), 1) || !lettered([]rune("é'a"), 1) || lettered([]rune("'90s"), 0) {
		t.Fatal("Latin letters and digits")
	}
	if lettered([]rune("이'는"), 1) || lettered([]rune("a'는"), 1) || lettered([]rune("이'a"), 1) {
		t.Fatal("a Hangul letter on either side")
	}
	bare, _ := New(Options{})
	got := strings.Join(bare.Lines("곰이 '사자 호랑이'는 싫다", false), "\n")
	if !strings.Contains(got, "'[여는 작은따옴표]") || !strings.Contains(got, "'[닫는 작은따옴표]+는[보조사:는]") {
		t.Fatal("a quote closed with a particle on it\n" + got)
	}
	if got := strings.Join(bare.Lines("it's", false), "\n"); strings.Contains(got, "작은따옴표") {
		t.Fatal("an apostrophe inside a Latin word\n" + got)
	}
}
