//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamCopula — end to end with the word table of testdata: a run of Latin letters or a number before 야 reads a dropped copula stem (abc+(이)+야, 3+(이)+야) and not before 다 or 네, where the closing is read on the run itself; a noun before the naming 기 reads a dropped stem only when a particle follows the 기 (문제+(이)+기+도, 문제+(이)+기+는, not 문제기 by itself); without a word table abc야 reads only the particle 야 after the run, with no dropped stem
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamCopula(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"abc야": "\nabc[열:영숫자]+(이)[서술격:이(빠짐)]+야[맺음:야]\n",
		"3야":   "\n3[수]+(이)[서술격:이(빠짐)]+야[맺음:야]\n",
		"abc다": "\nabc[열:영숫자]+다[맺음:다]\n",
		"abc네": "\nabc[열:영숫자]+네[맺음:네]\n",
		"문제기도": "\n문제[열쇠:체언]+(이)[서술격:이(빠짐)]+기[이름:기]+도[보조사:도]\n",
		"문제기는": "\n문제[열쇠:체언]+(이)[서술격:이(빠짐)]+기[이름:기]+는[보조사:는]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	for in, not := range map[string]string{
		"abc다": "(이)[서술격:이(빠짐)]",
		"abc네": "(이)[서술격:이(빠짐)]",
		"문제기":  "(이)[서술격:이(빠짐)]+기",
	} {
		if got := streamText(r, in); strings.Contains(got, not) {
			t.Errorf("%s: %q read\n%s", in, not, got)
		}
	}
	bare, _ := New(Options{})
	if got := streamText(bare, "abc야"); !strings.Contains(got, "\nabc[열:영숫자]+야[보조사:야]\n") || !strings.Contains(got, "\nabc[열:영숫자]+야[호격:야]\n") || strings.Contains(got, "(이)[서술격:이(빠짐)]") {
		t.Fatal("abc야 without a word table\n" + got)
	}
}
