//ff:func feature=data type=reader control=sequence
//ff:what TestSets — the members of each set are listed in the order of their lines and a set no line names has an empty list; a table that is not there and a set outside the names given are errors, the second one listing the names
package data

import (
	"strings"
	"testing"
)

func TestSets(t *testing.T) {
	useTables(t, map[string]string{
		"ok":  head + "집합, 낱말\n가\t하나\n나\t둘\n가\t셋\n",
		"one": head + "집합\n가\n",
		"bad": head + "집합, 낱말\n다\t하나\n",
	})
	got, err := Sets("ok", "가", "나", "비어 있음")
	if err != nil || len(got) != 3 || strings.Join(got["가"], ",") != "하나,셋" || strings.Join(got["나"], ",") != "둘" {
		t.Fatal(got, err)
	}
	if list, ok := got["비어 있음"]; !ok || len(list) != 0 {
		t.Fatal(got)
	}
	if _, err := Sets("one", "가"); err == nil || !strings.Contains(err.Error(), "표 one.tsv") {
		t.Fatal(err)
	}
	if _, err := Sets("bad", "가", "나"); err == nil || !strings.Contains(err.Error(), "표 bad.tsv 2번째 줄 1번째 칸(집합)") || !strings.Contains(err.Error(), "쓸 수 있는 이름: 가, 나") {
		t.Fatal(err)
	}
}
