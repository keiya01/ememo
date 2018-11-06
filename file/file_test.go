package file

import (
	"fmt"
	"os"
	"testing"
)

func Testテキストファイルからデータを読み取ることを確認する(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "引数に指定したファイルを一行ずつ読み込み出力する",
			args: args{
				value: "Hello World",
			},
			want: "Hello World\n",
		},
		{
			name: "引数に指定した改行込みのファイルを一行ずつ読み込み出力する",
			args: args{
				value: "Hello World\nHello World",
			},
			want: "Hello World\nHello World\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//初期化処理
			setFile := "file_test.txt"
			file, _ := os.OpenFile(setFile, os.O_RDWR|os.O_CREATE, 0666)
			//ファイルが残っていると分かりにくくなるので最後に削除する
			defer os.Remove(setFile)
			defer file.Close()
			fmt.Fprintln(file, tt.args.value)

			get := PrintReadFile(setFile)
			if get != tt.want {
				t.Errorf("値が一致していません: get = %s, want = %s", get, tt.want)
			}
		})
	}
}
