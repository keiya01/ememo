目標：出来るだけ簡単にメモを書いてtext fileに保存するCLIを作る

## Test of cmd package

### ユーザーから受け取ったスライス型のデータに3つ以上のデータが入っていることを確認する
- [x] 引数に受け取ったスライス型のデータに2つ以上のデータが入っていることを確認すること
- [x] スライス型のデータの数が3つ以下ならエラーを返すこと


### 拡張子txtの有無によって適切な値を返すことを確認する
- [x] 拡張子がついていなければ拡張子をつけて返すことを確認する
- [x] 拡張子がついていれば拡張子をつけずに返すことを確認する
- [x] 拡張子が複数ついている場合はひとつに変更することを確認する

### ユーザーからの入力を受け取ることを確認する
- [ ] 文字列で入力を受け取ることを確認する
- [ ] 入力が空ならエラーを返すことを確認する

## Test of files package

### テキストファイルからデータを読み取ることを確認する
- [x] 引数に指定したファイルを一行ずつ読み込む
- [x] 引数に指定した改行込みのファイルを一行ずつ読み込み出力する

## Test of cli package

### メモを入力できることを確認する
- [x] setフラグからユーザーの入力を受け取ること
- [x] -sフラグからユーザーの入力を受け取ること
- [x] 入力が空であればエラーを出力すること

### 入力された内容をtxtファイルに保存することを確認する
- [x] setフラグからユーザーの入力を受け取ったときに入力内容をファイルに保存すること
- [x] ファイルに保存した内容を出力すること
- [x] -sフラグからユーザーの入力を受け取ったときに入力内容をファイルに保存すること
- [x] ファイルに保存したら、保存に成功したことをメッセージで出力すること


### 結果によって文字の色を変えることを確認する
- [ ] エラーのとき文字の色を赤色にする
- [ ] 正常に保存したとき文字の色を緑色にする