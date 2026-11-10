//ff:func feature=mini type=reader control=iteration dimension=1
//ff:what TestFollows — 에+는, 에서+부터, 만+으로, 까지+에, 에서+와, 는+요 are allowed; 이+는, 는+도, 이+에, 는+에, 요+는 are not
package mini

import "testing"

func TestFollows(t *testing.T) {
	cases := []struct {
		left, right int
		want        bool
	}{
		{1, 4, true}, {1, 2, true}, {2, 1, true}, {1, 1, true}, {4, 5, true}, {3, 5, true}, {0, 1, true}, {2, 3, true},
		{3, 4, false}, {4, 4, false}, {3, 1, false}, {4, 1, false}, {5, 4, false}, {3, 3, false}, {4, 2, false},
	}
	for _, c := range cases {
		if follows(c.left, c.right) != c.want {
			t.Errorf("follows(%d, %d)", c.left, c.right)
		}
	}
}
