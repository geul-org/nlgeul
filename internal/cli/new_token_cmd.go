//ff:func feature=cli type=command control=sequence
//ff:what newTokenCmd — nlgeul token: read a token (syllable) stream into words with a graph for every marker candidate, with the spacing rules (article 2 and 41 to 50 of the spelling rules, --spacing) that show boundaries inside a unit, and print the words as boxes (the ways of a stretch in brackets at their right), or with --format plain as lines where the candidates do not agree, or with --format json as one line of JSON for a sentence
package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/geul-org/nlgeul/pkg/token"
	"github.com/spf13/cobra"
)

// newTokenCmd is `nlgeul token "<문장>"`. The word table comes from
// --words, else from the environment variable NLGEUL_WORDS.
func newTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   `token "<문장>"`,
		Short: "토큰(음절) 스트림을 표지별 판정 그래프로 읽어 어절로 끊어 올린다(실험)",
		Long:  tokenLong,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := cmd.Flags()
			trace, _ := f.GetBool("trace")
			cuts, _ := f.GetBool("cuts")
			stdin, _ := f.GetBool("stdin")
			words, _ := f.GetString("words")
			layers, _ := f.GetInt("layers")
			head, _ := f.GetInt("head")
			marks, _ := f.GetBool("marks")
			spacing, _ := f.GetString("spacing")
			format, _ := f.GetString("format")
			if format != "box" && format != "plain" && format != "json" {
				return fmt.Errorf("--format은 box, plain, json 중 하나다: %q", format)
			}
			if format == "json" && (cuts || marks) {
				return errors.New("--format json은 --cuts, --marks와 함께 줄 수 없다")
			}
			plain := format == "plain"
			if head < 0 {
				return fmt.Errorf("--head는 0 이상이다(0은 전부 찍는다): %d", head)
			}
			if words == "" {
				words = os.Getenv("NLGEUL_WORDS")
			}
			sentences, err := tokenInput(args, stdin, cmd.InOrStdin())
			if err != nil {
				return err
			}
			r, err := token.New(token.Options{Words: words, Layers: layers, Spacing: spacingOption(spacing)})
			if err != nil {
				return err
			}
			if format == "json" {
				return writeTokenJSON(cmd.OutOrStdout(), r, sentences)
			}
			runToken(cmd.OutOrStdout(), r, sentences, trace, cuts, marks, plain, head)
			return nil
		},
	}
	cmd.Flags().Bool("trace", false, "단위마다 걸린 토큰 열, 그래프, verdict, 기각 이유를 그 단위의 상자들 앞에 함께 찍는다(--format plain이면 그 단위의 경합 앞에)")
	cmd.Flags().String("format", "box", "출력 형식: box(기본, 상자), plain(옛 줄 형식: 경합 한 줄과 후보 줄들, 기계가 읽기 쉽다. 옛 탐색 그대로라 긴 붙여 쓴 덩어리는 box로), json(문장마다 JSON 한 줄, 스크립트용. 후보를 전부 담고 --head를 걸지 않으며 --trace는 담지 않는다. --cuts, --marks와는 함께 줄 수 없다)")
	cmd.Flags().Bool("cuts", false, "측정용: 단위마다 살아남은 후보들의 어절 끝 자리(토큰 수)만 찍는다")
	cmd.Flags().Bool("marks", false, "측정용: 단위마다 후보들의 어절 끝 자리와, 그 경계가 어느 띄어쓰기 조항의 자리인지(조항 번호를 /로 이음, 없으면 -)를 찍는다")
	cmd.Flags().String("spacing", "all", "켤 띄어쓰기 조항(「한글 맞춤법」 제2항, 제41~50항의 번호): all(전부), none(끔), 또는 쉼표로 이은 조항 이름(2,41,42,43,43다만,44,45,46,47,47다만,48,48다만,49,50)")
	cmd.Flags().Bool("stdin", false, "표준 입력의 줄마다 문장 하나로 읽고, 문장 사이에 빈 줄을 찍는다")
	cmd.Flags().String("words", "", "낱말 표 파일(없으면 환경 변수 NLGEUL_WORDS, 그것도 없으면 낱말 표 없이)")
	cmd.Flags().Int("layers", 0, "측정용: 낱말 표를 몇 겹까지 쓸지(1 여부, 2 갈래 맞춤, 3 구조 값; 0은 전부)")
	cmd.Flags().Int("head", 20, "상자 하나의 후보 줄과 어절경합 하나의 갈래를 앞의 n개까지만 찍고 나머지는 (k개 생략됨)으로 센다(--format plain이면 경합 하나의 후보 줄. 찍는 줄 수일 뿐 후보는 전부 만든다; 0은 전부 찍는다)")
	return cmd
}
