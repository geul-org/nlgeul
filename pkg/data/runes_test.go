//ff:func feature=data type=reader control=sequence
//ff:what TestRunes — the first column of letters is given line by line with each letter once, in the order it is first written; a table that is not there and a first column of more than one letter are errors
package data

import (
	"strings"
	"testing"
)

func TestRunes(t *testing.T) {
	useTables(t, map[string]string{
		"ok":  head + "기호, 뜻\n,\t1\n.\t2\n,\t3\n",
		"bad": head + "기호, 뜻\n,,\t1\n",
	})
	got, err := Runes("ok", 2)
	if err != nil || string(got) != ",." {
		t.Fatal(string(got), err)
	}
	if _, err := Runes("none", 2); err == nil || !strings.Contains(err.Error(), "표 none.tsv") {
		t.Fatal(err)
	}
	if _, err := Runes("bad", 2); err == nil || !strings.Contains(err.Error(), "표 bad.tsv 2번째 줄 1번째 칸(기호)") || !strings.Contains(err.Error(), "글자 하나가 아니다") {
		t.Fatal(err)
	}
}
