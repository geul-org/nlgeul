//ff:func feature=token type=reader control=sequence
//ff:what TestHaTails — 치 after 원 is read as 맺음:지 with the letters as they stand, asks for a stem whose 하 is dropped (irrHa, headHa) and is no slot; 컨대 likewise as 맺음:건대; a particle right after turns it down, but not before 요 (more); nothing at the start of the unit, at the head, without a word table, where the letters are not a written form, or where only the last letter matches
package token

import "testing"

func TestHaTails(t *testing.T) {
	got := wordsScan("원치", layerCost).haTails(2, false)
	if len(got) != 1 || got[0].start != 1 || got[0].irr != irrHa || got[0].head != headHa || len(got[0].pieces) != 1 || got[0].pieces[0].text != "치" || got[0].pieces[0].names != "맺음:지" || got[0].close == nil || got[0].close.form != "치" {
		t.Fatal(got)
	}
	if got = wordsScan("단언컨대", layerCost).haTails(4, false); len(got) != 1 || got[0].start != 2 || got[0].pieces[0].text != "컨대" || got[0].pieces[0].names != "맺음:건대" {
		t.Fatal("컨대", got)
	}
	if got = wordsScan("원치는", layerCost).haTails(2, false); len(got) != 0 {
		t.Fatal("a particle right after turns it down", got)
	}
	if got = wordsScan("원치는", layerCost).haTails(2, true); len(got) != 1 {
		t.Fatal("before 요 what follows is not asked", got)
	}
	if got = wordsScan("치", layerCost).haTails(1, false); len(got) != 0 {
		t.Fatal("no letter before it", got)
	}
	if got = wordsScan("원치", layerCost).haTails(0, false); got != nil {
		t.Fatal("at the head", got)
	}
	if got = scanOf("원치").haTails(2, false); len(got) != 0 {
		t.Fatal("no word table", got)
	}
	if got = wordsScan("원다", layerCost).haTails(2, false); len(got) != 0 {
		t.Fatal("다 is no written form of clause 40", got)
	}
	if got = wordsScan("원도록", layerCost).haTails(3, false); len(got) != 0 {
		t.Fatal("도록 is not 토록", got)
	}
}
