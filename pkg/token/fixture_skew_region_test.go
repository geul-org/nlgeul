//ff:func feature=token type=helper control=sequence
//ff:what skewRegion — a region of five letters made by hand from two ways, [0-2][2-3][3-4][4-5] (preference -4 in all) and [0-1][1-2][2-3][3-5] (-3.5), which share the word [2-3]; so that the words can be joined into two ways that are neither ([0-2][2-3][3-5] at -2.5, the best, and [0-1][1-2][2-3][3-4][4-5] at -5). Its nodes by place: 0 [0-1], 1 [0-2], 2 [1-2], 3 [2-3], 4 [3-4], 5 [3-5], 6 [4-5]
package token

func skewRegion() *region {
	first := streamCut([]Span{{0, 2}, {2, 3}, {3, 4}, {4, 5}}, []string{"a", "c", "d", "e"}, []float64{-1, -1, -1, -1})
	second := streamCut([]Span{{0, 1}, {1, 2}, {2, 3}, {3, 5}}, []string{"f", "g", "c", "h"}, []float64{-1, -1, -1, -0.5})
	return newRegion(Contest{Unit: 3, Glued: true, Text: "가나다라마"}, 0, []cut{first, second})
}
