# 目標：出来るだけ簡単にメモを書いてtext fileに保存するCLIを作る


# Test of input package

### ユーザーから受け取ったスライス型のデータに3つ以上のデータが入っていることを確認する
- [x] 引数に受け取ったスライス型のデータに2つ以上のデータが入っていることを確認すること
- [x] スライス型のデータの数が3つ以下ならエラーを返すこと

### 拡張子txtの有無によって適切な値を返すことを確認する
- [x] 拡張子がついていなければ拡張子をつけて返すことを確認する
- [x] 拡張子がついていれば拡張子をつけずに返すことを確認する
- [x] 拡張子が複数ついている場合はひとつに変更することを確認する

### ユーザーからの入力を受け取ることを確認する
- [x] 文字列で入力を受け取ることを確認する
- [x] 入力が空ならエラーを返すことを確認する

### ユーザーの入力をTempファイルに保存することを確認する
- [ ] ユーザーの入力がTempファイルに保存されていることを確認する
- [ ] 入力が空なら保存せず、エラーを返すことを確認する

# Test of files package

### テキストファイルからデータを読み取ることを確認する
- [x] 引数に指定したファイルを一行ずつ読み込む
- [x] 引数に指定した改行込みのファイルを一行ずつ読み込み出力する


# Test of cli package

### メモを入力できることを確認する
- [x] 入力した内容が指定したテキストファイルに保存されていることを確認する
- [x] ファイル名が空のときエラーが出力されていることを確認する

### ユーザーの入力を受け取ることができることを確認する
- [x] ユーザーの入力「Hello World」をTextFlagに登録できることを確認する
- [x] ユーザーの入力「Hello」をTextFlagに登録できることを確認する
- [x] 入力が空であればエラーを出力すること

### 入力された内容をtxtファイルに保存することを確認する
- [x] --textフラグからユーザーの入力を受け取ったときに入力内容をファイルに保存すること
- [x] -tフラグからユーザーの入力を受け取ったときに入力内容をファイルに保存すること
- [x] ファイルに保存した内容を出力すること
- [x] ファイルに保存したら、保存に成功したことをメッセージで出力すること
- [ ] -textフラグから空の入力を受け取った時に、空のファイルを作成するか[ Y/n ]で尋ねる。Yを受け取ったら空のファイルを作成し、それ以外を受け取ったら作成しない。

### マークダウンのコマンドが表示されることを確認する（出力するだけなのでテスト書かない）
- [x] bool型の--markフラグを受け取ったときにマークダウンのコマンドを表示すること
- [x] bool型の-mフラグを受け取ったときにマークダウンのコマンドを表示すること

### 指定されたtxtファイルの中身を出力することを確認する
- [x] --readフラグを受けとったときに入力されたファイル名の中身を出力する
- [x] -rフラグを受けとったときに入力されたファイル名の中身を出力する

### 指定されたファイルのTODOにチェックを入れられることを確認する
- [ ] --compフラグを受け取ったときに指定されたファイルの指定されたTODOの「[ ]」に「[x]」とチェックマークが入ることを確認する
- [ ] -cフラグを受け取ったときに指定されたファイルの指定されたTODOの「[ ]」に「[x]」とチェックマークが入ることを確認する

### 指定されたtxtファイルの中身を空にする
- [ ] --emptyフラグを受けとったときに入力されたファイル名の中身を空にする
- [ ] -eフラグを受けとったときに入力されたファイル名の中身を空にする

# Test of color package

### 結果によって文字の色を変えることを確認する
- [x] github.com/fatih/colorで実装


# Test of format package

### マークダウンによってフォーマットが変化することを確認する
- [x] 「-」を「●」に変換して出力する
- [x] 「=」を「◎」に変換して出力する
- [x] 「;」を「1行の終わり」として出力する