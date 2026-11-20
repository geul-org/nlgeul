//ff:func feature=parse type=reader control=sequence
//ff:what TestLoadCases — the rows are kept graph by graph in the order of their lines; a row read wrong is its error, and a word written for one graph on two lines is an error that names the line before (the same word for another graph is not)
package parse

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/data"
)

func TestLoadCases(t *testing.T) {
	names := []string{"그래프", "낱말", "바른 읽기", "세기", "근거 종류", "근거"}
	row := func(line int, cells ...string) data.Row {
		return data.Row{File: "c.tsv", Line: line, Names: names, Cells: cells}
	}
	a := row(2, "맺음:소", "담소", "-", "지움", "관찰", "1")
	b := row(3, "맺음:오", "가오", "-", "지움", "관찰", "2")
	c := row(5, "맺음:소", "적소", "-", "지움", "규범", "3")
	d := row(6, "맺음:오", "담소", "-", "지움", "관찰", "4")
	got, err := loadCases([]data.Row{a, b, c, d})
	if err != nil || len(got) != 2 || len(got["맺음:소"]) != 2 || got["맺음:소"][0].Word != "담소" || got["맺음:소"][1].Word != "적소" || got["맺음:소"][1].Row.Line != 5 || len(got["맺음:오"]) != 2 {
		t.Fatal(got, err)
	}
	if _, err := loadCases([]data.Row{a, row(4, "맺음:소", "담소", "-", "낮춤", "관찰", "1")}); err == nil || !strings.Contains(err.Error(), "4번째 칸(세기)") {
		t.Fatal(err)
	}
	if _, err := loadCases([]data.Row{a, b, row(7, "맺음:소", "담소", "-", "지움", "관찰", "9")}); err == nil || !strings.Contains(err.Error(), "표 c.tsv 7번째 줄 2번째 칸(낱말)") || !strings.Contains(err.Error(), "2번째 줄이 같은 그래프에 이 낱말을 이미 적었다") {
		t.Fatal(err)
	}
	if got, err := loadCases(nil); err != nil || len(got) != 0 {
		t.Fatal(got, err)
	}
}
