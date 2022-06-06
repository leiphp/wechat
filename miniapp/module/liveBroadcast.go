package module

import (
	"github.com/leiphp/wechat/utils"
)


var LivebroadcastEntity = Livebroadcast{}

type Livebroadcast struct {
	App utils.App
}

func (a *Livebroadcast) Init(app utils.App) *Livebroadcast {
	a.App = app
	return a
}
