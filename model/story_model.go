package model

import (
	"midnightapisvr/comm"
)

type StoryModel struct{}

func (*StoryModel) GetStoryBaseInfo(sid int) (*comm.StoryBaseDB, error) {
	conn, err := GetMidnightDBConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var ssql = `SELECT sid, conjunction_msg_type, conjunction_mid, time_delay, next_sid
				FROM story WHERE sid=?`
	var storyItem = new(comm.StoryBaseDB)

	err = conn.QueryRow(ssql, sid).Scan(&storyItem.Sid, &storyItem.ConMsgType,
		&storyItem.ConMid, &storyItem.TimeDelay, &storyItem.NextSid)
	if err != nil {
		return nil, err
	}

	return storyItem, nil
}

func (*StoryModel) GetMsgText(mid int) (*comm.StoryMsgTextDB, error) {
	conn, err := GetMidnightDBConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var ssql = `SELECT mid, content FROM msg_text WHERE mid=?`
	var msgTextInfo = new(comm.StoryMsgTextDB)
	err = conn.QueryRow(ssql, mid).Scan(&msgTextInfo.Mid, &msgTextInfo.Content)
	if err != nil {
		return nil, err
	}

	return msgTextInfo, nil
}

func (*StoryModel) GetMsgChoose(mid int) (*comm.StoryMsgChooseDB, error) {
	conn, err := GetMidnightDBConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var ssql = `SELECT mid, l_content, r_content, l_next_sid, r_next_sid
				FROM msg_choose WHERE mid=?`
	var msgChooseInfo = new(comm.StoryMsgChooseDB)
	err = conn.QueryRow(ssql, mid).Scan(&msgChooseInfo.Mid,
		&msgChooseInfo.LContent, &msgChooseInfo.RContent,
		&msgChooseInfo.LNextSid, &msgChooseInfo.RNextSid)

	if err != nil {
		return nil, err
	}

	return msgChooseInfo, nil
}

func (*StoryModel) GetMsgNews(mid int) (*comm.StoryMsgNewsDB, error) {
	conn, err := GetMidnightDBConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var ssql = `SELECT mid, title, content FROM msg_news WHERE mid=?`
	var msgNewsInfo = new(comm.StoryMsgNewsDB)
	err = conn.QueryRow(ssql, mid).Scan(&msgNewsInfo.Mid,
		&msgNewsInfo.Title, &msgNewsInfo.Content)

	if err != nil {
		return nil, err
	}

	return msgNewsInfo, nil
}

func (*StoryModel) GetMsgMoment(mid int) (*comm.StoryMsgMomentDB, error) {
	conn, err := GetMidnightDBConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var ssql = `SELECT mid, author, content, img_uri FROM msg_moment WHERE mid=?`
	var msgMomentInfo = new(comm.StoryMsgMomentDB)
	err = conn.QueryRow(ssql, mid).Scan(&msgMomentInfo.Mid,
		&msgMomentInfo.Author, &msgMomentInfo.Content, &msgMomentInfo.ImgURI)

	if err != nil {
		return nil, err
	}

	return msgMomentInfo, nil
}
