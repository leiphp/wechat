package module

import (
	"encoding/json"

	"github.com/leiphp/wechat/utils"
)


var UrlSchemeEntity = UrlScheme{}

type UrlScheme struct {
	App utils.App
}

func (a *UrlScheme) Init(app utils.App) *UrlScheme {
	a.App = app
	return a
}

//Generate 获取小程序 scheme 码，适用于短信、邮件、外部网页、微信内等拉起小程序的业务场景。通过该接口，可以选择生成到期失效和永久有效的小程序码，目前仅针对国内非个人主体的小程序开放，详见获取 URL scheme。
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/UrlScheme.generate.html
func (a *UrlScheme) Generate(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/generatescheme", body, a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
