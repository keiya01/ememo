package cli

import (
	"os"
	"reflect"
	"testing"

	"github.com/keiya01/ememo/input"
)

func Testメモを入力できることを確認するテスト(t *testing.T) {
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
		{
			name: "入力が空であればエラーを出力すること",
			args: args{
				input: "",
			},
			want: TextFlag{
				Value: "Hello World",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tf, err := NewTextFlag(tt.args.input)

			if tt.wantErr && err == nil {
				t.Errorf("エラーが出力されていません")
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(tf.Value, tt.want.Value) {
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
			want: "Hello World\n",
		},
		{
			name: "-textフラグからユーザーの入力を受け取ったときに入力内容をファイルに保存すること",
			args: args{
				fileName: "test",
				textFlag: "Hello World",
			},
			want: "Hello World\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//ファイルを上書きするとテストが通らないため初期化処理を行う
			fileName := input.AddExtension(tt.args.fileName)
			defer os.Remove(fileName)

			var tf TextFlag
			var get string

			tf.Value = tt.args.textFlag
			get = tf.save(tt.args.fileName)
			if get != tt.want {
				t.Errorf("値が一致していません: get = %v, want = %v", get, tt.want)
			}

		})
	}
}
