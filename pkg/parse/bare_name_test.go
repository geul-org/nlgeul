//ff:func feature=parse type=helper control=sequence
//ff:what TestBareName — a candidate with a name keeps it; one with none takes the first name the tables of special letters give the sign (= is 기호:같음), and stays with none for a sign those tables lack
package parse

import "testing"

func TestBareName(t *testing.T) {
	if got := bareName("가운뎃점", '·'); got != "가운뎃점" {
		t.Fatal(got)
	}
	if got := bareName("", '='); got != "기호:같음" {
		t.Fatal(got)
	}
	if got := bareName("빗금", '='); got != "빗금" {
		t.Fatal(got)
	}
	if got := bareName("", '가'); got != "" {
		t.Fatal(got)
	}
}
