//ff:func feature=token type=reader control=sequence
//ff:what TestMarked — after a boundary a stem whose letters read as markers there by themselves (에 before 서), or from the link vowel 으 on or after it up to the tail's end (으며, 으셨다), is marked; so is a word with no marker after it that reads as one run of markers or as two runs one after the other (는 and then 게); the 으 left of 을 when ㄹ is read as a piece is the link vowel, while a fused stem of two syllables is judged by its first letter as written; not marked are the unit's first word, letters that do not so read, markers that end before the tail's end, a tail that names its own stem, and a stem a piece is fused into outside the link vowel; the places are those the search fills
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMarked(t *testing.T) {
	seo := []piece{{text: "서", names: "부사격:서", verdict: 1, tail: rule.TailCase}}
	sc := scanOf("곰에서")
	sc.marks[2] = map[int]bool{1: true}
	if !sc.marked(1, 3, tail{start: 2, pieces: seo}) {
		t.Fatal("에 after a boundary, read as a marker, before 서")
	}
	if sc.marked(0, 3, tail{start: 2, pieces: seo}) {
		t.Fatal("the unit's first word")
	}
	if sc.marked(1, 3, tail{start: 2, only: "에", pieces: seo}) {
		t.Fatal("a tail that names its own stem")
	}
	if sc.marked(1, 3, tail{start: 2, base: '에', pieces: seo}) {
		t.Fatal("a stem a piece is fused into, which does not start with the link vowel")
	}
	sc.marks[2] = map[int]bool{}
	if sc.marked(1, 3, tail{start: 2, pieces: seo}) {
		t.Fatal("letters that do not read as markers")
	}

	whole := scanOf("한다는게")
	if whole.marked(2, 4, tail{start: 4}) {
		t.Fatal("a word with no marker after it whose letters read as no markers")
	}
	whole.marks[4] = map[int]bool{2: true}
	if !whole.marked(2, 4, tail{start: 4}) || whole.marked(0, 4, tail{start: 4}) {
		t.Fatal("a word with no marker after it that reads as one run of markers, after a boundary only")
	}
	whole.marks[4] = map[int]bool{3: true}
	whole.marks[3] = map[int]bool{2: true}
	if !whole.marked(2, 4, tail{start: 4}) {
		t.Fatal("는 and then 게: two runs of markers one after the other")
	}

	myeo := []piece{{text: "며", names: "맺음:며", verdict: 1, tail: rule.TailPred}}
	link := scanOf("곰으며")
	if link.marked(1, 3, tail{start: 2, pieces: myeo}) {
		t.Fatal("으 before 며 when nothing reads as markers")
	}
	link.marks[3] = map[int]bool{1: true}
	if !link.marked(1, 3, tail{start: 2, pieces: myeo}) {
		t.Fatal("으 before 며: from the link vowel on the letters read as markers (으며)")
	}
	link.marks[3] = map[int]bool{2: true}
	if !link.marked(1, 3, tail{start: 2, pieces: myeo}) || link.marked(1, 3, tail{start: 2, only: "으", pieces: myeo}) {
		t.Fatal("으 before 며: after the link vowel the letters read as markers (며), but for a tail that names its stem")
	}
	link.marks[3] = map[int]bool{}
	link.marks[2] = map[int]bool{1: true}
	if !link.marked(1, 3, tail{start: 2, pieces: myeo}) {
		t.Fatal("으 read as a marker by itself before 며")
	}
	if link.marked(1, 3, tail{start: 2, base: '으', pieces: myeo}) {
		t.Fatal("a stem a piece is fused into, when nothing up to the tail's end reads as markers")
	}

	da := []piece{{text: "다", names: "맺음:다", verdict: 1, tail: rule.TailPred}}
	eu := scanOf("곰으셨다")
	if eu.marked(1, 4, tail{start: 3, pieces: da}) {
		t.Fatal("으셨 before 다 when nothing reads 셨다 as markers")
	}
	eu.marks[3] = map[int]bool{2: true}
	if eu.marked(1, 4, tail{start: 3, pieces: da}) {
		t.Fatal("markers that end before the tail's end do not count for the link vowel")
	}
	eu.marks[3], eu.marks[4] = nil, map[int]bool{2: true}
	if !eu.marked(1, 4, tail{start: 3, pieces: da}) || eu.marked(2, 4, tail{start: 3, pieces: da}) {
		t.Fatal("으셨 before 다: the link vowel and then a run of markers; 셨 does not start with the link vowel")
	}
	eu.marks[4] = map[int]bool{1: true}
	if !eu.marked(1, 4, tail{start: 3, pieces: da}) || !eu.marked(1, 4, tail{start: 3, base: '시', pieces: da}) {
		t.Fatal("a run of markers from the link vowel on, with a piece fused into 으시 too")
	}

	geol := []piece{{text: "(ㄹ)걸", names: "맺음:ㄹ걸", verdict: 1, tail: rule.TailPred}}
	eul := scanOf("곰을걸")
	if eul.marked(1, 3, tail{start: 2, base: '으', pieces: geol}) {
		t.Fatal("으 out of 을 before (ㄹ)걸 when nothing reads as markers")
	}
	eul.marks[3] = map[int]bool{1: true}
	if !eul.marked(1, 3, tail{start: 2, base: '으', pieces: geol}) {
		t.Fatal("으 out of 을 is the link vowel: 을걸 reads as markers from it")
	}
	two := scanOf("곰가을걸")
	two.marks[4] = map[int]bool{1: true, 2: true}
	if two.marked(1, 4, tail{start: 3, base: '으', pieces: geol}) {
		t.Fatal("a fused stem of two syllables is judged by its first letter as written (가)")
	}

	yard := scanOf("마당에서")
	if yard.marked(2, 4, tail{start: 3, pieces: seo}) {
		t.Fatal("before the search no place reads as markers")
	}
	yard.search()
	if !yard.marks[3][2] || !yard.marked(2, 4, tail{start: 3, pieces: seo}) || yard.marked(1, 4, tail{start: 3, pieces: seo}) {
		t.Fatal("after the search 에 reads as a particle where it ends, 당에 does not")
	}
}
