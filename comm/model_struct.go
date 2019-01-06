package comm

// 定义数据访问层的 struct
type UserInfoDB struct {
	Uid       int
	Username  string
	Nickname  string
	AvatarURI string
}

type UserConfirmInfo struct {
	Uid          int
	PasswordHash string
}

type StoryBaseDB struct {
	Sid        int
	ConMsgType int
	ConMid     int
	TimeDelay  int64
	NextSid    int
}

type StoryMsgTextDB struct {
	Mid     int
	Content string
}

type StoryMsgChooseDB struct {
	Mid      int
	LContent string
	RContent string
	LNextSid int
	RNextSid int
}

type StoryMsgNewsDB struct {
	Mid     int
	Title   string
	Content string
}

type StoryMsgMomentDB struct {
	Mid     int
	Author  string
	Content string
	ImgURI  string
}

type UserProcessDB struct {
	Uid            int
	CurSid         int
	CurProcessTime int64
}
