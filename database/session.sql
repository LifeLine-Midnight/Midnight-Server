use `midnight`;

DROP TABLE IF EXISTS `user_session`;

CREATE TABLE `user_session` (
    `sid` int NOT NULL AUTO_INCREMENT,
    `stoken` varchar(64) NOT NULL,
    `uid` int NOT NULL,

    PRIMARY KEY (`sid`),
    UNIQUE KEY (`stoken`),
    UNIQUE KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
