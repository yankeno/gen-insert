# gen-insert

## 概要

CSVまたはTSVファイルを入力として与えることで、SQLのINSERT文を生成する。

```bash
$ gen-insert users.csv
INSERT statement 'users.sql' was generated successfully!

$ ls -l
-rw-r--r--   1 yankeno  group    2300 6 18  2024 insert_users.sql
-rw-r--r--   1 yankeno  group    1200 6 18  2024 users.csv
```

## オプション

| オプション | 説明                                                     |
|:------|:-------------------------------------------------------|
| -o    | 出力先のディレクトリ／ファイル名を指定する<br />* 指定しない場合は`./{入力ファイル名}.sql` |
| -t    | テーブル名を指定する<br />* 指定しない場合は入力ファイル名                      |
