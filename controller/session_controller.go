package controller

import (
	"midgo/httpsvr"
	"midgo/logger"

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
	logger.Info(":param username: %s", username)

	var ret = new(httpsvr.Resp)
	if len(username) == 0 || len(password) == 0 {
		var rtn = 0
		var msg = "param error, need username and password"

		ret.SetRtn(rtn)
		ret.SetMsg(msg)
		logger.Error("%d: %s", rtn, msg)

		return ret
	}

	var sessionSvc = new(service.SessionService)
	sessionInfo, err := sessionSvc.UserLogIn(username, password)
	if err != nil {
		var rtn = 403
		var msg = err.Error()

		ret.SetRtn(rtn)
		ret.SetMsg(msg)
		logger.Error("%d: %s", rtn, msg)

		return ret
	}

	var rtn = 0
	var msg = "okay"
	ret.SetRtn(rtn)
	ret.SetMsg(msg)
	ret.SetData(sessionInfo)
	logger.Info("%d: %s", rtn, msg)

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
	logger.Info(":param token: %s", token)

	var ret = new(httpsvr.Resp)
	if len(token) == 0 {
		var rtn = 403
		var msg = "permission denied: need token"

		ret.SetRtn(rtn)
		ret.SetMsg(msg)
		logger.Error("%d: %s", rtn, msg)

		return ret
	}

	var sessionSvc = new(service.SessionService)
	err := sessionSvc.UserLogOut(token)
	if err != nil {
		var rtn = 500
		var msg = err.Error()

		ret.SetRtn(rtn)
		ret.SetMsg(msg)
		logger.Error("%d: %s", rtn, msg)

		return ret
	}

	var rtn = 0
	var msg = "logout successfully"
	ret.SetRtn(rtn)
	ret.SetMsg(msg)
	logger.Info("%d: %s", rtn, msg)

	return ret
}
