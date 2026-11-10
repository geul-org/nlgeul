//ff:func feature=rule type=helper control=sequence
//ff:what InRange — 도움 함수(규칙 아님). 값이 i째 범위 안인가: 범위 목록이 비면 참, i가 목록을 넘으면 마지막 범위, 위 끝 0은 한도 없음
package rule

func InRange(ranges [][2]int, i, v int) bool {
	if len(ranges) == 0 {
		return true
	}
	r := ranges[min(i, len(ranges)-1)]
	return v >= r[0] && (r[1] == 0 || v <= r[1])
}
