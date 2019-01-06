package controller

import (
	"midgo/httpsvr"
	"midnightapisvr/comm"
	"midnightapisvr/service"
)

// SessionController 定义会话相关 api
type SessionController struct{}

// UserLogIn 用户登录
// [POST] /midnightapisvr/api/session/userlogin
// {
//     username: "guest123",
//     password: "guest123#"
// }
func (*SessionController) UserLogIn(c *httpsvr.Req) *httpsvr.Resp {
	var req = new(comm.UserLogInReq)
	httpsvr.JsonBodyDecode(c, req)

	var username = req.Username
	var password = req.Password

	var ret = new(httpsvr.Resp)
	if len(username) == 0 || len(password) == 0 {
		ret.SetRtn(400)
		ret.SetMsg("param error")
		return ret
	}

	var sessionSvc = new(service.SessionService)
	sessionInfo, err := sessionSvc.UserLogIn(username, password)
	if err != nil {
		ret.SetRtn(403)
		ret.SetMsg(err.Error())
		return ret
	}

	ret.SetRtn(0)
	ret.SetMsg("okay")
	ret.SetData(sessionInfo)
	return ret
}

// UserLogOut 用户登出
// [POST] /midnightapisvr/api/session/userlogout
// {
//     token: "xxx"
// }
func (*SessionController) UserLogOut(c *httpsvr.Req) *httpsvr.Resp {
	var req = new(comm.UserLogOutReq)
	httpsvr.JsonBodyDecode(c, req)
	var token = req.Token

	var ret = new(httpsvr.Resp)
	if len(token) == 0 {
		ret.SetRtn(403)
		ret.SetMsg("permission denied: need token")
		return ret
	}

	var sessionSvc = new(service.SessionService)
	err := sessionSvc.UserLogOut(token)
	if err != nil {
		ret.SetRtn(500)
		ret.SetMsg(err.Error())
		return ret
	}

	ret.SetRtn(0)
	ret.SetMsg("logout successfully")
	return ret
}
