package cli

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/keiya01/ememo/file"
	"github.com/keiya01/ememo/input"
	"github.com/keiya01/ememo/test"
)

func Testユーザーの入力を受け取ることができることを確認するテスト(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    TextFlag
		wantErr bool
	}{
		{
			name: "ユーザーの入力「Hello World」をTextFlagに登録できることを確認する",
			args: args{
				input: "Hello World",
			},
			want: TextFlag{
				Value: "Hello World",
			},
		},
		{
			name: "ユーザーの入力「Hello」をTextFlagに登録できることを確認する",
			args: args{
				input: "Hello",
			},
			want: TextFlag{
				Value: "Hello",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tf := NewTextFlag(tt.args.input)

			if !reflect.DeepEqual(tf.Value, tt.want.Value) {
				t.Errorf("値が一致していません: get = %v, want = %v", tf.Value, tt.want.Value)
			}

		})
	}
}

func Test入力された内容をtxtファイルに保存することを確認するテスト(t *testing.T) {
	type args struct {
		fileName string
		textFlag string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "-textフラグからユーザーの入力を受け取ったときに入力内容をファイルに保存すること",
			args: args{
				fileName: "test.txt",
				textFlag: "Hello World",
			},
			want: "Hello World [ ]\n",
		},
		{
			name: "-textフラグからユーザーの入力を受け取ったときに入力内容をファイルに保存すること",
			args: args{
				fileName: "test",
				textFlag: "Hello World",
			},
			want: "Hello World [ ]\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//ファイルを上書きするとテストが通らないため初期化処理を行う
			fileName := input.AddExtension(tt.args.fileName)
			defer os.Remove(fileName)

			var get string

			tf := NewTextFlag(tt.args.textFlag)
			get = tf.save(tt.args.fileName)
			if get != tt.want {
				t.Errorf("値が一致していません: get = %v, want = %v", get, tt.want)
			}

		})
	}
}

func Testメモを入力できることを確認するテスト(t *testing.T) {
	type args struct {
		input    string
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "入力した内容が指定したテキストファイルに保存されていることを確認する",
			args: args{
				input:    "Hello World",
				fileName: "test.txt",
			},
			want: "Hello World [ ]\n",
		},
		{
			name: "入力した内容が指定したテキストファイルに保存されていることを確認する",
			args: args{
				input:    "Hello",
				fileName: "test.txt",
			},
			want: "Hello [ ]\n",
		},
		{
			name: "ファイル名が空のときエラーが出力されていることを確認する",
			args: args{
				input:    "Hello World",
				fileName: "",
			},
			want:    "Hello World [ ]\n",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Tempファイルの作成
			tmpfile, err := ioutil.TempFile("", "example")
			if err != nil {
				log.Fatal(err)
			}
			defer os.Remove(tmpfile.Name()) // clean up

			//Tempファイルへ書き込む
			content := []byte(tt.args.fileName)
			if _, err := tmpfile.Write(content); err != nil {
				log.Fatal(err)
			}

			//Tempファイル情報をtmpfileに格納する
			if _, err := tmpfile.Seek(0, 0); err != nil {
				log.Fatal(err)
			}

			//もともとのos.Stdin情報をoldStdinに格納しておき、
			//最後にoldStdinをos.Stdinに代入して初期化する
			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }() // Restore original Stdin

			//tmpfile情報をos.Stdinに代入することで、os.Stdinはポインタ型なので
			//関数内でos.Stdinを読み込むことで参照することができる
			os.Stdin = tmpfile

			tf := NewTextFlag(tt.args.input)
			err = tf.FlagAction()
			defer os.Remove(tt.args.fileName)

			if tt.wantErr && err == nil {
				test.NotOutputtedErrorf(err, t)
			}

			get := file.PrintReadFile(tt.args.fileName)
			if !tt.wantErr && get != tt.want {
				test.MismatchErrorf(get, tt.want, t)
			}

			if err := tmpfile.Close(); err != nil {
				log.Fatal(err)
			}

		})
	}
}
