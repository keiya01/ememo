package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/keiya01/ememo/test"
)

func Testユーザーから受け取ったスライス型のデータに3つ以上のデータが入っていることを確認するテスト(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name    string
		args    args
		want    error
		wantErr bool
	}{
		{
			name: "引数に受け取ったスライス型のデータに3つ以上のデータが入っていることを確認すること",
			args: args{
				input: []string{
					"ememo",
					"text",
					"hello",
				},
			},
			want: nil,
		},
		{
			name: "スライス型のデータの数が3つ以下ならエラーを返すこと",
			args: args{
				input: []string{
					"ememo",
					"text",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckingUserInputArgumentValue(tt.args.input)
			if tt.wantErr && err == nil {
				test.NotOutputtedErrorf(err, t)
			}

			if !tt.wantErr && err != nil {
				t.Errorf("予期しないエラーが発生しています: err = %v", err)
			}
		})
	}
}

func Test拡張子txtの有無によって適切な値を返すことを確認する(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "拡張子がついていなければ拡張子をつけて返すことを確認する",
			args: args{
				fileName: "test",
			},
			want: "test.txt",
		},
		{
			name: "拡張子がついていれば拡張子をつけずに返すことを確認する",
			args: args{
				fileName: "test.txt",
			},
			want: "test.txt",
		},
		{
			name: "拡張子が複数ついている場合はひとつに変更することを確認する",
			args: args{
				fileName: "test.exec.txt.docs.txt",
			},
			want: "test.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			get := AddExtension(tt.args.fileName)
			if get != tt.want {
				test.MismatchErrorf(get, tt.want, t)
			}
		})
	}
}

func Testユーザーからの入力を受け取ることを確認するテスト(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "文字列で入力を受け取ることを確認する",
			args: args{
				input: "Hello World",
			},
			want: "Hello World",
		},
		{
			name: "文字列で入力を受け取ることを確認する",
			args: args{
				input: "Test Text",
			},
			want: "Test Text",
		},
		{
			name: "入力が空ならエラーを返すことを確認する",
			args: args{
				input: "",
			},
			want:    "入力値を空にすることは出来ません",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content := []byte(tt.args.input)
			tmpfile, err := ioutil.TempFile("", "example")
			if err != nil {
				log.Fatal(err)
			}
			defer os.Remove(tmpfile.Name()) // clean up

			//ファイルへの書き込む
			if _, err := tmpfile.Write(content); err != nil {
				log.Fatal(err)
			}

			//ファイル情報をtmpfileに格納する
			if _, err := tmpfile.Seek(0, 0); err != nil {
				log.Fatal(err)
			}

			//もともとのos.Stdin情報をoldStdinに格納しておき、
			//最後にoldStdinをos.Stdinに代入して初期化する
			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }() // Restore original Stdin

			//tmpfile情報をos.Stdinに代入することで、os.Stdinはポインタ型なので
			//GetUserInputValue()で参照することができる
			os.Stdin = tmpfile
			get, err := GetUserInputValue()
			if tt.wantErr && err == nil {
				test.NotOutputtedErrorf(err, t)
			}
			if !tt.wantErr && get != tt.want {
				test.MismatchErrorf(get, tt.want, t)
			}

			if err := tmpfile.Close(); err != nil {
				log.Fatal(err)
			}
		})
	}
}
