//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamHaoEndings — end to end with the word table of testdata, the 하오체 and 하십시오체 endings added to the end of closings.tsv: 습디다 and 습디까 after a consonant and ㅂ디다, ㅂ디까 under a vowel; 구려; 소이까 after a consonant; 오니까 and 오리이까 (으오니까, 으오리이까 after a consonant); 렷다, 리다 and 리까 (으렷다, 으리다, 으리까 after a consonant); 노니 after a verb's stem and not after an adjective's
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamHaoEndings(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"작습디다":   "\n작[열쇠:형용사]+습디다[맺음:습디다]\n",
		"갑디다":    "\n가[열쇠:동사]+(ㅂ)디다[맺음:ㅂ디다]\n",
		"먹었습디까":  "\n먹[열쇠:동사]+었[선어말:었]+습디까[맺음:습디까]\n",
		"갑디까":    "\n가[열쇠:동사]+(ㅂ)디까[맺음:ㅂ디까]\n",
		"가구려":    "\n가[열쇠:동사]+구려[맺음:구려]\n",
		"먹소이까":   "\n먹[열쇠:동사]+소이까[맺음:소이까]\n",
		"하오니까":   "\n하[열쇠:동사]+오니까[맺음:오니까]\n",
		"먹으오니까":  "\n먹[열쇠:동사]+으오니까[맺음:으오니까]\n",
		"가오리이까":  "\n가[열쇠:동사]+오리이까[맺음:오리이까]\n",
		"먹으오리이까": "\n먹[열쇠:동사]+으오리이까[맺음:으오리이까]\n",
		"가렷다":    "\n가[열쇠:동사]+렷다[맺음:렷다]\n",
		"먹으렷다":   "\n먹[열쇠:동사]+으렷다[맺음:으렷다]\n",
		"가리다":    "\n가[열쇠:동사]+리다[맺음:리다]\n",
		"먹으리다":   "\n먹[열쇠:동사]+으리다[맺음:으리다]\n",
		"가리까":    "\n가[열쇠:동사]+리까[맺음:리까]\n",
		"먹으리까":   "\n먹[열쇠:동사]+으리까[맺음:으리까]\n",
		"가노니":    "\n가[열쇠:동사]+노니[맺음:노니]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	for in, not := range map[string]string{
		"작노니":  "작[열쇠:형용사]+노니",
		"가습디다": "가[열쇠:동사]+습디다",
		"먹ㅂ디다": "먹[열쇠:동사]+(ㅂ)디다",
	} {
		if got := streamText(r, in); strings.Contains(got, not) {
			t.Errorf("%s: %q read\n%s", in, not, got)
		}
	}
}
