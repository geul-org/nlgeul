//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderLinesCopula — end to end with the word table of testdata: a run of Latin letters or a number before 야 reads a dropped copula stem (abc+(이)+야, 3+(이)+야) and not before 다 or 네, where the closing is read on the run itself; a noun before the naming 기 reads a dropped stem only when a particle follows the 기 (문제+(이)+기+도, 문제+(이)+기+는, not 문제기 by itself); without a word table abc야 reads only the particle 야 after the run, with no dropped stem
package token

import (
	"strings"
	"testing"
)

func TestReaderLinesCopula(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"abc야": "  abc[열:영숫자]+(이)[서술격:이(빠짐)]+야[맺음:야]\t성립 ",
		"3야":   "  3[수]+(이)[서술격:이(빠짐)]+야[맺음:야]\t성립 ",
		"abc다": "  abc[열:영숫자]+다[맺음:다]\t성립 ",
		"abc네": "  abc[열:영숫자]+네[맺음:네]\t성립 ",
		"문제기도": "  문제[열쇠:체언]+(이)[서술격:이(빠짐)]+기[이름:기]+도[보조사:도]\t성립 ",
		"문제기는": "  문제[열쇠:체언]+(이)[서술격:이(빠짐)]+기[이름:기]+는[보조사:는]\t성립 ",
	} {
		if got := strings.Join(r.Lines(in, false), "\n"); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	for in, not := range map[string]string{
		"abc다": "(이)[서술격:이(빠짐)]",
		"abc네": "(이)[서술격:이(빠짐)]",
		"문제기":  "(이)[서술격:이(빠짐)]+기",
	} {
		if got := strings.Join(r.Lines(in, false), "\n"); strings.Contains(got, not) {
			t.Errorf("%s: %q read\n%s", in, not, got)
		}
	}
	bare, _ := New(Options{})
	if got := strings.Join(bare.Lines("abc야", false), "\n"); !strings.Contains(got, "  abc[열:영숫자]+야[보조사:야|호격:야]\t") || strings.Contains(got, "(이)[서술격:이(빠짐)]") {
		t.Fatal("abc야 without a word table\n" + got)
	}
}
