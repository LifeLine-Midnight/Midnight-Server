package controller

import (
	"midgo/httpsvr"
	"midgo/logger"

	"midnightapisvr/comm"
	"midnightapisvr/service"
)

// UserController 定义用户相关信息的 api
type UserController struct{}

// UserRegister 用户注册
// [POST] /midnightapisvr/api/user/userregister
// Request Body (application/json) {
//     username: "guest123",
//     nickname: "john",
//     password: "guest123#",
// }
func (*UserController) UserRegister(c *httpsvr.Req) *httpsvr.Resp {
	var req = new(comm.UserRegisterReq)
	httpsvr.JsonBodyDecode(c, req)

	var username = req.Username
	var password = req.Password
	var nickname = req.Nickname
	logger.Info(":param username: %s", username)
	logger.Info(":param nickname: %s", nickname)

	var ret = new(httpsvr.Resp)
	if len(username) < 6 || len(password) < 6 || len(nickname) == 0 {
		var rtn = 400
		var msg = "param error, check if username, password length < 6"

		ret.SetRtn(rtn)
		ret.SetMsg(msg)
		logger.Error("%d: %s", rtn, msg)

		return ret
	}

	var userSvc = new(service.UserService)
	const AVATAR_URI = "/midnightstatic/images/avatar_default.jpg"
	userInfo, err := userSvc.UserRegister(username, nickname, password, AVATAR_URI)
	if err != nil {
		var rtn = 500
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
	ret.SetData(userInfo)
	logger.Info("%d: %s", rtn, msg)

	return ret
}

// GetUserInfo 获取用户信息
// [GET] /midnightapisvr/api/user/getuserinfo
// :param token: 登录时候得到的 token
func (*UserController) GetUserInfo(c *httpsvr.Req) *httpsvr.Resp {
	c.ParseForm()
	token := c.FormValue("token")
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
	uid, err := sessionSvc.GetUidByToken(token)
	if err != nil || uid <= 0 {
		var rtn = 403
		var msg = "permission denied: invalid token"

		ret.SetRtn(rtn)
		ret.SetMsg(msg)
		logger.Error("%d: %s", rtn, msg)

		return ret
	}

	var userSvc = new(service.UserService)
	userInfo, err := userSvc.GetUserInfo(uid)
	if err != nil {
		var rtn = 500
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
	ret.SetData(userInfo)
	logger.Info("%d: %s", rtn, msg)

	return ret
}
