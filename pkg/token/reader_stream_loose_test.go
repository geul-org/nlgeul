//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamLoose — with the word table of testdata: 내게 and 네게도 read the pronoun and 에게's 게, 어디다 and 여기다가 the place word and 에다 or 에다가, 가야겠다, 먹어야겠네 and 갔어야겠지 (하) between 어야 and 겠, 되도, 되요 and 됬다 an 어 left out after 되 next to 되 and the particle 도, and 줄은 the modifying 은 after 줄 as the second candidate, behind the table's 줄+은 and with a lower verdict; without it 내게 and 어디다 read the fused form's name; 학교다 has no 에다, a 겠다 written apart no (하), and without a word table 줄은 no modifying 은
package token

import (
	"fmt"
	"strings"
	"testing"
)

func TestReaderStreamLoose(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	bare, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	present := []struct {
		r          *Reader
		text, want string
	}{
		{r, "내게", "\n내[열쇠:체언]+게[부사격:에게]\n"},
		{r, "네게도", "\n네[열쇠:체언]+게[부사격:에게]+도[보조사:도]\n"},
		{bare, "내게", "\n내[줄임:내게]+게[부사격:에게]\n"},
		{r, "어디다", "\n어디[열쇠:체언]+다[부사격:에다]\n"},
		{r, "여기다가", "\n여기[열쇠:체언]+다가[부사격:에다가]\n"},
		{bare, "어디다", "\n어디[줄임:어디다]+다[부사격:에다]\n"},
		{r, "가야겠다", "\n가[열쇠:동사]+(아)야[맺음:어야(합침)]+(하)[보조:하]+겠[선어말:겠]+다[맺음:다]\n"},
		{r, "먹어야겠네", "\n먹[열쇠:동사]+어야[맺음:어야]+(하)[보조:하]+겠[선어말:겠]+네[맺음:네]\n"},
		{r, "갔어야겠지", "\n가[열쇠:동사]+(았)[선어말:ㅆ]+어야[맺음:어야]+(하)[보조:하]+겠[선어말:겠]+지[맺음:지]\n"},
		{r, "가야겠어요", "\n가[열쇠:동사]+(아)야[맺음:어야(합침)]+(하)[보조:하]+겠[선어말:겠]+어[맺음:어]+요[보조사:요]\n"},
		{r, "되도", "\n되[열쇠:동사]+(어)도[맺음:어도(합침)]\n"},
		{r, "되도", "\n되+도[보조사:도]\n"},
		{bare, "되도", "\n되+(어)도[맺음:어도(합침)]\n"},
		{r, "되요", "\n되[열쇠:동사]+(어)[맺음:어(합침)]+요[보조사:요]\n"},
		{r, "됬다", "\n되[열쇠:동사]+(었)[선어말:ㅆ]+다[맺음:다]\n"},
	}
	for _, c := range present {
		if got := streamText(c.r, c.text); !strings.Contains(got, c.want) {
			t.Errorf("missing %q in\n%s", c.want, got)
		}
	}
	absent := []struct {
		r          *Reader
		text, want string
	}{
		{r, "학교다", "[부사격:에다]"},
		{bare, "학교다", "[부사격:에다]"},
		{r, "가야 겠다", "[보조:하]"},
		{bare, "줄은", "[꾸밈:은]"},
		{bare, "줄을", "[꾸밈:을]"},
	}
	for _, c := range absent {
		if got := streamText(c.r, c.text); strings.Contains(got, c.want) {
			t.Errorf("%q in\n%s", c.want, got)
		}
	}
	lines := scoredLines(r, "줄은")
	if len(lines) < 2 || !strings.HasPrefix(lines[0], "줄[의존명사:줄]+은[보조사:은]\t") || !strings.HasPrefix(lines[1], "줄[열쇠:동사]+은[꾸밈:은]\t") {
		t.Fatal(lines)
	}
	_, first, _ := strings.Cut(lines[0], "\t성립 ")
	_, second, _ := strings.Cut(lines[1], "\t성립 ")
	var table, loose float64
	if _, err = fmt.Sscan(first, &table); err != nil {
		t.Fatal(err)
	}
	if _, err = fmt.Sscan(second, &loose); err != nil || loose <= 0 || loose >= table {
		t.Fatal("the modifying 은 after 줄 stands, rebutted below the table's candidate", table, loose, err)
	}
}
