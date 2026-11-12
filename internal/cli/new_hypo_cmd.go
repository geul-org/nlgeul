//ff:func feature=cli type=command control=sequence
//ff:what newHypoCmd — nlgeul hypo: read a sentence as nlgeul token does, pick one reading for every bundle in every way, and print the hypotheses sorted by the sorting formula
package cli

import (
	"os"

	"github.com/geul-org/nlgeul/internal/token"
	"github.com/spf13/cobra"
)

// newHypoCmd is `nlgeul hypo "<문장>"`. The word table comes from
// --words, else from the environment variable NLGEUL_WORDS, as for
// nlgeul token.
func newHypoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   `hypo "<문장>"`,
		Short: "토큰 층의 읽기를 묶음마다 하나씩 골라 이은 가설을 정렬해 나열한다(실험)",
		Long: `문장 하나를 nlgeul token과 같이 읽고, 묶음(nlgeul token이 한 줄이나 경합 하나로
찍는 구간)마다 읽기를 하나씩 골라 이은 조합(가설)을 모두 나열한다. 읽기는
nlgeul token이 찍는 읽기 그대로다(재워 둔 읽기는 쓰지 않는다).

첫 줄은 "가설 <수>"이고, 그 뒤 가설마다 한 줄이다: 순위, 가설의 성립과 선호,
묶음별 읽기(" | "로 가른다).

정렬: 가설의 성립은 고른 읽기들의 성립 가운데 최솟값, 선호는 고른 읽기들의 선호의
합이다. 성립이 큰 가설이 먼저고, 성립이 같으면 선호가 큰 가설이 먼저다. 그것도
같으면 앞 묶음부터 nlgeul token이 낸 읽기 순서를 따른다.

가설이 1000000개를 넘는 문장은 가설 수만 찍고 나열하지 않으며 오류로 끝난다.

낱말 표는 nlgeul token과 같이 --words 또는 환경 변수 NLGEUL_WORDS로 준다.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			words, _ := cmd.Flags().GetString("words")
			if words == "" {
				words = os.Getenv("NLGEUL_WORDS")
			}
			r, err := token.New(token.Options{Words: words})
			if err != nil {
				return err
			}
			return runHypo(cmd.OutOrStdout(), r, args[0])
		},
	}
	cmd.Flags().String("words", "", "낱말 표 파일(없으면 환경 변수 NLGEUL_WORDS, 그것도 없으면 낱말 표 없이)")
	return cmd
}
