//ff:func feature=token type=reader control=sequence
//ff:what TestLines — stream-graph-algorithm.md section 4: 다음으로써 is sent up as a contest whose first two branches are 다음+으로써 and 다음+으로 써; spaces are fixed boundaries; the same unit always gives the same lines
package token

import (
	"strings"
	"testing"
)

func TestLines(t *testing.T) {
	got := Lines("다음으로써", false)
	if len(got) < 3 || !strings.HasPrefix(got[0], "다음으로써\t경합(") || !strings.HasPrefix(got[1], "  다음+으로써[부사격:으로써]\t") || !strings.HasPrefix(got[2], "  다음+으로[부사격:으로] 써\t") {
		t.Fatal(got)
	}
	two := Lines("아버지가 아버지가", false)
	half := len(two) / 2
	if half == 0 || strings.Join(two[:half], "\n") != strings.Join(two[half:], "\n") {
		t.Fatal(two)
	}
	if traced := Lines("다음으로써", true); !strings.Contains(strings.Join(traced, "\n"), "맺음:다\t-1.00\t기각: 앞에 줄기가 없다") {
		t.Fatal(traced)
	}
	if Lines("", false) != nil {
		t.Fatal("empty")
	}
}
