package cli

import (
	"reflect"
	"testing"
)

func Testメモを入力できることを確認するテスト(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name    string
		args    args
		want    MyFlags
		wantErr bool
	}{
		{
			name: "textフラグからユーザーの入力を受け取ること",
			args: args{
				input: []string{
					"ememo",
					"text",
					"Hello World",
				},
			},
			want: MyFlags{
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
			want: MyFlags{
				TextFlag: "Hello World",
			},
		},
		{
			name: "入力が空であればエラーを出力すること",
			args: args{
				input: nil,
			},
			want: MyFlags{
				TextFlag: "Hello World",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mf MyFlags
			var err error

			err = checkingUserInputValue(tt.args.input)
			if tt.wantErr && err == nil {
				t.Errorf("エラーが出力されていません: get = %s, want = %s", err, tt.want)
				return
			}

			StartCli(&mf, tt.args.input)
			if !tt.wantErr && !reflect.DeepEqual(mf.TextFlag, tt.want.TextFlag) {
				t.Errorf("値が一致していません: get = %v, want = %v", mf.TextFlag, tt.want.TextFlag)
			}

		})
	}
}
