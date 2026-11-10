# cli

cobra 루트 명령 `kopas`(`newRootCmd`), 하위 명령 `parse`(`newParseCmd`, `runParse`), 실행(`Execute`). `parse`는 인자를 정확히 하나 받고 `parse.Lines`가 돌려준 줄(`조각+조각:문장성분`)을 표준 출력에 찍는다. 인자가 없거나 둘 이상이면 오류(종료 코드 1)다.
