//ff:func feature=mini type=reader control=iteration dimension=1
//ff:what TestAffixMark — 하, 되, 시키 are set off from a word before them, fused or not; alone, or after nothing, they are not; other stems stay whole
package mini

import "testing"

func TestAffixMark(t *testing.T) {
	cases := []struct {
		stem  string
		fused bool
		want  string
	}{
		{"지시해", true, "지시+해"},
		{"작성하", false, "작성+하"},
		{"시작됩", true, "시작+됩"},
		{"발전시키", false, "발전+시키"},
		{"가셨다고합", true, "가+시+었+다+고+합"},
		{"했", true, "했"},
		{"시키", false, "시키"},
		{"먹", false, "먹"},
		{"고팠", true, "고팠"},
		{"학", false, "학"},
	}
	for _, c := range cases {
		if got := affixMark([]rune(c.stem), c.fused); got != c.want {
			t.Errorf("%s: %s", c.stem, got)
		}
	}
}
