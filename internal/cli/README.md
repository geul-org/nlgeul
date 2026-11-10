# cli

cobra 루트 명령 `kopas`(`newRootCmd`), 하위 명령 `parse`(`newParseCmd`, `runParse`), 실행(`Execute`). `parse`는 인자를 정확히 하나 받고 `parse.Lines`가 돌려준 줄(`조각+조각:문장성분`)을 표준 출력에 찍는다. 인자가 없거나 둘 이상이면 오류(종료 코드 1)다.

하위 명령 `token`(`newTokenCmd`, `runToken`)은 인자 하나를 받아 `Reader.Lines`의 줄을 찍는다(`Reader.Bundles`의 구조를 찍은 것이다). `--trace`는 판정 과정을, `--cuts`는 측정용 어절 끝 자리를, `--asleep`는 재워 둔 읽기도 찍는다. `--words`(또는 환경 변수 `KOPAS_WORDS`)로 낱말 표를, `--layers`로 그 표를 몇 겹까지 쓸지를 주고, `--stdin`이면 인자 대신 표준 입력의 줄마다 문장을 읽는다(`tokenInput`).
