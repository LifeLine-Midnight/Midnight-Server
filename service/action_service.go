package service

import (
	"errors"
	"time"

	"midnightapisvr/comm"
	"midnightapisvr/model"
)

type ActionService struct{}

// GetUpdateCurrentAction 获取当前进度
func (*ActionService) GetCurrentAction(uid int) (map[string]interface{}, error) {
	var processModel = new(model.ProcessModel)
	processInfo, err := processModel.GetUserProcess(uid)
	if err != nil {
		return nil, err
	}

	var actionInfo = make(map[string]interface{})

	// 判断是否到了执行时间
	if processInfo.CurProcessTime > time.Now().Unix() {
		actionInfo["base_info"] = &comm.StoryBaseRes{
			Sid:        processInfo.CurSid,
			ConMsgType: comm.MSG_NONE,
		}

		return actionInfo, nil
	}

	var storyModel = new(model.StoryModel)
	storyBaseInfo, err := storyModel.GetStoryBaseInfo(processInfo.CurSid)
	if err != nil {
		return nil, err
	}

	actionInfo["base_info"] = &comm.StoryBaseRes{
		Sid:        storyBaseInfo.Sid,
		ConMsgType: storyBaseInfo.ConMsgType,
	}

	var conInfo interface{}
	switch storyBaseInfo.ConMsgType {
	case comm.MSG_TEXT:
		msgInfo, err := storyModel.GetMsgText(storyBaseInfo.ConMid)
		if err != nil {
			return nil, err
		}
		conInfo = &comm.StoryMsgTextRes{
			Content: msgInfo.Content,
		}
	case comm.MSG_CHOOSE:
		chooseInfo, err := storyModel.GetMsgChoose(storyBaseInfo.ConMid)
		if err != nil {
			return nil, err
		}
		conInfo = &comm.StoryMsgChooseRes{
			LContent: chooseInfo.LContent,
			RContent: chooseInfo.RContent,
		}
	case comm.MSG_NEWS:
		newsInfo, err := storyModel.GetMsgNews(storyBaseInfo.ConMid)
		if err != nil {
			return nil, err
		}
		conInfo = &comm.StoryMsgNewsRes{
			Title:   newsInfo.Title,
			Content: newsInfo.Content,
		}
	case comm.MSG_MOMENT:
		momentInfo, err := storyModel.GetMsgMoment(storyBaseInfo.ConMid)
		if err != nil {
			return nil, err
		}
		conInfo = &comm.StoryMsgMomentRes{
			Author:  momentInfo.Author,
			Content: momentInfo.Content,
			ImgURI:  momentInfo.ImgURI,
		}
	case comm.MSG_ONLINE, comm.MSG_OFFLINE:
		conInfo = nil
	default:
		return nil, errors.New("invaild ConMsgType")
	}
	actionInfo["conjunction_info"] = conInfo

	return actionInfo, nil
}

// UpdateProcess 更新当前进度，除双选外
func (*ActionService) ConfirmUpdateProcess(uid int, sid int) error {
	var processModel = new(model.ProcessModel)
	processInfo, err := processModel.GetUserProcess(uid)
	if err != nil {
		return err
	}

	if processInfo.CurSid != sid {
		return errors.New("param sid != CurSid")
	}

	var storyModel = new(model.StoryModel)
	storyBaseInfo, err := storyModel.GetStoryBaseInfo(processInfo.CurSid)
	if err != nil {
		return err
	}

	if storyBaseInfo.ConMsgType == comm.MSG_CHOOSE {
		return errors.New("cannot confirm choose process")
	}

	nextExeTime := time.Now().Unix() + storyBaseInfo.TimeDelay
	err = processModel.UpdateUserProcess(uid, storyBaseInfo.NextSid, nextExeTime)
	if err != nil {
		return err
	}

	return nil
}

// ChooseUpdateProcess 获取用户选择，更新选择后的进度
func (*ActionService) ChooseUpdateProcess(uid int, sid int, option int) error {
	var processModel = new(model.ProcessModel)
	processInfo, err := processModel.GetUserProcess(uid)
	if err != nil {
		return err
	}

	if processInfo.CurSid != sid {
		return errors.New("param sid != CurSid")
	}

	var storyModel = new(model.StoryModel)
	storyBaseInfo, err := storyModel.GetStoryBaseInfo(processInfo.CurSid)
	if err != nil {
		return err
	}

	if storyBaseInfo.ConMsgType != comm.MSG_CHOOSE {
		return errors.New("cannot update choose process")
	}

	chooseInfo, err := storyModel.GetMsgChoose(storyBaseInfo.ConMid)
	if err != nil {
		return err
	}

	var nextSid int

	switch option {
	case comm.LCHOOSE:
		nextSid = chooseInfo.LNextSid
	case comm.RCHOOSE:
		nextSid = chooseInfo.RNextSid
	default:
		return errors.New("invalid option")
	}

	nextExeTime := time.Now().Unix() + storyBaseInfo.TimeDelay
	err = processModel.UpdateUserProcess(uid, nextSid, nextExeTime)
	if err != nil {
		return err
	}

	return nil
}
