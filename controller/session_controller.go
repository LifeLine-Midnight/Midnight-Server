package controller

import (
	"midgo/httpsvr"
	"midnightapisvr/service"
)

// SessionController 定义会话相关 api
type SessionController struct{}

// UserLogIn 用户登录
func (*SessionController) UserLogIn(c *httpsvr.Req) *httpsvr.Resp {
	c.ParseForm()
	username := c.FormValue("username")
	password := c.FormValue("password")

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
func (*SessionController) UserLogOut(c *httpsvr.Req) *httpsvr.Resp {
	c.ParseForm()
	token := c.FormValue("token")

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
