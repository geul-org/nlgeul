//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNounCopulasLines — end to end with the word table of testdata, which has no 터 or 테 (the stem is printed with the form's name, 줄임:테): 테 after a modifying form written apart (할 테니까, 갈 텐데, 할 테고) or onto it (갈테니, 갈텐데, 할테야) reads 터 with the copula's stem (이) before the closing, as a word that does not attach; a word that begins with 테 (테스트, 테이블), 테 after a word that ends in no modifying closing (나테니, 곰테니), before ㅆ fused into its vowel (텠다) or before a closing the copula does not take (테자) has no such candidate; without a word table 할 테니까 still reads it
package token

import (
	"strings"
	"testing"
)

func TestNounCopulasLines(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"할 테니까": "\n  터[줄임:테]+(이)[서술격:이]+니까[맺음:니까]\t성립 ",
		"갈 텐데":  "\n  터[줄임:테]+(이)[서술격:이]+(ㄴ)데[맺음:ㄴ데]\t성립 ",
		"할 테고":  "\n  터[줄임:테]+(이)[서술격:이]+고[맺음:고]\t성립 ",
		"갈테니":   "\n  가[열쇠:동사]+(ㄹ)[꾸밈:ㄹ] 터[줄임:테]+(이)[서술격:이]+니[맺음:니]\t성립 ",
		"갈텐데":   "\n  가[열쇠:동사]+(ㄹ)[꾸밈:ㄹ] 터[줄임:테]+(이)[서술격:이]+(ㄴ)데[맺음:ㄴ데]\t성립 ",
		"할테야":   "\n  하[열쇠:동사]+(ㄹ)[꾸밈:ㄹ] 터[줄임:테]+(이)[서술격:이]+야[맺음:야]\t성립 ",
	} {
		if got := strings.Join(r.Lines(in, false), "\n"); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	for _, in := range []string{"테스트", "테이블", "갈 테이블", "나테니", "곰테니", "할 텠다", "할 테자"} {
		if got := strings.Join(r.Lines(in, false), "\n"); strings.Contains(got, "줄임:테") {
			t.Errorf("%s: 터 read\n%s", in, got)
		}
	}
	bare, _ := New(Options{})
	if got := strings.Join(bare.Lines("할 테니까", false), "\n"); !strings.Contains(got, "\n  터[줄임:테]+(이)[서술격:이]+니까[맺음:니까]\t성립 ") {
		t.Fatal("할 테니까 without a word table\n" + got)
	}
}
