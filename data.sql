/*创建数据库*/
CREATE DATABASE IF NOT EXISTS `video_cms`;
USE `video_cms`;


/*用户表*/
CREATE TABLE IF NOT EXISTS `users`(
    `id` INT UNSIGNED  AUTO_INCREMENT,
    `login_name` VARCHAR(64) UNIQUE KEY,
    `pwd` TEXT,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `video_info`(
    `id` VARCHAR(64) NOT NULL,
    `author_id` INT UNSIGNED,
    `name` TEXT,
    `display_ctime` TEXT,
    `create_time` DATETIME DEFAULT now(),
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `comments`(
    `id` VARCHAR(64) NOT NULL,
    `video_id` VARCHAR(64),
    `author_id` INT UNSIGNED,
    `content` TEXT,
    `time` DATETIME DEFAULT now(),
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `sessions`(
    `session_id` VARCHAR(255) NOT NULL,
    `TTL` TINYTEXT,/*过期时间*/
    `login_name` VARCHAR(64),
    PRIMARY KEY (`session_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;