package controller

import (
	"midgo/httpsvr"

	"midnightapisvr/comm"
	"midnightapisvr/service"
)

type ActionController struct{}

// GetCurrentAction API：获取当前执行的步骤
// [GET] /midnightapisvr/api/action/getcurrentaction
// :param token: token
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

// NormalMsgACK 出选择外的消息类型 ACK
// [POST] /midnightapisvr/api/action/normalmsgack
// {
// 	token: "xxx",
// 	sid: 1
// }
func (*ActionController) NormalMsgACK(c *httpsvr.Req) *httpsvr.Resp {
	var req = new(comm.MsgAckReq)
	httpsvr.JsonBodyDecode(c, req)
	var token = req.Token
	var sid = req.Sid

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
	ret.SetMsg("ack recieved :-)")
	return ret
}

// 选择 ACK
// [POST] /midnightapisvr/api/action/makechoice
// {
//     token: "xxx",
//     sid: 1,
//     option: 0 // 0为左边选项 1为右边选项
// }
func (*ActionController) MakeChoice(c *httpsvr.Req) *httpsvr.Resp {
	var req = new(comm.ChoiceReq)
	httpsvr.JsonBodyDecode(c, req)
	var token = req.Token
	var sid = req.Sid
	var option = req.Option

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
