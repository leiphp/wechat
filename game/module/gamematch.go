package module

import (
	"encoding/json"

	"github.com/leiphp/wechat/utils"
)


var GamematchEntity = Gamematch{}

type Gamematch struct {
	App utils.App
}

func (a *Gamematch) Init(app utils.App) *Gamematch {
	a.App = app
	return a
}

//CreateMatchRule 小游戏创建对局匹配规则，并返回一个matchid。每个小游戏可以创建多个matchid对应不同的匹配规则。小游戏持有的matchid数量上限为20。
//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/gamematch/gamematch.createMatchRule.html
func (a *Gamematch) CreateMatchRule(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/business/gamematch/creatematchrule", body, a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//DeleteMatchRule 小游戏删除对局匹配规则，每个规则对应一个唯一的matchid。
// 每个小游戏持有的matchid数量有限制，可以通过此接口删除无效的matchid。
// 删除接口在删除matchid的同时，也会释放掉matchid对应的匹配池。如果是线上正在使用的matchid，请谨慎使用该接口。
//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/gamematch/gamematch.deleteMatchRule.html
func (a *Gamematch) DeleteMatchRule(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/business/gamematch/deletematchrule", body, a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//Getallmatchrule 获取小游戏拥有的所有matchid及其对应的匹配规则，以及matchid的打开状态。
//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/gamematch/gamematch.getallmatchrule.html
func (a *Gamematch) Getallmatchrule(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/business/getallmatchrule", body, a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//SetMatchIdOpenState
// 修改matchid对应的打开状态。
// 每个matchid都有两种状态：打开(1)和关闭(0)
// 用户在调用加入匹配接口时，只能加入状态为打开的matchid。状态为打开的matchid才会分配匹配池。
// 小程序通过updateMatchRule修改matchid对应的规则时，只能修改状态为关闭的matchid
// 可以通过getAllMatchRule拉取小程序所有matchid的配置信息和打开状态。
// 将打开状态设置为关闭（0）时，匹配服务会释放掉matchid对应的匹配池。如果是线上正在使用的matchid，请谨慎变更matchid的状态。
//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/gamematch/gamematch.setMatchIdOpenState.html
func (a *Gamematch) SetMatchIdOpenState(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/business/gamematch/setmatchopenstate", body, a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//UpdateMatchRule 修改matchid对应的匹配规则。 小游戏修改matchid对应的规则时，只能修改状态为关闭的matchid。 matchid的状态可以通过setMatchIdOpenState修改。
//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/gamematch/gamematch.updateMatchRule.html
func (a *Gamematch) UpdateMatchRule(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/business/gamematch/updatematchrule", body, a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
