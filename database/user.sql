use `midnight`;

DROP TABLE IF EXISTS `user_info`;

CREATE TABLE `user_info` (
    `uid` int NOT NULL AUTO_INCREMENT,
    `username` varchar(16) NOT NULL,
    `nickname` varchar(16) NOT NULL,
    `avatar_uri` varchar(128) NOT NULL,
    `passwordhash` varchar(64) NOT NULL,

    PRIMARY KEY (`uid`),
    UNIQUE KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
