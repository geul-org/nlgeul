//ff:func feature=data type=reader control=iteration dimension=1
//ff:what Sets — a table of two columns, the name of a set and a member, as the members of each set in the order of their lines; a set that is not among the names given is an error, and every name given has a list (empty when no line names it)
package data

import (
	"slices"
	"strings"
)

func Sets(name string, sets ...string) (map[string][]string, error) {
	rows, err := Rows(name, 2)
	if err != nil {
		return nil, err
	}
	out := map[string][]string{}
	for _, set := range sets {
		out[set] = nil
	}
	for _, row := range rows {
		if !slices.Contains(sets, row.Cells[0]) {
			return nil, row.Errf(0, "정해진 이름이 아니다. 쓸 수 있는 이름: %s", strings.Join(sets, ", "))
		}
		out[row.Cells[0]] = append(out[row.Cells[0]], row.Cells[1])
	}
	return out, nil
}
