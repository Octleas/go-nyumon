# Go 入門編

Go 入門編の解答サンプルリポジトリ

## 環境構築

- Go 1.24.0 以上
- Visual Studio Code
- Git

上記のうち、Visual Studio Code と Git についてのインストール方法は、[Railway 準備編](https://techbowl.notion.site/Railway-ceba695d5014460e9733c2a46318cdec) をご確認いただき、挑戦の準備をしてください。
※ GitHub Codespaces についての資料はスキップしてください。

その他リポジトリの更新の仕方や、トラブルシューティングについても上記の Railway 準備編にございます。 何かあった際は、まずそちらを確認しましょう。

### Go セットアップ

下記 3つのいずれかの方法でインストールすることができます。

1. [インストーラーからのインストール](#インストーラーからのインストール)
2. [Homebrew](#homebrew)
3. [asdf](#asdf)

バージョン管理したいなど、特になければ [インストーラーからのインストール](#インストーラーからのインストール) で問題ありません。

#### インストーラーからのインストール

インストーラーからセットアップする場合は Windows, MacOS, Linux いずれも同じ手順です。

Go 公式の [All releases](https://go.dev/dl/) から自分が使用しているマシンの OS のインストーラーをダウンロードします。

バージョンは 1.24.0 以降にしてください。

![](https://github.com/TechBowl-japan/railway-assets/blob/main/go-nyumon/installer-page.png?raw=true)

- Apple Silicon Mac (M1など)
  - Kind: Installer
  - OS: macOS
  - Arch: ARM64
- Intel Mac
  - Kind: Installer
  - OS: macOS
  - Arch: x86_64
- Windows(Intel CPUなど)
  - Kind: Installer
  - OS: Windows
  - Arch: x86
- Windows(AMD CPUなど)
  - Kind: Installer
  - OS: Windows
  - Arch: x86_64

ダウンロードしたファイルを実行し、画面に従ってインストールをしてください。

#### Homebrew

```shell
$ brew install go
```

#### asdf

```shell
$ asdf plugin add golang https://github.com/asdf-community/asdf-golang.git
$ asdf install golang 1.24.0
```

## ディレクトリ構成

```
go-nyumon
├── app ... station6から作成するcliツール
├── stations ... station5までのコード
│   ├── station01
│   ├── station02
│   ├── station03
│   ├── station04
│   └── station05
├── tests
├── README.md
├── go.mod
└── go.sum
```
