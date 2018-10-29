package cmd

import "testing"

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
			err := CheckingUserInputValue(tt.args.input)
			if tt.wantErr && err == nil {
				t.Errorf("エラーが発生していません: err = %v", err)
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			get := AddExtension(tt.args.fileName)
			if get != tt.want {
				t.Errorf("値が一致していません: get = %s want = %s", get, tt.want)
			}
		})
	}
}
