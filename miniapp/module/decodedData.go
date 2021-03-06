package module

import (
	"github.com/leiphp/wechat/utils"
)


var DecodedDataEntity = DecodedData{}

type DecodedData struct {
	App utils.App
}

func (a *DecodedData) Init(app utils.App) *DecodedData {
	a.App = app
	return a
}

//DecodeCryptoData 解密数据
func (a *DecodedData) DecodeCryptoData(sessionKey, encryptedData, iv string) (interface{}, error) {
	wxCrypt := WxBizDataCrypt{
		AppId:      a.App.GetConfig().Appid,
		SessionKey: sessionKey,
	}
	result, err := wxCrypt.Decrypt(encryptedData, iv, true)
	return result, err
}