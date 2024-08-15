# ebitengine-tutorial

ebitenで作ったゲームをiOS用にビルドして実機で動かせるようにしてみる

## 構造

```
.
├── README.md
├── game
│   └── main.go
├── go.mod
├── go.sum
├── main.go
└── mobile
    ├── ios
    │   ├── Mobile.xcframework
    │   │   ├── Info.plist
    │   │   ├── ios-arm64
    │   │   └── ios-arm64_x86_64-simulator
    │   ├── tutorial
    │   └── tutorial.xcodeproj
    └── mobile.go

```

| ディレクトリ / ファイル | 説明 |
| - | - |
| game/ | 本体のプログラムを格納する |
| main.go | PC用に本体プログラムを呼び出す | 
| mobile/ios/Mobile.xcframework | iOS用に書き出したバイナリ | 
| mobile/ios/tutorial | iOSアプリ本体 | 
| mobile/ios/tutorial.xcodeprof | プロジェクトファイル | 
| mobile/mobile.go | モバイル用に本体プログラムを呼び出す | 

## iOSでのビルド

プロジェクトルートで下記を実行

```shlel
$ ebitenmobile bind -target ios -o ./mobile/ios/Mobile.xcframework ./mobile
```

mobile/ios/Mobile.xcframework/ios-arm64/Mobile.frameworkをxcodeprojでインポートして使用する。