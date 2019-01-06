# midnightapisvr API

* Host: golfqiu.cn

## 1. 用户注册

[POST] /midnightapisvr/api/user/userregister

**Request Body (application/json)**
```
{
    username: "guest123",
    nickname: "john",
    password: "guest123#"
}
```

**Response (application/json)**
```
{
    rtn: 0,
    msg: "okay",
    data: {
        username: "guset123",
        nickname: "John",
        avatar_uri: "/midnightstatic/images/avatar_default.jpg"
    }
}
```

---

## 2. 用户登录

[POST] /midnightapisvr/api/session/userlogin

**Request Body (application/json)**
```
{
    username: "guest123",
    password: "guest123#"
}
```

**Response (application/json)**
```
{
    "rtn": 0,
    "msg": "okay",
    "data": {
        "username": "guest123",
        "token": "fcf41657f02f88137a1bcf068a32c0a31546766298712192281"
    }
}
```

---

## 3. 用户信息获取

[GET] /midnightapisvr/api/user/getuserinfo

**Request**
```
:param token: "fcf41657f02f88137a1bcf068a32c0a31546766298712192281"
```

**Response (application/json)**
```
{
    "rtn": 0,
    "msg": "okay",
    "data": {
        "username": "guest123",
        "nickname": "john",
        "avatar_uri": "/midnightstatic/images/avatar_default.jpg"
    }
}
```

---

## 4. 获取当前执行的 action

[GET] /midnightapisvr/api/action/getcurrentaction

**Request**
```
:param token: "fcf41657f02f88137a1bcf068a32c0a31546766298712192281"
```

**Response(application/json)**
```
data.base_info.conjunction_msg_type 的值，共有好多种返回值

// conjunction_msg_type = 0 什么都不用干
{
    "rtn": 0,
    "msg": "okay",
    "data": {
        "base_info": {
            "sid": 1,
            "conjunction_msg_type": 0
        },
        "conjunction_info": null
    }
}

// conjunction_msg_type = 1 对方发来文字消息
{
    "rtn": 0,
    "msg": "okay",
    "data": {
        "base_info": {
            "sid": 2,
            "conjunction_msg_type": 1
        },
        "conjunction_info": {
            "content": "请问，你看得到吗"
        }
    }
}

// conjunction_msg_type = 2 选择题
{
    "rtn": 0,
    "msg": "okay",
    "data": {
        "base_info": {
            "sid": 4,
            "conjunction_msg_type": 2
        },
        "conjunction_info": {
            "l_content": "你是谁啊",
            "r_content": "什么鬼软件，不经过我同意就加好友？"
        }
    }
}

// conjunction_msg_type = 3 新闻
{
    "rtn": 0,
    "msg": "okay",
    "data": {
        "base_info": {
            "sid": 8,
            "conjunction_msg_type": 3
        },
        "conjunction_info": {
            "title": "打代码可防止脱发",
            "content": "xxxx, xxxx 新闻正文..."
        }
    }
}

// conjunction_msg_type = 4 朋友圈更新
{
    "rtn": 0,
    "msg": "okay",
    "data": {
        "base_info": {
            "sid": 9,
            "conjunction_msg_type": 4
        },
        "conjunction_info": {
            "author": "tfboys",
            "content": "这是一条朋友圈的内容，下面是图片的 uri 哦",
            "img_uri": "/midnightstatic/images/moment/moment_default.jpg"
        }
    }
}

// conjunction_msg_type = 5 对方上线
{
    "rtn": 0,
    "msg": "okay",
    "data": {
        "base_info": {
            "sid": 1,
            "conjunction_msg_type": 5
        },
        "conjunction_info": null
    }
}

// conjunction_msg_type = 6 对方下线
{
    "rtn": 0,
    "msg": "okay",
    "data": {
        "base_info": {
            "sid": 1,
            "conjunction_msg_type": 6
        },
        "conjunction_info": null
    }
}
```

---

以下两个接口都为 ACK 接口，分别对应普通消息类型和选择消息类型。

在前台调用 getcurrentaction 接口时，后台默认不进行游戏进度更新，直到前台调用 ACK 后，后台会进行进度更新。

## 5. 普通消息 ACK

适用于除了选择消息以外的所有消息类型的 ACK。

[POST] /midnightapisvr/api/action/normalmsgack

**Request(application/json)**
```
{
    token: "fcf41657f02f88137a1bcf068a32c0a31546766298712192281",
    sid: 1
}
```

**Response(application/json)**
```
{
    "rtn": 0,
    "msg": "ack recieved :-)",
    "data": null
}
```

---

## 6. 选择消息 ACK

[POST] /midnightapisvr/api/action/makechoice

**Request(application/json)**
```
{
    token: "fcf41657f02f88137a1bcf068a32c0a31546766298712192281",
    sid: 1,
    option: 0 // 0为左边选项 1为右边选项
}
```

**Response(application/json)**
```
{
    "rtn": 0,
    "msg": "ack recieved :-)",
    "data": null
}
```

---

## 7. 用户登出

[POST] /midnightapisvr/api/session/userlogout

**Request(application/json)**
```
{
    token: "fcf41657f02f88137a1bcf068a32c0a31546766298712192281"
}
```

**Response(application/json)**
```
{
    "rtn": 0,
    "msg": "okay",
    "data": null
}
```