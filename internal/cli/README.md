# cli

cobra 루트 명령 `nlgeul`(`newRootCmd`), 하위 명령 `token`, 실행(`Execute`). `parse`와 `hypo` 하위 명령은 없다(`internal/parse`와 `internal/hypo`는 내부 함수로 남아 있다).

하위 명령 `token`(`newTokenCmd`, `runToken`)은 인자 하나를 받아 `Reader.Lines`의 줄을 찍는다(`Reader.Bundles`의 구조를 찍은 것이다). `--trace`는 판정 과정을, `--cuts`는 측정용 어절 끝 자리를 찍는다. `--words`(또는 환경 변수 `NLGEUL_WORDS`)로 낱말 표를, `--layers`로 그 표를 몇 겹까지 쓸지를 주고, `--stdin`이면 인자 대신 표준 입력의 줄마다 문장을 읽는다(`tokenInput`).
