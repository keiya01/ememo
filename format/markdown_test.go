package format

import (
	"testing"

	"github.com/keiya01/ememo/test"
)

func Testマークダウンによってフォーマットが変化することを確認するテスト(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "「-」を「●」に変換して出力する",
			args: args{
				input: "-Hello World",
			},
			want: " ● Hello World",
		},
		{
			name: "「-」を「●」に変換して出力する",
			args: args{
				input: "-Test",
			},
			want: " ● Test",
		},
		{
			name: "「=」を「◎」に変換して出力する",
			args: args{
				input: "=Hello World",
			},
			want: " ◎ Hello World",
		},
		{
			name: "「=」を「◎」に変換して出力する",
			args: args{
				input: "=Test",
			},
			want: " ◎ Test",
		},
		{
			name: "「;」を「1行の終わり」として出力する",
			args: args{
				input: "-Test;-Hello;-World;",
			},
			want: " ● Test\n ● Hello\n ● World\n",
		},
		{
			name: "「;」を「1行の終わり」として出力する",
			args: args{
				input: "-Test;=Hello;-World;",
			},
			want: " ● Test\n ◎ Hello\n ● World\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			get := ChengeToMarkdown(tt.args.input)
			if get != tt.want {
				test.MismatchErrorf(get, tt.want, t)
			}
		})
	}
}
