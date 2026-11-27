//ff:func feature=parse type=helper control=sequence
//ff:what TestFormsByLetter — the lines are filed under their letter in the order they are written; a letter with two lines keeps both
package parse

import "testing"

func TestFormsByLetter(t *testing.T) {
	got := formsByLetter([]form{{'○', "원", "가림", "○", ""}, {'①', "원", "수", "1", "1"}, {'○', "원", "글머리", "○", ""}})
	if len(got) != 2 || len(got['○']) != 2 || got['○'][0].content != "가림" || got['○'][1].content != "글머리" || got['①'][0].number != "1" {
		t.Fatalf("%v", got)
	}
}
