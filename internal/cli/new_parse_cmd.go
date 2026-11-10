//ff:func feature=cli type=command control=sequence
//ff:what newParseCmd — nlgeul parse: one line "어절:문장성분" for each word of the sentence
package cli

import "github.com/spf13/cobra"

// newParseCmd is `nlgeul parse "<문장>"`.
func newParseCmd() *cobra.Command {
	return &cobra.Command{
		Use:   `parse "<문장>"`,
		Short: "문장의 어절마다 문장성분을 문법 표만으로 가려 찍는다",
		Long: `문장 하나를 띄어쓰기로 나눠 어절마다 "어절:문장성분"을 한 줄씩 찍는다.
이름은 주어, 목적어, 보어, 서술어, 관형어, 부사어, 독립어, 미정이고
풀이가 여럿이면 "|"로 잇는다. 단어와 조사, 용언의 줄기와
꼬리 조각 사이에는 "+"를 넣는다
(고양이+가:주어, 학교+에서+는, 곰+이+다:서술어, 지시+해서, 가+시+었+다+고).
사전 없이 조사·어미 표와 표에 적힌 닫힌 집합 낱말(의존명사, 보조 용언,
지시·수 관형사, 접속 부사)만으로 판단하는 실험이다.

입력 문장의 전제조건:
  1. 동사 서술은 반드시 마지막에 온다.
  2. 띄어쓰기를 올바로 한다.
  3. 조사를 반드시 올바로 붙인다.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runParse(cmd.OutOrStdout(), args[0])
		},
	}
}
