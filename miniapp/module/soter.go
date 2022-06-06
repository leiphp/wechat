package module

import (
	"encoding/json"
	"github.com/leiphp/wechat/utils"
)


var SoterEntity = Soter{}

type Soter struct {
	App utils.App
}

func (a *Soter) Init(app utils.App) *Soter {
	a.App = app
	return a
}

//VerifySignature 生物认证秘钥签名验证
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/soter/soter.verifySignature.html
func (a *Soter) VerifySignature(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/cgi-bin/soter/verify_signature", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	return result , err
}
