# Go Wechat

[![Go Doc](https://godoc.org/github.com/leiphp/wechat?status.svg)](https://godoc.org/github.com/leiphp/wechat)
[![Production Ready](https://img.shields.io/badge/production-ready-blue.svg)](https://github.com/leiphp/wechat)
[![License](https://img.shields.io/github/license/leiphp/wechat.svg?style=flat)](https://github.com/leiphp/wechat)

# Installation
```
go get -u -v github.com/leiphp/wechat
```
suggested using `go.mod`:
```
require github.com/leiphp/wechat
```

# Usage
```golang

appid := "xx"
secret := "xx"

//sample
app := miniapp.New(appid,secret)

//hook
var cacheToken utils.Token
app := New(appid,secret, func(appidAndAccessToken ...string) *utils.Token {
    if contextToken,err := utils.ExtractAppidAndAccessToken(appidAndAccessToken...);err == nil{
        // write token logic
        cacheToken.Token = contextToken.Token
        cacheToken.UpdateTime =  int(time.Now().Unix())
        return &cacheToken
    }
    //read token logic
    return &cacheToken
    })

//To Do
```

# License

`Go Wechat` is licensed under the [MIT License](LICENSE), 100% free and open-source, forever.