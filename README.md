# go-timer

コマンドライン上で動作するストップウォッチ

# Features

CLI上でストップウォッチ機能、タイマー機能が利用できます。

# Installation

```bash
go get github.com/riku-smile/go-timer
or
git clone https://github.com/riku-smile/go-timer.git

go install
```

# Usage

*stopwatch

```bash
go-timer stopwatch
```

*timer
```bash
go-timer timer <time>
```

# Note

## stopwatch
  ストップウォッチ機能を提供します。
  Ctrl+Cで終了した地点までの時間を計測します。

## timer <time>
  タイマー機能を提供します。
  <time>の箇所に、半角数字を用いて()h<>m[]s の形でタイマーを設定してください。
	それぞれ、()時間<>分[]秒 という意味になります。
	例： 1h30m45s -> １時間３０分４５秒のタイマーを設定する。

# Author

* Riku-smile

# License

"go-timer" is under [MIT license](https://en.wikipedia.org/wiki/MIT_License).
