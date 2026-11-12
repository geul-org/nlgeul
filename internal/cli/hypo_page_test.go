//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestHypoPage — no option gives page 0 and hypo.PerPage; --page gives its page, of hypo.PerPage hypotheses or of --per-page in either order; a --page below 1, a --per-page below 1 and a --per-page without --page, even one of hypo.PerPage, are refused, the --page first
package cli

import (
	"strconv"
	"testing"

	"github.com/geul-org/nlgeul/internal/hypo"
)

func TestHypoPage(t *testing.T) {
	cases := []struct {
		args          []string
		page, perPage int
		err           string
	}{
		{nil, 0, hypo.PerPage, ""},
		{[]string{"--page", "2"}, 2, hypo.PerPage, ""},
		{[]string{"--page", "3", "--per-page", "7"}, 3, 7, ""},
		{[]string{"--per-page", "1", "--page", "1"}, 1, 1, ""},
		{[]string{"--per-page", "7"}, 0, 0, "--per-page는 --page와 함께 준다"},
		{[]string{"--per-page", strconv.Itoa(hypo.PerPage)}, 0, 0, "--per-page는 --page와 함께 준다"},
		{[]string{"--page", "0"}, 0, 0, "--page는 1 이상이어야 한다: 0"},
		{[]string{"--page=-3"}, 0, 0, "--page는 1 이상이어야 한다: -3"},
		{[]string{"--page", "0", "--per-page", "0"}, 0, 0, "--page는 1 이상이어야 한다: 0"},
		{[]string{"--page", "1", "--per-page", "0"}, 0, 0, "--per-page는 1 이상이어야 한다: 0"},
		{[]string{"--page", "1", "--per-page=-5"}, 0, 0, "--per-page는 1 이상이어야 한다: -5"},
	}
	for _, c := range cases {
		// The hypo command as nlgeul builds it, so that the options and
		// their defaults are its own.
		cmd, _, err := newRootCmd().Find([]string{"hypo"})
		if err != nil || cmd.Name() != "hypo" || cmd.ParseFlags(c.args) != nil {
			t.Fatalf("%q: the hypo command does not take them: %v", c.args, err)
		}
		page, perPage, err := hypoPage(cmd)
		got := ""
		if err != nil {
			got = err.Error()
		}
		if got != c.err || (err == nil && (page != c.page || perPage != c.perPage)) {
			t.Errorf("%q: page %d, %d per page, error %q; want %d, %d, %q", c.args, page, perPage, got, c.page, c.perPage, c.err)
		}
	}
}
