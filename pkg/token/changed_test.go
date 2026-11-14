//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestChanged — each kind of change gives its form: 살 → 사 with no type, 듣 → 들, 짓 → 지, 덥 → 더우, 돕 → 도우 and 도오, 파랗 → 파라 and 파래, 그렇 → 그러 and 그래, 모르 → 몰르, 이르 → 이르르, 푸 → 프; a regular stem that ends in no ㄹ has none; a type the stem's shape does not fit gives nothing
package token

import (
	"fmt"
	"testing"
)

func TestChanged(t *testing.T) {
	cases := map[string]string{
		fmt.Sprint(changed("살", 0)):                                      fmt.Sprint([]alt{{"사", irrL}}),
		fmt.Sprint(changed("만들", 0)):                                     fmt.Sprint([]alt{{"만드", irrL}}),
		fmt.Sprint(changed("듣", irrD)):                                   fmt.Sprint([]alt{{"들", irrD}}),
		fmt.Sprint(changed("짓", irrS)):                                   fmt.Sprint([]alt{{"지", irrS}}),
		fmt.Sprint(changed("덥", irrB)):                                   fmt.Sprint([]alt{{"더우", irrB}}),
		fmt.Sprint(changed("돕", irrB)):                                   fmt.Sprint([]alt{{"도우", irrB}, {"도오", irrBO}}),
		fmt.Sprint(changed("파랗", irrH)):                                  fmt.Sprint([]alt{{"파라", irrH}, {"파래", irrHE}}),
		fmt.Sprint(changed("그렇", irrH)):                                  fmt.Sprint([]alt{{"그러", irrH}, {"그래", irrHE}}),
		fmt.Sprint(changed("누렇", irrH)):                                  fmt.Sprint([]alt{{"누러", irrH}, {"누레", irrHE}}),
		fmt.Sprint(changed("모르", irrReu)):                                fmt.Sprint([]alt{{"몰르", irrReu}}),
		fmt.Sprint(changed("이르", irrReo)):                                fmt.Sprint([]alt{{"이르르", irrReo}}),
		fmt.Sprint(changed("푸", irrU)):                                   fmt.Sprint([]alt{{"프", irrU}}),
		fmt.Sprint(changed("먹", 0)):                                      fmt.Sprint([]alt(nil)),
		fmt.Sprint(changed("먹", irrD|irrS|irrB|irrH|irrReu|irrReo|irrU)): fmt.Sprint([]alt(nil)),
		fmt.Sprint(changed("르", irrReu)):                                 fmt.Sprint([]alt(nil)),
	}
	for got, want := range cases {
		if got != want {
			t.Fatal(got, want)
		}
	}
}
