package comm

type UserInfoRes struct {
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	AvatarURI string `json:"avatar_uri"`
}

// SessionInfo 会话信息
type SessionInfo struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type StoryBaseRes struct {
	Sid        int `json:"sid"`
	ConMsgType int `json:"conjunction_msg_type"`
}

type StoryMsgTextRes struct {
	Content string `json:"content"`
}

type StoryMsgChooseRes struct {
	LContent string `json:"l_content"`
	RContent string `json:"r_content"`
}

type StoryMsgNewsRes struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type StoryMsgMomentRes struct {
	Author  string `json:"author"`
	Content string `json:"content"`
	ImgURI  string `json:"img_uri"`
}
