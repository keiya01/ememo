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
			name: "textフラグからユーザーの入力を受け取ること",
			args: args{
				input: []string{
					"ememo",
					"set",
					"Hello World",
				},
			},
			want: CliFlags{
				SetFlag: "Hello World",
			},
		},
		{
			name: "-tフラグからユーザーの入力を受け取ること",
			args: args{
				input: []string{
					"ememo",
					"-s",
					"Hello World",
				},
			},
			want: CliFlags{
				SetFlag: "Hello World",
			},
		},
		{
			name: "入力が空であればエラーを出力すること",
			args: args{
				input: nil,
			},
			want: CliFlags{
				SetFlag: "Hello World",
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

			if !tt.wantErr && !reflect.DeepEqual(cf.SetFlag, tt.want.SetFlag) {
				t.Errorf("値が一致していません: get = %v, want = %v", cf.SetFlag, tt.want.SetFlag)
			}

		})
	}
}

func Test入力された内容をtxtファイルに保存することを確認するテスト(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name     string
		args     args
		want     string
		testType string
		wantErr  bool
	}{
		{
			name: "ファイルに保存した内容を出力すること",
			args: args{
				fileName: "test.txt",
			},
			want:     "Hello World",
			testType: "read",
		},
		{
			name: "textフラグからユーザーの入力を受け取ったときに入力内容をファイルに保存すること",
			args: args{
				fileName: "test.txt",
			},
			want:     "Hello World",
			testType: "save",
		},
		{
			name: "textフラグからユーザーの入力を受け取ったときに入力内容をファイルに保存すること",
			args: args{
				fileName: "test",
			},
			want:     "Hello World",
			testType: "save",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//ファイルを上書きするとテストが通らないため初期化処理を行う
			fileName := cmd.AddExtension(tt.args.fileName)
			os.Remove(fileName)

			var cf CliFlags
			var get string

			cf.SetFlag = tt.want
			get = cf.saveInputText(tt.args.fileName)
			if tt.testType == "save" && get != tt.want {
				t.Errorf("値が一致していません: get = %v, want = %v", get, tt.want)
			}

			get = printReadFile(tt.args.fileName)
			if tt.testType == "read" && get != tt.want {
				t.Errorf("値が一致していません: get = %v, want = %v", get, tt.want)
			}

		})
	}
}
