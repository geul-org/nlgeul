# cli

cobra 루트 명령 `nlgeul`(`newRootCmd`), 하위 명령 `parse`, 실행(`Execute`). `combo` 하위 명령은 없다(`internal/combo`는 함수로만 남아 있다).

하위 명령 `parse`(`newParseCmd`, `runParse`, 줄 고르기 `parseLines`)은 인자 하나를 받아 기본으로 `Reader.BoxLines`의 상자를 찍는다(스트림의 갈래를 어절마다 상자로, 어절경합은 오른쪽 괄호로 그린 것이다). `--format json`이면 문장마다 `Reader.JSONSentence`를 JSON 한 줄로 찍는다(`writeParseJSON`. 후보를 전부 담고 `--head`와 `--trace`를 쓰지 않으며 `--cuts`·`--marks`와 함께 주면 오류. `--format`은 `box`(기본)와 `json`이고, 옛 `plain`은 없앴다(주면 오류). `--head`는 상자 하나의 후보 줄과 어절경합 하나의 갈래를 줄인다. `--trace`는 판정 과정을(그 단위의 상자들 앞에), `--cuts`는 측정용 어절 끝 자리를, `--marks`는 거기에 경계마다 띄어쓰기 조항 표시를 더해 찍는다. `--spacing`(기본 `all`, `none`이나 쉼표로 이은 조항 이름, `spacingOption`)은 켤 띄어쓰기 조항을 `parse.Options.Spacing`으로 준다. 스크립트는 `--format json`을 쓴다. `--words`(또는 환경 변수 `NLGEUL_WORDS`)로 낱말 표를, `--layers`로 그 표를 몇 겹까지 쓸지를 주고, `--stdin`이면 인자 대신 표준 입력의 줄마다 문장을 읽는다(`parseInput`).
