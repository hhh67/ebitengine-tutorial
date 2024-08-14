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
    └── mobile.go
```

| ディレクトリ / ファイル | 説明 |
| - | - |
| game/ | 本体のプログラムを格納する |
| mobile/ | モバイル用に本体プログラムを呼び出す | 
| main.go | PC用に本体プログラムを呼び出す | 
