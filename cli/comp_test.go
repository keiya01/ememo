package cli

import (
	"reflect"
	"testing"

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
				Value: "test.txt",
			},
		},
		{
			name: "引数にtestを渡したときにCompFlagのValueにtest.txtが格納されること",
			args: args{
				fileName: "test",
			},
			want: CompFlag{
				Value: "test.txt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cf := NewCompFlag(tt.args.fileName)
			if !reflect.DeepEqual(cf.Value, tt.want.Value) {
				test.MismatchErrorf(cf.Value, tt.want.Value, t)
			}
		})
	}
}
