//ff:func feature=rule type=helper control=iteration dimension=1
//ff:what FitsShape — 도움 함수(규칙 아님). 숫자 마디들이 꼴 기준에 맞는가: 마디 수, 마디마다 숫자뿐인지와 자리 수·값의 범위, 첫 마디의 머리
package rule

import (
	"strconv"
	"strings"
)

func FitsShape(parts []string, s *Shape) bool {
	if len(parts) < s.Min || (s.Max > 0 && len(parts) > s.Max) || !strings.HasPrefix(parts[0], s.Head) {
		return false
	}
	for i, p := range parts {
		v, err := strconv.Atoi(p)
		if err != nil || strings.Trim(p, "0123456789") != "" || !InRange(s.Lens, i, len(p)) || !InRange(s.Vals, i, v) {
			return false
		}
	}
	return true
}
