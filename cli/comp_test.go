package cli

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/keiya01/ememo/format"
	"github.com/keiya01/ememo/test"
)

func TestCompFlagが生成されることを確認するテスト(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    CompFlag
		wantErr bool
	}{
		{
			name: "引数にtest.txtを渡したときにCompFlagのValueにtest.txtが格納されること",
			args: args{
				fileName: "test.txt",
			},
			want: CompFlag{
				FileName: "test.txt",
			},
		},
		{
			name: "引数にtestを渡したときにCompFlagのValueにtest.txtが格納されること",
			args: args{
				fileName: "test",
			},
			want: CompFlag{
				FileName: "test.txt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cf := NewCompFlag(tt.args.fileName)
			if !reflect.DeepEqual(cf.FileName, tt.want.FileName) {
				test.MismatchErrorf(cf.FileName, tt.want.FileName, t)
			}
		})
	}
}

func Test指定されたファイルのTODOにチェックを入れられることを確認するテスト(t *testing.T) {
	type args struct {
		text     string
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
			name: "test.txtの中の「Hello World [ ]」の「[ ]」に「[x]」のようにチェックマークが入ることを確認する",
			args: args{
				text:     "Hello World",
				input:    "1",
				fileName: "test.txt",
			},
			want: "Hello World [x]\n",
		},
		{
			name: "test.txtの中の「Hello [ ]」の「[ ]」に「[x]」のようにチェックマークが入ることを確認する",
			args: args{
				text:     "Hello",
				input:    "1",
				fileName: "test.txt",
			},
			want: "Hello [x]\n",
		},
		{
			name: "入力値が数字以外のときにエラーを返すことを確認する",
			args: args{
				text:     "Hello",
				input:    "test",
				fileName: "test.txt",
			},
			want:    "Hello [x]\n",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test.InputValueCheck(tt.args.input, func() {
				cf := NewCompFlag(tt.args.fileName)

				fileData, _ := os.OpenFile(cf.FileName, os.O_RDWR|os.O_CREATE, 0666)
				defer os.Remove(cf.FileName)
				defer fileData.Close()

				contents := format.ChengeToMarkdown(tt.args.text)
				//書き込み処理
				fmt.Fprintln(fileData, contents)

				get, err := cf.FlagAction()
				if tt.wantErr && err == nil {
					test.NotOutputtedErrorf(err, t)
				}

				if !tt.wantErr && get != tt.want {
					test.MismatchErrorf(get, tt.want, t)
				}

			})

		})
	}
}
