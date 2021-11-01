package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args[1:]

	// masterと名前のつくブランチを対象にしたマージ操作は禁止
	if args[0] == "merge" {
		if -1 != strings.Index(args[1], "master") {
			fmt.Print("\n\n[ERROR] DON'T MERGE INTO THE MASTER BRANCH FOR ANY REASON!\n\n\n")
			os.Exit(0)
		}
	}

	// masterとdevelopmentブランチに直接commitするのは禁止
	if args[0] == "commit" {
		// 現在のブランチを取得
		result, _ := exec.Command("git", []string{"rev-parse", "--abbrev-ref", "HEAD"}...).Output()
		branchName := strings.TrimRight(string(result), "\n")
		if branchName == "master" || branchName == "development" {
			fmt.Print("\n\n[ERROR] DON'T MERGE INTO THE MASTER OR DEVELOPMENT BRANCH FOR ANY REASON!\n\n\n")
			os.Exit(0)
		}
	}

	// gitコマンドを生成
	cmd := exec.Command("git", args...)

	// 色情報を保持するために標準出力と標準エラー出力を置換
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// gitコマンド実行
	cmd.Run()
}
