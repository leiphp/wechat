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

//Generate JSON
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.generate.html
func (a *UrlScheme) Generate(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/generatescheme", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}
