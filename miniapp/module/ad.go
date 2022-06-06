package module

import (
	"github.com/leiphp/wechat/utils"
)


var AdEntity = Ad{}

type Ad struct {
	App utils.App
}

func (a *Ad) Init(app utils.App) *Ad {
	a.App = app
	return a
}
