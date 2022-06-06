package module

import (
    "encoding/json"
    "github.com/leiphp/wechat/utils"
)


var SubscribeMessageEntity = SubscribeMessage{}

type SubscribeMessage struct {
    App utils.App
}

func (a *SubscribeMessage) Init(app utils.App) *SubscribeMessage {
    a.App = app
    return a
}

//Send 发送订阅消息
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
func (a *SubscribeMessage) Send(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/cgi-bin/message/subscribe/send", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}
