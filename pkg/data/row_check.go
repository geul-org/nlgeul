//ff:func feature=data type=reader control=iteration dimension=1
//ff:what Row.check — a row has as many columns as the table names, none of them empty, and is not the line written before it at the line given (0 for none)
package data

import (
	"fmt"
	"strings"
)

func (r Row) check(before int) error {
	if len(r.Cells) != len(r.Names) {
		return fmt.Errorf("표 %s %d번째 줄: 칸이 %d개다. 칸 %d개(%s)가 탭으로 나뉘어 있어야 한다", r.File, r.Line, len(r.Cells), len(r.Names), strings.Join(r.Names, ", "))
	}
	if before != 0 {
		return fmt.Errorf("표 %s %d번째 줄: %d번째 줄과 같은 줄이다", r.File, r.Line, before)
	}
	for i, cell := range r.Cells {
		if cell == "" {
			return r.Errf(i, "빈 칸이다. 값이 없으면 %s를 적는다", none)
		}
	}
	return nil
}
