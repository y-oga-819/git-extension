package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args[1:]

	// masterと名前のつくブランチをマージするのは禁止
	if args[0] == "merge" {
		if -1 != strings.Index(args[1], "master") {
			fmt.Print("\n\n[ERROR] DON'T MERGE INTO THE MASTER BRANCH FOR ANY REASON!\n\n\n")
			os.Exit(0)
		}
	}

	// masterとdevelopmentブランチに直接commitやマージするのは禁止
	if args[0] == "commit" || args[0] == "merge" {
		// 現在のブランチを取得
		result, _ := exec.Command("git", []string{"rev-parse", "--abbrev-ref", "HEAD"}...).Output()
		branchName := strings.TrimRight(string(result), "\n")
		if branchName == "master" || branchName == "development" {
			fmt.Print("\n\n[ERROR] DON'T MERGE/COMMIT INTO THE MASTER OR DEVELOPMENT BRANCH FOR ANY REASON!\n\n\n")
			os.Exit(0)
		}
	}

	// どうしてもmasterやdevelopmentに何か操作したい時に備えてforceコマンドを実装しておく
	if args[0] == "force" {
		args = args[1:]
	}

	// hubコマンドを利用しているなら差し替える
	gitCommand := "git"
	if existsHubCommand() {
		gitCommand = "hub"
	}

	// gitコマンドを生成
	cmd := exec.Command(gitCommand, args...)

	// 色情報を保持するために標準出力と標準エラー出力を置換
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// gitコマンド実行
	cmd.Run()
}

func existsHubCommand() bool {
	result, _ := exec.Command("type", "hub").Output()
	hubCommandExists := strings.TrimRight(string(result), "\n")
	return hubCommandExists != "hub not found"
}
