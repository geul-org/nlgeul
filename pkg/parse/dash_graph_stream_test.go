//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestDashGraphStream — end to end, with no word table: a row of hyphens or box-drawing lines written for a dash is 줄표 alone, at a word's edge and between letters (the letters around it are read as words, and the reading that keeps the row in a word stays); one hyphen alone is 줄표 and then 빼기, and 빼기 and then 줄표 between two numbers and inside a bracket; a rule across the page, a table's line, a frame, an arrow, a list's mark, an option, the minus cut between two amounts and one hyphen between letters are not read as a dash
package parse

import (
	"slices"
	"strings"
	"testing"
)

func TestDashGraphStream(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in   string
		want []string
	}{
		{"그는 ─ 말했다", []string{"그[대명사:그]+는[보조사:는]", "─[줄표]", "말했+다[맺음:다]"}},
		{"갔다 ── 그리고", []string{"갔+다[맺음:다]", "──[줄표]", "그리고[접속 부사:그리고]"}},
		{"갔다 --- 그리고", []string{"갔+다[맺음:다]", "---[줄표]", "그리고[접속 부사:그리고]"}},
		{"갔다 -- 그리고", []string{"갔+다[맺음:다]", "--[줄표]", "그리고[접속 부사:그리고]"}},
		{"갔다--그리고", []string{"갔+다[맺음:다] --[줄표] 그리고[접속 부사:그리고]", "갔다--그리고"}},
		{"답이다──누구", []string{"답+이[서술격:이]+다[맺음:다] ──[줄표] 누구"}},
		{"갔다 ---그리고", []string{"갔+다[맺음:다]", "---[줄표]", "그리고[접속 부사:그리고]"}},
		{"갔다---.", []string{"갔+다[맺음:다]", "---[줄표]", ".[마침표]"}},
		{"보아.---그리고", []string{"보[보조 용언:보]+아[맺음:아]+.[마침표] ---[줄표] 그리고[접속 부사:그리고]"}},
		{"-- 그는", []string{"--[줄표]", "그[대명사:그]+는[보조사:는]"}},
		{"이름 ----- 값", []string{"이름", "-----[줄표]", "값"}},
		{"서울 - 부산", []string{"서울", "-[줄표]", "-[빼기]", "부산"}},
		{"100 - 20", []string{"100[수]", "-[빼기]", "-[줄표]", "20[수]"}},
		{"3,000원 - 1,000원", []string{"3,000[수]+원[단위:원]", "-[빼기]", "-[줄표]", "1,000[수]+원[단위:원]"}},
		{"가 × [나 - 다]", []string{"[[열기:1]", "-[빼기]", "-[줄표]", "][닫기:1]"}},
	} {
		if lines := streamLines(r, c.in); !linesInOrder(lines, c.want) {
			t.Errorf("%q: %q, want in this order %q", c.in, lines, c.want)
		}
	}
	for _, in := range []string{"---", "제목\n---\n글", "|---|---|", "| 이름 | 값 |\n|---|---|", "-----------", "+-----+-----+", "---- ----", "├────┼──────┤", "┌─┐", "a --> b", "a -> b", "- 항목 하나", "--head 를 준다", "-x 를 준다", "3,000원-1,000원", "가-나", "-고 를 붙인다", "10-20"} {
		named := func(line string) bool {
			return strings.Contains(line, "[줄표]") || strings.Contains(line, "-["+minusName+"]")
		}
		if lines := streamLines(r, in); slices.ContainsFunc(lines, named) {
			t.Errorf("%q: read as a dash or a minus: %q", in, lines)
		}
	}
}
