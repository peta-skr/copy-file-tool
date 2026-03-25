package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// CLIフラグの定義
	src := flag.String("src", "", "コピー元ルート")
	target := flag.String("target", "", "探すフォルダ")
	dest := flag.String("dest", "", "コピー先")
	flag.Parse()

	// 必須引数のチェック
	if *src == "" || *target == "" || *dest == "" {
		fmt.Fprintln(os.Stderr, "使い方: copy_tool -src <フォルダA> -target <フォルダE名> -dest <フォルダF>")
		flag.Usage()
		os.Exit(1)
	}

	// コピー元ルート直下のサブフォルダ一覧を取得
	entries, err := os.ReadDir(*src)
	if err != nil {
		fmt.Fprintln(os.Stderr, "フォルダの中身の読み込みでエラーが発生しました。")
		os.Exit(1)
	}

	totalCopied, totalSkipped := 0, 0

	for _, entry := range entries {
		// ファイルは無視してフォルダのみ処理
		if !entry.IsDir() {
			continue
		}

		// 対象フォルダのパス: src/サブフォルダ/target
		targetDir := filepath.Join(*src, entry.Name(), *target)

		// target フォルダが存在しないサブフォルダはスキップ
		_, err := os.Stat(targetDir)
		if os.IsNotExist(err) {
			continue
		}

		// コピー先: dest/サブフォルダ/
		destDir := filepath.Join(*dest, entry.Name())

		c, s, err := copyDir(targetDir, destDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "コピーエラー (%s): %v\n", entry.Name(), err)
			continue
		}

		fmt.Printf("[%s/%s] -> [%s/%s]: コピー %d件, スキップ %d件\n",
			entry.Name(), *target, filepath.Base(*dest), entry.Name(), c, s)
		totalCopied += c
		totalSkipped += s
	}

	fmt.Printf("\n完了: 合計コピー %d件, スキップ %d件\n", totalCopied, totalSkipped)
}

// copyDir はsrcDir内のファイルをdestDirへコピーする。
// 既にdestDirに同名ファイルが存在する場合はスキップする。
func copyDir(srcDir, destDir string) (copied, skipped int, err error) {
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return 0, 0, err
	}

	for _, entry := range entries {
		// サブフォルダは再帰しない（フラットにコピー）
		if entry.IsDir() {
			continue
		}

		destPath := filepath.Join(destDir, entry.Name())

		// 同名ファイルが既に存在する場合はスキップ
		_, err := os.Stat(destPath)
		if err == nil {
			skipped++
			continue
		}

		// コピー先フォルダがなければ作成
		if err := os.MkdirAll(destDir, 0755); err != nil {
			return copied, skipped, err
		}

		srcPath := filepath.Join(srcDir, entry.Name())
		if err := copyFile(srcPath, destPath); err != nil {
			return copied, skipped, fmt.Errorf("ファイルのコピー失敗 %s: %w", entry.Name(), err)
		}
		copied++
	}

	return copied, skipped, nil
}

// copyFile は1つのファイルをsrcからdestへコピーする。
func copyFile(src, dest string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
