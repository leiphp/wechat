package module

import (
	"encoding/json"
	"github.com/leiphp/wechat/utils"
)


var PhoneNumberEntity = PhoneNumber{}

type PhoneNumber struct {
	App utils.App
}

func (a *PhoneNumber) Init(app utils.App) *PhoneNumber {
	a.App = app
	return a
}

//GetPhoneNumber
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/phonenumber/phonenumber.getPhoneNumber.html
func (a *PhoneNumber) GetPhoneNumber(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/business/getuserphonenumber", body, a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
