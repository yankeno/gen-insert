# gen-insert

## 概要

CSVまたはTSVファイルを入力として与えることで、SQLのINSERT文を生成するコマンドです。

```bash
$ gen-insert ./users.csv
INSERT statement 'insert_users.sql' was generated successfully!

$ ls -l
-rw-r--r--   1 yankeno  group    2300 6 18  2024 insert_users.sql
-rw-r--r--   1 yankeno  group    1200 6 18  2024 users.csv
```

## オプション

| オプション | 説明          |
|:------|:------------|
| -o    | 出力先のパスを指定する |
| -n    | ファイル名を指定する  |
