//ff:func feature=rule type=helper control=sequence
//ff:what Follows — 조건식(kopas parse의 읽기 코드가 쓴다; 자리 번호만 본다. 갈래 표는 Links). 조사 사슬에서 자리 left의 조사 뒤에 자리 right의 조사가 올 수 있는가
package rule

// Places rise from left to right. A case particle (3) takes nothing
// of place 3 or 4 after it, a place-4 particle only 요 (5), and an
// adverbial particle (1) may follow a place-2 one (만+으로, 까지+에).
func Follows(left, right int) bool {
	if left == 2 && right == 1 {
		return true
	}
	if left >= 3 && right <= 4 {
		return false
	}
	return right >= left
}
