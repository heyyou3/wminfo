package window

import (
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/icccm"
)

type Reader interface {
	ClientListGet(xu *xgbutil.XUtil) ([]xproto.Window, error)
	WmClassGet(xu *xgbutil.XUtil, win xproto.Window) (*icccm.WmClass, error)
	WmNameGet(xu *xgbutil.XUtil, win xproto.Window) (string, error)
}

type Info struct {
	ID      xproto.Window
	Name    string
	WmClass *icccm.WmClass
}

type Client struct {
	x  *xgbutil.XUtil
	wr Reader
}

type windowReaderImpl struct {
}

func (w *windowReaderImpl) ClientListGet(xu *xgbutil.XUtil) ([]xproto.Window, error) {
	return ewmh.ClientListGet(xu)
}

func (w *windowReaderImpl) WmClassGet(xu *xgbutil.XUtil, win xproto.Window) (*icccm.WmClass, error) {
	return icccm.WmClassGet(xu, win)
}

func (w *windowReaderImpl) WmNameGet(xu *xgbutil.XUtil, win xproto.Window) (string, error) {
	return ewmh.WmNameGet(xu, win)
}

func New() (*Client, error) {
	x, err := xgbutil.NewConn()

	if err != nil {
		return nil, err
	}

	return &Client{
		x:  x,
		wr: &windowReaderImpl{},
	}, nil
}

func (c *Client) FetchWindowInfo() ([]*Info, error) {
	wids, err := c.wr.ClientListGet(c.x)

	if err != nil {
		return nil, err
	}

	var res []*Info

	for _, wid := range wids {
		name, err := c.wr.WmNameGet(c.x, wid)
		if err != nil {
			continue
		}

		class, err := c.wr.WmClassGet(c.x, wid)
		if err != nil {
			continue
		}

		res = append(res, &Info{
			ID:      wid,
			Name:    name,
			WmClass: class,
		})
	}

	return res, nil
}
