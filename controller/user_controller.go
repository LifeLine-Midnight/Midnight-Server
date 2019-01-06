package controller

import (
	"fmt"
	"midgo/httpsvr"

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
	fmt.Println(req)

	var username = req.Username
	var password = req.Password
	var nickname = req.Nickname

	var ret = new(httpsvr.Resp)
	if len(username) < 6 || len(password) < 6 || len(nickname) == 0 {
		ret.SetRtn(400)
		ret.SetMsg("param error, check if username, password length < 6")
		return ret
	}

	var userSvc = new(service.UserService)
	const AVATAR_URI = "/midnightstatic/images/avatar_default.jpg"
	userInfo, err := userSvc.UserRegister(username, nickname, password, AVATAR_URI)
	if err != nil {
		ret.SetRtn(500)
		ret.SetMsg(err.Error())
		return ret
	}

	ret.SetRtn(0)
	ret.SetMsg("okay")
	ret.SetData(userInfo)
	return ret
}

// GetUserInfo 获取用户信息
// [GET] /midnightapisvr/api/user/getuserinfo
// :param token: 登录时候得到的 token
func (*UserController) GetUserInfo(c *httpsvr.Req) *httpsvr.Resp {
	c.ParseForm()
	token := c.FormValue("token")

	var ret = new(httpsvr.Resp)
	if len(token) == 0 {
		ret.SetRtn(403)
		ret.SetMsg("permission denied: need token")
		return ret
	}

	var sessionSvc = new(service.SessionService)
	uid, err := sessionSvc.GetUidByToken(token)
	if err != nil || uid <= 0 {
		ret.SetRtn(403)
		ret.SetMsg("permission denied: invalid token")
		return ret
	}

	var userSvc = new(service.UserService)
	userInfo, err := userSvc.GetUserInfo(uid)
	if err != nil {
		ret.SetRtn(500)
		ret.SetMsg(err.Error())
		return ret
	}

	ret.SetRtn(0)
	ret.SetMsg("okay")
	ret.SetData(userInfo)

	return ret
}
