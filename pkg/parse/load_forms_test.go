//ff:func feature=parse type=helper control=sequence
//ff:what TestLoadForms — the lines of the tables given are read in order, a line a form with its five columns (- in the last is no number); a first column that is not one letter is an error
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/data"
)

func TestLoadForms(t *testing.T) {
	one := []data.Row{{Cells: []string{"①", "원", "수", "1", "1"}}}
	two := []data.Row{{Cells: []string{"㈎", "괄호", "한글", "가", "-"}}}
	got, err := loadForms(one, two)
	if err != nil || len(got) != 2 || got[0] != (form{'①', "원", "수", "1", "1", ""}) || got[1] != (form{'㈎', "괄호", "한글", "가", "", ""}) {
		t.Fatalf("%v %v", got, err)
	}
	if _, err := loadForms([]data.Row{{Cells: []string{"①②", "원", "수", "1", "1"}, Names: []string{"글자", "꼴", "내용", "값", "수 값"}}}); err == nil {
		t.Fatal("two letters in the first column: no error")
	}
	if len(formList) < 1000 || len(forms['①']) != 1 || forms['①'][0].name() != "원:수" {
		t.Fatalf("the tables held: %d lines, ① %v", len(formList), forms['①'])
	}
}
