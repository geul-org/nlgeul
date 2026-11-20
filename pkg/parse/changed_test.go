//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestChanged — each kind of change gives its form: 살 → 사 with no type, 듣 → 들, 짓 → 지, 덥 → 더우, 돕 → 도우 and 도오, 파랗 → 파라 and 파래, 그렇 → 그러 and 그래, 모르 → 몰르, 이르 → 이르르, 푸 → 프, 원하 → 원 and 가능하 → 가능 with 하 dropped (clause 40), but not the one-syllable 하; 생각하 → 생각, 깨끗하 → 깨끗, 섭섭하 → 섭섭 also with 하 dropped wholly after a final consonant that sounds [ㄱ], [ㄷ], [ㅂ] (note 2), but not 가능하 (ㅇ) nor 깔끔하 (ㅁ) nor 넓하 (ㄼ, [ㄹ]); a regular stem that ends in no ㄹ has none; a type the stem's shape does not fit gives nothing
package parse

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
		fmt.Sprint(changed("원하", 0)):                                     fmt.Sprint([]alt{{"원", irrHa}}),
		fmt.Sprint(changed("가능하", 0)):                                    fmt.Sprint([]alt{{"가능", irrHa}}),
		fmt.Sprint(changed("하", 0)):                                      fmt.Sprint([]alt(nil)),
		fmt.Sprint(changed("생각하", 0)):                                    fmt.Sprint([]alt{{"생각", irrHa}, {"생각", irrHaW}}),
		fmt.Sprint(changed("깨끗하", 0)):                                    fmt.Sprint([]alt{{"깨끗", irrHa}, {"깨끗", irrHaW}}),
		fmt.Sprint(changed("섭섭하", 0)):                                    fmt.Sprint([]alt{{"섭섭", irrHa}, {"섭섭", irrHaW}}),
		fmt.Sprint(changed("깔끔하", 0)):                                    fmt.Sprint([]alt{{"깔끔", irrHa}}),
		fmt.Sprint(changed("넓하", 0)):                                     fmt.Sprint([]alt{{"넓", irrHa}}),
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
