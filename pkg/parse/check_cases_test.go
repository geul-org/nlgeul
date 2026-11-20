//ff:func feature=parse type=reader control=sequence
//ff:what TestCheckCases — a table whose graphs are all marker graphs or rules of a sentence of the index passes, and so does an empty one; a graph that is neither is an error that names its first row, the graphs taken in the order of their names
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/data"
)

func TestCheckCases(t *testing.T) {
	old := caseExceptions
	t.Cleanup(func() { caseExceptions = old })
	names := []string{"그래프", "낱말", "바른 읽기", "세기", "근거 종류", "근거"}
	row := func(line int, graph string) caseException {
		return caseException{Graph: graph, Row: data.Row{File: "c.tsv", Line: line, Names: names, Cells: []string{graph, "x", "-", "지움", "관찰", "1"}}}
	}
	ix := &index{byKey: map[rune][]*entry{'소': {{name: "맺음:소"}}, '오': {{name: "맺음:오"}}}}
	caseExceptions = map[string][]caseException{"맺음:소": {row(2, "맺음:소")}, "맺음:오": {row(3, "맺음:오")}}
	if err := ix.checkCases(); err != nil {
		t.Fatal(err)
	}
	ix.sentence = []sentenceRule{{name: "문장:시험"}}
	caseExceptions = map[string][]caseException{"맺음:소": {row(2, "맺음:소")}, "문장:시험": {row(4, "문장:시험")}}
	if err := ix.checkCases(); err != nil {
		t.Fatal("a rule of a sentence is named", err)
	}
	caseExceptions = map[string][]caseException{}
	if err := ix.checkCases(); err != nil {
		t.Fatal(err)
	}
	caseExceptions = map[string][]caseException{"맺음:소": {row(2, "맺음:소")}, "맺음:하": {row(5, "맺음:하"), row(6, "맺음:하")}, "맺음:가": {row(7, "맺음:가")}}
	if err := ix.checkCases(); err == nil || err.Error() != `표 c.tsv 7번째 줄 1번째 칸(그래프) "맺음:가": 색인에 이 이름의 표지 그래프도 문장 규칙도 없다` {
		t.Fatal(err)
	}
}
