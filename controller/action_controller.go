package controller

import (
	"midgo/httpsvr"
	"midgo/logger"

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

	var actionSvc = new(service.ActionService)
	actionInfo, err := actionSvc.GetCurrentAction(uid)
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
	ret.SetData(actionInfo)
	logger.Info("%d: %s", rtn, msg)

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
	logger.Info(":param token: %s", token)
	logger.Info(":param sid: %d", sid)

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

	var actionSvc = new(service.ActionService)
	err = actionSvc.ConfirmUpdateProcess(uid, sid)
	if err != nil {
		var rtn = 500
		var msg = err.Error()

		ret.SetRtn(rtn)
		ret.SetMsg(msg)
		logger.Error("%d: %s", rtn, msg)

		return ret
	}

	var rtn = 0
	var msg = "ack recieved :-)"
	ret.SetRtn(rtn)
	ret.SetMsg(msg)
	logger.Info("%d: %s", rtn, msg)

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
	logger.Info(":param token: %s", token)
	logger.Info(":param sid: %d", sid)
	logger.Info(":param option: %d", option)

	var ret = new(httpsvr.Resp)
	if len(token) == 0 {
		var rtn = 403
		var msg = "permission denied: need token"

		ret.SetRtn(rtn)
		ret.SetMsg(msg)
		logger.Error("%d: %s", rtn, msg)

		return ret
	}

	if option != comm.LCHOOSE && option != comm.RCHOOSE {
		var rtn = 400
		var msg = "param error: invalid option"

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

	var actionSvc = new(service.ActionService)
	err = actionSvc.ChooseUpdateProcess(uid, sid, option)
	if err != nil {
		var rtn = 500
		var msg = err.Error()

		ret.SetRtn(rtn)
		ret.SetMsg(msg)
		logger.Error("%d: %s", rtn, msg)

		return ret
	}

	var rtn = 0
	var msg = "ack recieved :-)"
	ret.SetRtn(rtn)
	ret.SetMsg(msg)
	logger.Info("%d: %s", rtn, msg)

	return ret
}
