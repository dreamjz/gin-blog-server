-- 若不存在则创建数据库 gin-blog ，字符集 utf8 ，校对规则 utf8_general_ci 不区分大小写
CREATE DATABASE IF NOT EXISTS `gin_blog` CHARSET utf8 COLLATE utf8_general_ci;

-- Create Table For gin_blog
USE `gin_blog`;

-- User Auth Table
CREATE TABLE IF NOT EXISTS `blog_auth` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(100) DEFAULT '' COMMENT 'Username',
  `password` VARCHAR(100) DEFAULT '' COMMENT 'Password',
  `created_by` VARCHAR(100) DEFAULT '' COMMENT 'Username created by',
  `updated_by` VARCHAR(100) DEFAULT '' COMMENT 'Username updated by',
  `created_at` datetime COMMENT 'Created time',
  `updated_at` datetime COMMENT 'updated time',
  `deleted_at` datetime COMMENT 'deleted time',
  PRIMARY KEY (id),
  UNIQUE(username)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'User auth table';

-- Create user admin
INSERT IGNORE INTO `blog_auth` (`username`,`password`,`created_at`) VALUES ('admin','admin',NOW());

-- Article Table
CREATE TABLE IF NOT EXISTS `blog_article` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `author` VARCHAR(100) NOT NULL COMMENT 'author',
    `title` VARCHAR(120) NOT NULL COMMENT 'title',
    `summary` VARCHAR(120) COMMENT 'summary',
    `content` TEXT NOT NULL COMMENT 'article content',
    `importance` TINYINT DEFAULT 0 COMMENT 'importance',
    `status` TINYINT NOT NULL COMMENT 'status 0 draft 1 published',
    `created_by` VARCHAR(100) DEFAULT '' COMMENT 'Article created by',
    `updated_by` VARCHAR(100) DEFAULT '' COMMENT 'Article updated by',
    `created_at` datetime COMMENT 'Created time',
    `updated_at` datetime COMMENT 'updated time',
    `deleted_at` datetime COMMENT 'deleted time',
    PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'Article table';

