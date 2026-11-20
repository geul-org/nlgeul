//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestSaisiotParts — with the small word table of 사이시옷 (words-saisiot.tsv): 전셋값, 장맛비, 아랫목, 배춧잎 (ㄴㄴ added), 처갓집, 머릿결 and 전깃불이 (back 불이) are two nominals of the table; the first place that fits is taken, past one whose back is no word (나뭇가짓값); none for 아랫층 (ㅊ) and 짜릿함 (ㅎ), for 오얏나무 and 이것밖 (the ㅅ is the word's own, though 오야·나무 and 이거·밖 are nominals), for 귓밥 (the front has one syllable, though 귀·밥 are nominals), for 전셋값 with no word table, for a ㅅ under the last syllable, for a front or back no nominal of the table (아랫갑, 앉잣값) and for Latin letters
package parse

import "testing"

func TestSaisiotParts(t *testing.T) {
	ix := streamReader(t, "words-saisiot.tsv").ix
	cases := []struct{ stem, front, back string }{
		{"전셋값", "전세", "값"}, {"장맛비", "장마", "비"}, {"아랫목", "아래", "목"}, {"배춧잎", "배추", "잎"},
		{"처갓집", "처가", "집"}, {"머릿결", "머리", "결"}, {"전깃불이", "전기", "불이"}, {"나뭇가짓값", "나뭇가지", "값"},
	}
	for _, c := range cases {
		if front, back, ok := ix.saisiotParts(c.stem); !ok || front != c.front || back != c.back {
			t.Errorf("%s: %s %s %v", c.stem, front, back, ok)
		}
	}
	for _, stem := range []string{"아랫층", "짜릿함", "오얏나무", "이것밖", "귓밥", "아랫", "아랫갑", "앉잣값", "abc"} {
		if front, back, ok := ix.saisiotParts(stem); ok || front != "" || back != "" {
			t.Errorf("%s: %s %s", stem, front, back)
		}
	}
	for _, w := range []string{"오야", "나무", "이거", "밖", "귀", "밥"} {
		if ix.lex[w].kinds&kindNominal == 0 {
			t.Fatal(w, "is a nominal of the table, so only the rule keeps it apart")
		}
	}
	if _, _, ok := streamReader(t, "").ix.saisiotParts("전셋값"); ok {
		t.Fatal("no word table")
	}
}
