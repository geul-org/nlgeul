//ff:func feature=data type=reader control=sequence
//ff:what TestMap — a table of two columns is a map from the first column to the second; a table that is not there, one with another number of columns and a first column written on two lines are errors
package data

import (
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	useTables(t, map[string]string{
		"ok":  head + "꼴, 뜻\n가\t하나\n나\t둘\n",
		"one": head + "꼴\n가\n",
		"dup": head + "꼴, 뜻\n가\t하나\n가\t둘\n",
	})
	got, err := Map("ok")
	if err != nil || len(got) != 2 || got["가"] != "하나" || got["나"] != "둘" {
		t.Fatal(got, err)
	}
	if _, err := Map("one"); err == nil || !strings.Contains(err.Error(), "표 one.tsv 1번째 줄") {
		t.Fatal(err)
	}
	if _, err := Map("dup"); err == nil || !strings.Contains(err.Error(), "표 dup.tsv 3번째 줄 1번째 칸(꼴)") || !strings.Contains(err.Error(), "앞 줄에 이미 적혀 있다") {
		t.Fatal(err)
	}
}
