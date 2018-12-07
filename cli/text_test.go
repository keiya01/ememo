package cli

import (
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
			name: "ユーザーの入力「hello.txt」をTextFlagに登録できることを確認する",
			args: args{
				input: "hello.txt",
			},
			want: TextFlag{
				Value: "hello.txt",
			},
		},
		{
			name: "ユーザーの入力「hello」をTextFlagに登録できることを確認する",
			args: args{
				input: "hello",
			},
			want: TextFlag{
				Value: "hello.txt",
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
		textFlag string
		inputVal []string
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
				textFlag: "test.txt",
				inputVal: []string{
					"Hello World",
				},
			},
			want: "Hello World [ ]\n",
		},
		{
			name: "-textフラグからユーザーの入力を受け取ったときに入力内容をファイルに保存すること",
			args: args{
				textFlag: "test",
				inputVal: []string{
					"Hello World",
				},
			},
			want: "Hello World [ ]\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//ファイルを上書きするとテストが通らないため初期化処理を行う
			fileName := input.AddExtension(tt.args.textFlag)
			defer os.Remove(fileName)

			tf := NewTextFlag(tt.args.textFlag)
			get := tf.save(tt.args.inputVal)
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
				input:    "",
				fileName: "test",
			},
			want:    "Hello World [ ]\n",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test.InputValueCheck(tt.args.input, func() {
				tf := NewTextFlag(tt.args.fileName)
				err := tf.FlagAction()
				defer os.Remove(tt.args.fileName)

				if tt.wantErr && err == nil {
					test.NotOutputtedErrorf(err, t)
				}

				get := file.PrintReadFile(tf.Value)
				if !tt.wantErr && get != tt.want {
					test.MismatchErrorf(get, tt.want, t)
				}

			})

		})
	}
}
