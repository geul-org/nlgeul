//ff:func feature=data type=reader control=iteration dimension=1
//ff:what Row.Flags — the values of a column that holds names of a closed set joined with |, summed as bits; - is none (0); a name that is not of the set, or one written twice, is an error
package data

import "strings"

func (r Row) Flags(col int, names map[string]int) (int, error) {
	out := 0
	if r.Cells[col] == none {
		return 0, nil
	}
	for _, name := range strings.Split(r.Cells[col], joint) {
		v, ok := names[name]
		if !ok {
			return 0, r.Errf(col, "\"%s\"은 정해진 이름이 아니다. 쓸 수 있는 이름: %s", name, listOf(names))
		}
		if out&v != 0 {
			return 0, r.Errf(col, "\"%s\"이 두 번 적혀 있다", name)
		}
		out |= v
	}
	return out, nil
}
