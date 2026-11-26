//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamPaths — the readings of a path that the letters alone do not settle, as the stream sends them: the Hangul after a path's last slash is its name first, and the path cut at the slash with a word after it and the new word of the whole letters stay behind; a particle on a directory's slash stays first before the name reading (so do a contracted particle and the copula with a closing fused into it, and the path is not cut from them as from a word); a path with no head and one slash whose name holds Hangul joined by a sign is a path first; a path's head alone is a path and a slash alone the sign; a hidden name with a slash is a path and a name with the extension bat a file name; two names joined by one slash keep the reading with no name; the readings that were there stay (words set side by side by a slash, a path with a particle, a path's name joined by a hyphen, a directory's slash, a slash before a word)
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamPaths(t *testing.T) {
	r := streamReader(t, "words.tsv")
	for _, c := range []struct {
		in, first string
		has, not  []string
	}{
		{"~/work/notes/회의록 본다", "~/work/notes/회의록[경로]", []string{"\n~/work/notes/[경로] 회의록\n", "\n~/work/notes/회의록\n"}, nil},
		{"~/work/notes/결과는 좋다", "~/work/notes/결과는[경로]", []string{"\n~/work/notes/[경로] 결과[열쇠:체언]+는[보조사:는]\n", "\n~/work/notes/결과+는[보조사:는]\n"}, nil},
		{"~/work/notes/에 둔다", "~/work/notes/[경로]+에[부사격:에]", []string{"\n~/work/notes/에[경로]\n", "\n~/work/notes/에\n"}, nil},
		{"~/work/notes/에겐 없다", "~/work/notes/[경로]+에겐[부사격:에겐]", []string{"\n~/work/notes/에겐[경로]\n"}, []string{"[경로] 에겐"}},
		{"~/work/notes/인 것", "~/work/notes/[경로]+이[서술격:이]+(ㄴ)[꾸밈:ㄴ]", []string{"\n~/work/notes/인[경로]\n"}, nil},
		{"../창고/토지 본다", "../창고/토지[경로]", []string{"\n../창고/[경로] 토지\n"}, nil},
		{"notes/회의-결과 본다", "notes/회의-결과[경로]", []string{"\nnotes/회의-결과\n"}, nil},
		{"notes/memo-결과 본다", "notes/memo-결과[경로]", []string{"\nnotes/memo-결과\n"}, nil},
		{"docs/계약서_양식 본다", "docs/계약서_양식[경로]", []string{"\ndocs/계약서_양식\n"}, nil},
		{"notes/회의-v2를 본다", "notes/회의-v2[경로]+를[목적격:를]", []string{"\nnotes/회의-v2+를[목적격:를]\n"}, nil},
		{"~/ 아래에 둔다", "~/[경로]", nil, []string{"[빗금]"}},
		{"./ 와 ../ 를 본다", "./[경로]", []string{"\n../[경로]\n"}, []string{"[빗금]"}},
		{"/ 를 쓴다", "/[빗금]", nil, []string{"[경로]"}},
		{".cache/db를 지운다", ".cache/db[경로]+를[목적격:를]", []string{"\n.cache/db+를[목적격:를]\n"}, nil},
		{"run.bat 연다", "run.bat[파일명]", nil, nil},
		{"ab/cd를 연다", "ab/cd[경로]+를[목적격:를]", []string{"\nab/cd+를[목적격:를]\n"}, []string{"[빗금]"}},
		{"입력/출력을 본다", "입력/출력[대비:/]+을[목적격:을]", []string{"\n입력[열쇠:체언] /[빗금] 출력[열쇠:체언]+을[목적격:을]\n"}, []string{"[경로]"}},
		{"~/a/b를 본다", "~/a/b[경로]+를[목적격:를]", nil, []string{"b를[경로]"}},
		{"~/a/b-가격표 본다", "~/a/b-가격표[경로]", nil, nil},
		{"pkg/parser/ 를 본다", "pkg/parser/[경로]", nil, nil},
		{"files/ 를 본다", "files[열:영숫자]", []string{"\n/[빗금]\n"}, []string{"[경로]"}},
		{"/창고 본다", "/창+고[맺음:고]", []string{"\n/창고\n"}, []string{"[경로]"}},
		{"TV/라디오 광고", "TV/라디오", nil, []string{"[경로]"}},
	} {
		lines := streamLines(r, c.in)
		if len(lines) == 0 || lines[0] != c.first {
			t.Errorf("%s: first %q", c.in, lines)
		}
		checkLines(t, c.in, "\n"+strings.Join(lines, "\n")+"\n", c.has, c.not)
	}
}
