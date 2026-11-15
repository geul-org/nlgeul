//ff:func feature=data type=helper control=sequence
//ff:what TestMust — a value read without an error is given as it is; an error stops with a panic that carries the error's text
package data

import (
	"errors"
	"testing"
)

func TestMust(t *testing.T) {
	if got := Must(5, nil); got != 5 {
		t.Fatal(got)
	}
	defer func() {
		if r := recover(); r != "표 x.tsv 틀렸다" {
			t.Fatal(r)
		}
	}()
	Must(0, errors.New("표 x.tsv 틀렸다"))
	t.Fatal("no panic")
}
