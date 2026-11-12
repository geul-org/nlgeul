//ff:func feature=data type=reader control=sequence
//ff:what TestCounts — a table of a name and a count is a map; a table that is not there, a count that is not a number and a name written on two lines are errors that name the line
package data

import (
	"strings"
	"testing"
)

func TestCounts(t *testing.T) {
	useTables(t, map[string]string{
		"ok":  head + "이름, 횟수\n가\t5\n나\t0\n",
		"one": head + "이름\n가\n",
		"bad": head + "이름, 횟수\n가\t많이\n",
		"dup": head + "이름, 횟수\n가\t1\n가\t2\n",
	})
	got, err := Counts("ok")
	if err != nil || len(got) != 2 || got["가"] != 5 || got["나"] != 0 {
		t.Fatal(got, err)
	}
	if _, err := Counts("one"); err == nil || !strings.Contains(err.Error(), "표 one.tsv") {
		t.Fatal(err)
	}
	if _, err := Counts("bad"); err == nil || !strings.Contains(err.Error(), "표 bad.tsv 2번째 줄 2번째 칸(횟수)") {
		t.Fatal(err)
	}
	if _, err := Counts("dup"); err == nil || !strings.Contains(err.Error(), "표 dup.tsv 3번째 줄 1번째 칸(이름)") || !strings.Contains(err.Error(), "앞 줄에 이미 적혀 있다") {
		t.Fatal(err)
	}
}
