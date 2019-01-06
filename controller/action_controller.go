package controller

import (
	"strconv"

	"midgo/httpsvr"

	"midnightapisvr/comm"
	"midnightapisvr/service"
)

type ActionController struct{}

func (*ActionController) GetCurrentAction(c *httpsvr.Req) *httpsvr.Resp {
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

	var actionSvc = new(service.ActionService)
	actionInfo, err := actionSvc.GetCurrentAction(uid)
	if err != nil {
		ret.SetRtn(500)
		ret.SetMsg(err.Error())
		return ret
	}

	ret.SetRtn(0)
	ret.SetMsg("okay")
	ret.SetData(actionInfo)
	return ret
}

func (*ActionController) NormalMsgACK(c *httpsvr.Req) *httpsvr.Resp {
	c.ParseForm()
	token := c.FormValue("token")
	sid, _ := strconv.Atoi(c.FormValue("sid"))

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

	var actionSvc = new(service.ActionService)
	err = actionSvc.ConfirmUpdateProcess(uid, sid)
	if err != nil {
		ret.SetRtn(500)
		ret.SetMsg(err.Error())
		return ret
	}

	ret.SetRtn(0)
	ret.SetMsg("recieve ack :-)")
	return ret
}

func (*ActionController) MakeChoice(c *httpsvr.Req) *httpsvr.Resp {
	c.ParseForm()
	token := c.FormValue("token")
	sid, _ := strconv.Atoi(c.FormValue("sid"))
	option, _ := strconv.Atoi(c.FormValue("option"))

	var ret = new(httpsvr.Resp)
	if len(token) == 0 {
		ret.SetRtn(403)
		ret.SetMsg("permission denied: need token")
		return ret
	}

	if option != comm.LCHOOSE && option != comm.RCHOOSE {
		ret.SetRtn(400)
		ret.SetMsg("param error: invalid option")
		return ret
	}

	var sessionSvc = new(service.SessionService)
	uid, err := sessionSvc.GetUidByToken(token)
	if err != nil || uid <= 0 {
		ret.SetRtn(403)
		ret.SetMsg("permission denied: invalid token")
		return ret
	}

	var actionSvc = new(service.ActionService)
	err = actionSvc.ChooseUpdateProcess(uid, sid, option)
	if err != nil {
		ret.SetRtn(500)
		ret.SetMsg(err.Error())
		return ret
	}

	ret.SetRtn(0)
	ret.SetMsg("recieve ack :-)")
	return ret
}
