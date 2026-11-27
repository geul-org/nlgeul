//ff:func feature=parse type=helper control=sequence
//ff:what TestOrderJoint — a tail's first piece is the prefix 제 of a number where it is the letter 제 and a digit or an enclosed number is right after it (요추제1, 갑제13, 제①); not where a Hangul syllable follows (제가), nothing follows, the piece is another letter (용) or the tail has no piece
package parse

import "testing"

func TestOrderJoint(t *testing.T) {
	sc := wordsScan("요추제1 갑제① 연봉제가 연봉제", layerCost)
	je := []piece{{text: "제", names: "접미:제"}}
	if !sc.orderJoint(tail{start: 2, pieces: je}) || !sc.orderJoint(tail{start: 6, pieces: je}) {
		t.Fatal("a number right after 제")
	}
	if sc.orderJoint(tail{start: 11, pieces: je}) || sc.orderJoint(tail{start: 16, pieces: je}) {
		t.Fatal("a syllable or nothing after 제")
	}
	if sc.orderJoint(tail{start: 2, pieces: []piece{{text: "용", names: "접미:용"}}}) || sc.orderJoint(tail{start: 2}) {
		t.Fatal("another letter or no piece")
	}
}
