//ff:func feature=data type=helper control=sequence
//ff:what TestListOf — the names of a set are listed in the order of their letters, joined with commas
package data

import "testing"

func TestListOf(t *testing.T) {
	if got := listOf(map[string]int{"다": 3, "가": 1, "나": 2}); got != "가, 나, 다" {
		t.Fatal(got)
	}
	if got := listOf(map[string]int{}); got != "" {
		t.Fatal(got)
	}
}
