# move_file

指定フォルダ構造の中から特定のサブフォルダを見つけ、別の場所へコピーするCLIツールです。

## 使い方

```
copy_tool -src <コピー元ルート> -target <探すフォルダ名> -dest <コピー先>
```

### フラグ

| フラグ | 説明 |
|--------|------|
| `-src` | コピー元のルートフォルダ |
| `-target` | 各サブフォルダ内で探すフォルダ名 |
| `-dest` | コピー先のルートフォルダ |

### 例

```
copy_tool -src /data/projects -target output -dest /backup
```

`/data/projects` 直下の各サブフォルダの中に `output` フォルダがあれば、その中身を `/backup/<サブフォルダ名>/` へコピーします。

```
/data/projects/
  projectA/
    output/        ← コピー元
      result.csv
  projectB/
    output/        ← コピー元
      result.csv
  projectC/        ← output がないのでスキップ

/backup/
  projectA/
    result.csv     ← コピー先
  projectB/
    result.csv     ← コピー先
```

## 動作仕様

- コピー先に同名ファイルが既に存在する場合はスキップ（上書きしない）
- `target` フォルダが存在しないサブフォルダは無視
- コピー先フォルダが存在しない場合は自動作成

## ビルド

```
go build -o copy_tool .
```

## 動作環境

- Go 1.18 以上
