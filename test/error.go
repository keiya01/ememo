package test

import (
	"testing"
)

func MismatchErrorf(get, want interface{}, tb testing.TB) {
	tb.Helper()
	tb.Errorf("値が一致していません: get = %v want = %v", get, want)
}

func NotOutputtedErrorf(err error, tb testing.TB) {
	tb.Helper()
	tb.Errorf("エラーが出力されていません: err = %v", err)
}
