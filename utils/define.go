package utils

const domain Domain = "https://api.weixin.qq.com"

var expiredToken = MapBoolean{
	"42001":true, //40001 获取 access_token 时 AppSecret 错误，或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，或查看是否正在为恰当的公众号调用接口
	"40001":true, //42001 access_token 超时，请检查 access_token 的有效期，请参考基础支持 - 获取 access_token 中，对 access_token 的详细机制说明
}

type App interface {
	GetAccessToken(reflush ...bool) (*Token)
	GetConfig() (Config)
}

type Config struct {
	Appid  string
	Secret string
}

type Token struct {
	Token      string
	UpdateTime int
}

type ContextToken struct {
	Appid      string
	Token 	   string
}

//Hook [0] Appid [1] AccessToken
type Hook func(appidAndAccessToken ...string) *Token

type Domain string

type MapStr map[string]string

type Query map[string]string

type MapBoolean map[string]bool

type MapInterface map[string]interface{}
