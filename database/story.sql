use `midnight`;

DROP TABLE IF EXISTS `story`;
DROP TABLE IF EXISTS `msg_moment`;
DROP TABLE IF EXISTS `msg_news`;
DROP TABLE IF EXISTS `msg_choose`;
DROP TABLE IF EXISTS `msg_text`;

/*对方的文字消息*/
CREATE TABLE `msg_text` (
    `mid` int NOT NULL,
    `content` varchar(128) NOT NULL,

    PRIMARY KEY (`mid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*双选消息*/
CREATE TABLE `msg_choose` (
    `mid` int NOT NULL,
    `l_content` varchar(128) NOT NULL,
    /*左边选项的 story 的 table 名*/
    `l_next_sid` int NOT NULL,
    `r_content` varchar(128) NOT NULL,
    /*左边选项的 story 的 table 名*/
    `r_next_sid` int NOT NULL,

    PRIMARY KEY (`mid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*新闻消息*/
CREATE TABLE `msg_news` (
    `mid` int NOT NULL,
    `title` varchar(32) NOT NULL,
    `content` text NOT NULL,

    PRIMARY KEY (`mid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*朋友圈消息*/
CREATE TABLE `msg_moment` (
    `mid` int NOT NULL,
    `author` varchar(32) NOT NULL,
    `content` varchar(128) NOT NULL,
    `img_uri` varchar(64) NOT NULL,

    PRIMARY KEY (`mid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `story` (
    `sid` int NOT NULL,
    /*关联消息类型：对方发言，双选，新闻等*/
    `conjunction_msg_type` int NOT NULL,
    /*关联消息 id*/
    `conjunction_mid` int NOT NULL DEFAULT -1,

    /*下一条的延迟时间 s*/
    `time_delay` int NOT NULL DEFAULT -1,
    `next_sid` int NOT NULL DEFAULT -1,

    PRIMARY KEY (`sid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
