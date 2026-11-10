//ff:func feature=mini type=reader control=sequence
//ff:what follows — may a particle of place right stand after one of place left (particles.md section 5)
package mini

// Places rise from left to right. A case particle (3) takes nothing
// of place 3 or 4 after it, a place-4 particle only 요 (5), and an
// adverbial particle (1) may follow a place-2 one (만+으로, 까지+에).
func follows(left, right int) bool {
	if left == 2 && right == 1 {
		return true
	}
	if left >= 3 && right <= 4 {
		return false
	}
	return right >= left
}
