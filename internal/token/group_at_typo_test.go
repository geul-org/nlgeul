//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestGroupAtTypo — a closing written wrong is a piece of its own, apart from the right closing of the same letters (거에요: the letters 에요 are the right closing 에요 and, written wrong, the closing 예요); its text is the right closing, it keeps the letters it is written with, and a form under a final consonant (할께) is a piece with the right closing in brackets
package token

import "testing"

func TestGroupAtTypo(t *testing.T) {
	sc := typoScan("거에요", layerCost)
	got, from, _ := sc.groupAt(classClosing, 3, -1, false, false, false)
	var right, wrong piece
	for key, pc := range got {
		if from[key].wrote == "" && pc.text == "에요" {
			right = pc
		}
		if from[key].wrote != "" {
			wrong = pc
		}
	}
	if right.names != "맺음:에요" || right.wrote != "" {
		t.Fatal("the right closing of the letters", got)
	}
	if wrong.names != "맺음:예요" || wrong.text != "예요" || wrong.wrote != "에요" || wrong.verdict <= 0 {
		t.Fatal("the closing written wrong", got)
	}
	sc = typoScan("할께", layerCost)
	got, _, starts := sc.groupAt(classClosing, 2, -1, false, false, false)
	found := false
	for key, pc := range got {
		if pc.wrote == "ㄹ께" {
			found = found || (pc.text == "(ㄹ)게" && pc.names == "맺음:ㄹ게" && starts[key] == 1)
		}
	}
	if !found {
		t.Fatal(got)
	}
}
