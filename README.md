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
make install

echo 'alias git=git-extension' >> ~/.zshrc
```
