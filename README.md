# git-extension
- masterと名前のつくブランチをマージするのは禁止
- masterとdevelopmentブランチに直接commitやmergeするのは禁止
- masterブランチを親に新規ブランチを作成するのは禁止

### どうしてもmasterやdevelopmentブランチに操作をしたい時は...
- force サブコマンドを用意してあります
```
git force checkout -b master
```

# インストール
```bash
git clone https://github.com/y-oga-819/git-extension.git

cd get-extension
make install # 要 /usr/local/bin/ 配下への write 権限

echo 'alias git=git-extension' >> ~/.zshrc # 自身の環境に応じてシェルプロファイルは適宜書き換えてください
exec $SHELL -l
```
