[![Unit Tests Status](https://github.com/Woody1193/miletos-test/actions/workflows/test.yml/badge.svg)](https://github.com/Woody1193/miletos-test/actions)

# miletos-test
このリポジトリには、Miletosに提示されたケーススタディの解決策として役立つコードが含まれています。将来の面接のための候補者に情報提供するために使用しないでくださいが、Goの良いコーディングプラクティスの例として役立つかもしれません。このリポジトリは、さまざまな請求書作成ツールを含む架空のCLIのコードを含むことを意図しています。

## 言語のこと
[![English](https://img.shields.io/badge/lang-en-red.svg)](https://github.com/Woody1193/miletos-test/blob/main/README.md)
[![日本語](https://img.shields.io/badge/lang-jp-green.svg)](https://github.com/Woody1193/miletos-test/blob/main/README.jp.md)

## checkコマンド
checkコマンドは、請求書ファイルと売掛金ファイル間の取引が複数の基準に従って一致するかどうかを検証します。現在、以下のチェックがサポートされています：

- 特定の売掛金に対して請求書が存在すること
- 請求額が受領額と一致すること
- 受領日が請求書の支払期日より前であること
- 受領日が未来の1か月以内であること
- 支払われていない請求書が存在する場合、請求書の支払期日がまだ来ていないこと

### コマンド構文
checkコマンドの構文は以下の通りです：

``` bash
invtools check --invoice <invoice_file> --receivables <receivables_file> --output <output_file> --error <error_file>
```

### コマンド入力
checkコマンドには、2つの入力ファイルと2つの出力ファイルが必要です。これらが明示的に指定されていない場合、請求書ファイル、売掛金ファイル、出力ファイル、エラーファイルは、それぞれ`invoice.csv`、`receivables.csv`、``output.csv`、`errors.csv`として生成されます。

`invoice_file`パラメータは、請求書ファイルのファイルパスを指定します。このファイルは、以下の列を含むCSVファイルである必要があります。

- **ID (string):** 請求書のユニークな識別子。
- **Amount (integer):** 日本円での請求額。
- **Due Date (date):** 請求書の支払期限日。 YYYY-MM-DDThh:mm:ssZ形式で指定します。

`receivables_file`パラメータは、受取可能金額ファイルのファイルパスを指定します。このファイルは、以下の列を含むCSVファイルである必要があります。

- **ID (string):** 受取可能金額のユニークな識別子。
- **Amount (integer):** 日本円で受け取られた金額。
- **Date (date):** 金額が受け取られた日付。 YYYY-MM-DDThh:mm:ssZ形式で指定します。

### コマンドの出力
check コマンドは、エラーがない場合は出力を生成しません。 `output_file` には、ルールチェックに失敗したすべてのインスタンスが含まれます。

`output_file`には、以下の情報が含まれます：

- **ID (string)：** エラーが発生した請求書のID
- **Invoices File Line (integer):** このエラーに対応する請求書ファイル内の行番号
- **Receivables File Line (integer):** このエラーに対応する売掛金ファイル内の行番号
- **Description (string):** 発生した問題の説明。複数ある場合があります。

`error_file`は、このコマンドによって生成され、処理できなかったデータとその理由に関する情報が含まれます。このファイルにデータが書き込まれる場合、コマンドによるチェックは行われないため、それは不可能です。次の条件のいずれかが満たされる場合、データがこのファイルに書き込まれます：

- 請求書IDが欠落している場合
- 金額が整数値ではない場合
- 支払期日または日付が指定された形式に従って解析できない場合
- 同じ請求書IDに複数の項目がある場合。この場合、請求書IDの最初の出現が受け入れられ、その後の出現はすべて拒否されます。

`error_file`には、次の情報が含まれます。

- **File (string):** 不正なデータが含まれたファイル名
- **Line (integer):** エラーを含む行数
- **Description (string):** エラーの説明

###エラー条件
チェックコマンドは、次の状況下でエラーを生成する可能性があります：

- `invoice_file`または`receivables_file`パラメータが指定されていないか、ファイルが存在しない場合、エラーメッセージが生成されます。
- CSVファイルが正しいフォーマットでない場合、エラーメッセージが生成されます。
- `output_file`または`error_file`のどちらかが書き込めない場合、エラーメッセージが生成されます。