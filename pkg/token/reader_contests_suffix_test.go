//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderContestsSuffix — with the word table of testdata, the structure holds the copula after the suffix 들, 씩 or 끼리 (학교들이다, 마당씩이다, 우리끼리다 with a dropped stem, 학교들일까, 문제끼리인데); the subject candidate of 학교들이 stays and has no copula candidate; after 쯤 and the suffix 이 of a name no copula is read; 학교이다 and 3명당이다 keep their candidates
package token

import "testing"

func TestReaderContestsSuffix(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"학교들이다":  "열쇠:체언+접미:들+서술격:이+맺음:다",
		"마당씩이다":  "열쇠:체언+접미:씩+서술격:이+맺음:다",
		"우리끼리다":  "열쇠:체언+접미:끼리+서술격:이(빠짐)+맺음:다",
		"학교들일까":  "열쇠:체언+접미:들+서술격:이+맺음:ㄹ까",
		"문제끼리인데": "열쇠:체언+접미:끼리+서술격:이+맺음:ㄴ데",
		"학교들이":   "열쇠:체언+접미:들+주격:이",
		"학교이다":   "열쇠:체언+서술격:이+맺음:다",
		"3명당이다":  "수:+단위:명+접미:당+서술격:이+맺음:다",
	} {
		if !hasNames(r.Contests(in), want) {
			t.Errorf("%s: no candidate %s", in, want)
		}
	}
	for in, not := range map[string]string{
		"학교들이":  "열쇠:체언+접미:들+서술격:이",
		"학교쯤이다": "열쇠:체언+접미:쯤+서술격:이+맺음:다",
		"영숙이이다": "열쇠:체언+접미:이+서술격:이+맺음:다",
	} {
		if hasNames(r.Contests(in), not) {
			t.Errorf("%s: candidate %s", in, not)
		}
	}
}
