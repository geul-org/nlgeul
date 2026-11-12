//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestLoadWordsTypo — a line with a fifth column goes to the misspelt forms and is in neither the entries nor the total of counts; the lines of four columns are read as before; lines with the same two forms join their kinds, those with other right forms keep the order of the lines; a line of three or six columns is refused and says which
package token

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadWordsTypo(t *testing.T) {
	plain, _, plainTotal, err := loadWords("testdata/words.tsv")
	if err != nil {
		t.Fatal(err)
	}
	lex, typos, total, err := loadWords("testdata/words-typo.tsv")
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := lex["맞기"]; ok || len(lex) != len(plain)+3+8 || total != plainTotal+34000+80000+6400+22100 {
		t.Fatal("a misspelt line is no word", len(lex), len(plain), total, plainTotal)
	}
	if got := typos["맞기"]; len(got) != 1 || got[0] != (typo{"맡기", kindVerb}) || len(typos) != 4 {
		t.Fatal(typos)
	}
	if got := typos["역활"]; len(got) != 1 || got[0] != (typo{"역할", kindNominal}) {
		t.Fatal(typos)
	}
	if lex["맡기"].count != 6400 || lex["맡기"].kinds != kindVerb {
		t.Fatal(lex["맡기"])
	}
	dir := t.TempDir()
	ok := filepath.Join(dir, "ok.tsv")
	os.WriteFile(ok, []byte("맞기\t동사\t규칙\t0\t맡기\n맞기\t명사\t\t0\t맡기\n맞기\t명사\t\t0\t맞이\n맡기\t동사\t규칙\t5\t\n"), 0o600)
	_, typos, total, err = loadWords(ok)
	if err != nil || total != 5 || len(typos["맞기"]) != 2 || typos["맞기"][0] != (typo{"맡기", kindVerb | kindNominal}) || typos["맞기"][1] != (typo{"맞이", kindNominal}) {
		t.Fatal(typos, total, err)
	}
	for n, line := range []string{"맞기\t동사\t규칙\t0\t맡기\t여섯\n", "맞기\t동사\t규칙\n"} {
		bad := filepath.Join(dir, "bad.tsv")
		os.WriteFile(bad, []byte("# x\n맡기\t동사\t규칙\t5\t\n"+line), 0o600)
		if _, _, _, err := loadWords(bad); err == nil || !strings.Contains(err.Error(), "3번째 줄") || !strings.Contains(err.Error(), "칸 넷이나 다섯") {
			t.Fatal(n, err)
		}
	}
}
