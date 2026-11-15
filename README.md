# copy-file-tool

指定したフォルダの直下にある各サブフォルダから、決まった名前のフォルダ（例: `output`）を探し、その中のファイルを別の場所へサブフォルダ名ごとにコピーするGoのCLIツール。

## 動かし方

`go.mod` の指定はGo 1.26.1。外部の依存は無い。

```
go build -o copy_tool .
copy_tool -src <コピー元ルート> -target <探すフォルダ名> -dest <コピー先> [-overwrite]
```

`-h` で使い方を表示する。

| フラグ | 説明 |
|---|---|
| `-src` | コピー元のルートフォルダ |
| `-target` | 各サブフォルダの中で探すフォルダ名 |
| `-dest` | コピー先のルートフォルダ |
| `-overwrite` | コピー先に同じ名前のファイルがあるとき上書きする（省略時はスキップ） |

例えば次のように実行すると、`/data/projects` 直下の各サブフォルダに `output` フォルダがあれば、その中のファイルを `/backup/<サブフォルダ名>/` へコピーする。

```
copy_tool -src /data/projects -target output -dest /backup
```

```
/data/projects/
  projectA/
    output/        ← コピー元
      result.csv
  projectB/
    output/        ← コピー元
      result.csv
  projectC/        ← output が無いのでスキップ

/backup/
  projectA/
    result.csv     ← コピー先
  projectB/
    result.csv     ← コピー先
```

動作は次のとおり。

- コピー先に同じ名前のファイルがあればスキップする（`-overwrite` を付けたときは上書きする）
- `-target` のフォルダが無いサブフォルダは無視する
- コピー先のフォルダが無ければ作る
- `-target` のフォルダの直下にあるファイルだけをコピーし、その中のサブフォルダはコピーしない
- サブフォルダごとのコピー件数・スキップ件数と、最後に合計を表示する

## 使った技術

- Go（標準ライブラリのみ。`flag`、`os`、`path/filepath`、`io`）

## 学んだこと

- `flag.Usage` を差し替えて、日本語の説明と使用例を含むヘルプを表示する方法
