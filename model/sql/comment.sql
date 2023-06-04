DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`
(
    `id`          bigint(20)                      NOT NULL AUTO_INCREMENT,
    `comment_id`  bigint(20) unsigned             NOT NULL,
    `content`     text COLLATE utf8mb4_general_ci NOT NULL,
    `post_id`     bigint(20)                      NOT NULL,
    `author_id`   bigint(20)                      NOT NULL,
    `parent_id`   bigint(20)                      NOT NULL DEFAULT '0',
    `status`      tinyint(3) unsigned             NOT NULL DEFAULT '1',
    `create_time` timestamp                       NULL     DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp                       NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_comment_id` (`comment_id`),
    KEY `idx_author_Id` (`author_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;