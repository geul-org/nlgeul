//ff:func feature=token type=reader control=sequence
//ff:what pref — a marker's preference: log10 of its share of the tables' observed counts; a marker the tables give no count is neither preferred nor not (it costs what its syllables would cost as stem syllables)
package token

import "math"

func (ix *index) pref(count int, form string) float64 {
	if count == 0 {
		return float64(len([]rune(form))) * math.Log10(1.0/syllables)
	}
	return math.Log10(float64(count) / ix.total)
}
