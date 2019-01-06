use `midnight`;

DROP TABLE IF EXISTS `user_process`;

CREATE TABLE `user_process` (
    `uid` int NOT NULL,
    `cur_sid` int NOT NULL DEFAULT 1,
    `cur_process_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
