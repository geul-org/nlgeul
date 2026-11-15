//ff:func feature=token type=reader control=sequence
//ff:what TestCaseExceptionOf — a line of the case exception table gives the graph, the word, the reading (none for -), the kind of its ground and the ground, and keeps the row; a strength other than 지움 and a kind of ground other than 규범 and 관찰 are errors that name the column and list the names
package token

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/data"
)

func TestCaseExceptionOf(t *testing.T) {
	names := []string{"그래프", "낱말", "바른 읽기", "세기", "근거 종류", "근거"}
	row := data.Row{File: "c.tsv", Line: 2, Names: names, Cells: []string{"맺음:소", "담소", "담소[열쇠:체언]", "지움", "관찰", "46/46"}}
	c, err := caseExceptionOf(row)
	if err != nil || c.Graph != "맺음:소" || c.Word != "담소" || c.Reading != "담소[열쇠:체언]" || c.GroundKind != "관찰" || c.Ground != "46/46" || c.Row.Line != 2 {
		t.Fatal(c, err)
	}
	row.Cells = []string{"맺음:고", "그리고", "-", "지움", "규범", "조항"}
	if c, err = caseExceptionOf(row); err != nil || c.Reading != "" || c.GroundKind != "규범" {
		t.Fatal(c, err)
	}
	row.Cells = []string{"맺음:소", "담소", "-", "낮춤", "관찰", "46/46"}
	if _, err = caseExceptionOf(row); err == nil || !strings.Contains(err.Error(), "표 c.tsv 2번째 줄 4번째 칸(세기)") || !strings.Contains(err.Error(), "쓸 수 있는 이름: 지움") {
		t.Fatal(err)
	}
	row.Cells = []string{"맺음:소", "담소", "-", "지움", "사전", "46/46"}
	if _, err = caseExceptionOf(row); err == nil || !strings.Contains(err.Error(), "5번째 칸(근거 종류)") || !strings.Contains(err.Error(), "쓸 수 있는 이름: 관찰, 규범") {
		t.Fatal(err)
	}
}
