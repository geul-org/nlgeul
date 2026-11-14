# cli

cobra 루트 명령 `nlgeul`(`newRootCmd`), 하위 명령 `token`, 실행(`Execute`). `parse`와 `combo` 하위 명령은 없다(`internal/parse`와 `internal/combo`는 내부 함수로 남아 있다).

하위 명령 `token`(`newTokenCmd`, `runToken`, 줄 고르기 `tokenLines`)은 인자 하나를 받아 기본으로 `Reader.BoxLines`의 상자를 찍는다(스트림의 갈래를 어절마다 상자로, 어절경합은 오른쪽 괄호로 그린 것이다). `--plain`이면 옛 줄 형식 `Reader.LinesHead`의 줄(`Reader.Contests`의 구조를 찍은 것이다)을 찍는다. `--head`는 상자 하나의 후보 줄과 어절경합 하나의 갈래를(`--plain`이면 경합 하나의 후보 줄을) 줄인다. `--trace`는 판정 과정을(상자 출력에서는 그 단위의 상자들 앞에), `--cuts`는 측정용 어절 끝 자리를, `--marks`는 거기에 경계마다 띄어쓰기 조항 표시를 더해 찍는다. `--spacing`(기본 `all`, `none`이나 쉼표로 이은 조항 이름, `spacingOption`)은 켤 띄어쓰기 조항을 `token.Options.Spacing`으로 준다. 글자로 읽는 스크립트는 `--plain`을 붙인다. `--words`(또는 환경 변수 `NLGEUL_WORDS`)로 낱말 표를, `--layers`로 그 표를 몇 겹까지 쓸지를 주고, `--stdin`이면 인자 대신 표준 입력의 줄마다 문장을 읽는다(`tokenInput`).
