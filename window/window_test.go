package window

import (
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/icccm"
	"testing"
)

type mockWindowReader struct {
}

func (m *mockWindowReader) ClientListGet(xu *xgbutil.XUtil) ([]xproto.Window, error) {
	_ = xu
	return []xproto.Window{123456, 789012}, nil
}

func (m *mockWindowReader) WmClassGet(xu *xgbutil.XUtil, win xproto.Window) (*icccm.WmClass, error) {
	_, _ = xu, win
	return &icccm.WmClass{Class: "Vivaldi"}, nil
}

func (m *mockWindowReader) WmNameGet(xu *xgbutil.XUtil, win xproto.Window) (string, error) {
	_, _ = xu, win
	return "hoge", nil
}

func TestFetchWindowInfo(t *testing.T) {
	cases := []struct {
		desc     string
		expected []*Info
	}{
		{desc: "期待する複数の Window 情報が返却されること", expected: []*Info{
			{ID: 123456, Name: "hoge", WmClass: &icccm.WmClass{Class: "Vivaldi"}},
			{ID: 789012, Name: "hoge", WmClass: &icccm.WmClass{Class: "Vivaldi"}},
		},
		},
	}
	for i, tc := range cases {
		x := &xgbutil.XUtil{}
		client := &Client{x: x, wr: &mockWindowReader{}}
		res, _ := client.FetchWindowInfo()
		actual := res[i]
		if actual.ID != tc.expected[i].ID {
			t.Fatalf("Window ID が一致しない %s: expected: %v got: %v",
				tc.desc, tc.expected[i], actual)
		}
		if actual.Name != tc.expected[i].Name {
			t.Fatalf("Window名が一致しない %s: expected: %v got: %v",
				tc.desc, tc.expected[i], actual)
		}
		if actual.WmClass.Class != tc.expected[i].WmClass.Class {
			t.Fatalf("クラス名が一致しない %s: expected: %v got: %v",
				tc.desc, tc.expected[i], actual)
		}
		if len(res) != len(tc.expected) {
			t.Fatalf("配列の数が一致しない %s: expected: %v got: %v",
				tc.desc, tc.expected[i], actual)
		}
	}
}
