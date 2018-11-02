package cli

import (
	"os"
	"reflect"
	"testing"

	"github.com/keiya01/ememo/cmd"
)

func Testメモを入力できることを確認するテスト(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name    string
		args    args
		want    CliFlags
		wantErr bool
	}{
		{
			name: "-textフラグからユーザーの入力を受け取ること",
			args: args{
				input: []string{
					"ememo",
					"-text",
					"Hello World",
				},
			},
			want: CliFlags{
				TextFlag: "Hello World",
			},
		},
		{
			name: "-tフラグからユーザーの入力を受け取ること",
			args: args{
				input: []string{
					"ememo",
					"-t",
					"Hello World",
				},
			},
			want: CliFlags{
				TextFlag: "Hello World",
			},
		},
		{
			name: "入力が空であればエラーを出力すること",
			args: args{
				input: nil,
			},
			want: CliFlags{
				TextFlag: "Hello World",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var cf CliFlags

			err := StartCli(&cf, tt.args.input)
			if tt.wantErr && err == nil {
				t.Errorf("エラーが出力されていません")
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(cf.TextFlag, tt.want.TextFlag) {
				t.Errorf("値が一致していません: get = %v, want = %v", cf.TextFlag, tt.want.TextFlag)
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
			fileName := cmd.AddExtension(tt.args.fileName)
			defer os.Remove(fileName)

			var cf CliFlags
			var get string

			cf.TextFlag = tt.args.textFlag
			get = cf.save(tt.args.fileName)
			if get != tt.want {
				t.Errorf("値が一致していません: get = %v, want = %v", get, tt.want)
			}

		})
	}
}
