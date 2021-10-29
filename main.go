package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	flag.Parse()

	// masterと名前のつくブランチを対象にしたマージ操作は禁止
	if flag.Arg(0) == "merge" {
		if -1 != strings.Index(flag.Arg(1), "master") {
			fmt.Print("\n\n[ERROR] DON'T MERGE INTO THE MASTER BRANCH FOR ANY REASON!\n\n\n")
			os.Exit(0)
		}
	}

	// masterとdevelopmentブランチに直接commitするのは禁止
	if flag.Arg(0) == "commit" {
		// 現在のブランチを取得
		result, _ := exec.Command("git", []string{"rev-parse", "--abbrev-ref"}...).Output()
		branchName := string(result)
		if branchName == "master" || branchName == "development" {
			fmt.Print("\n\n[ERROR] DON'T MERGE INTO THE MASTER BRANCH FOR ANY REASON!\n\n\n")
			os.Exit(0)
		}
	}

	// gitコマンドを生成
	cmd := exec.Command("git", flag.Args()...)

	// 色情報を保持するために標準出力と標準エラー出力を置換
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// gitコマンド実行
	cmd.Run()
}
