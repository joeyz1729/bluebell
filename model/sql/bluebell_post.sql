create table post
(
    id           bigint auto_increment
        primary key,
    post_id      bigint                              not null comment '帖子id',
    title        varchar(128)                        not null comment '标题',
    content      varchar(8192)                       not null comment '内容',
    author_id    bigint                              not null comment '作者的用户id',
    community_id bigint                              not null comment '所属社区',
    status       tinyint   default 1                 not null comment '帖子状态',
    create_time  timestamp default CURRENT_TIMESTAMP null comment '创建时间',
    update_time  timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    constraint idx_post_id
        unique (post_id)
)
    collate = utf8mb4_general_ci;

create index idx_author_id
    on post (author_id);

create index idx_community_id
    on post (community_id);

INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (1, 14283784123846656, '学习day1', '学习content', 28018727488323585, 1, 1, '2020-08-09 09:58:39', '2020-08-09 09:58:39');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (2, 14373128436191232, '深度优先搜索', 'dfs, bfs, 回溯, 剪枝, ...', 28018727488323585, 2, 1, '2020-08-09 15:53:40', '2020-08-09 15:53:40');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (3, 14373246019309568, '火法无敌了', '版本答案暗牧火法增辉', 28018727488323585, 3, 1, '2020-08-09 15:54:08', '2020-08-09 15:54:08');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (6, 19434165682311168, '20号赛季', '但是没新东西', 28018727488323585, 4, 1, '2020-08-23 15:04:26', '2020-08-23 15:04:26');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (9, 21810865955147776, '范型是什么', ' 如题，zsbd', 28018727488323585, 1, 1, '2020-08-30 04:28:35', '2020-08-30 04:28:35');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (10, 21810938202034176, '暑假计划', '暑假content', 28018727488323585, 1, 1, '2020-08-30 04:28:52', '2020-08-30 04:28:52');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (11, 1, 'test', 'just for test', 1, 1, 1, '2020-09-12 14:03:18', '2020-09-12 14:03:18');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (12, 92636388033302528, 'test', 'just a test', 1, 1, 1, '2020-09-12 15:03:56', '2020-09-12 15:03:56');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (13, 92636388142354432, 'test', 'just a test', 1, 1, 1, '2020-09-12 15:03:56', '2020-09-12 15:03:56');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (15, 123, 'test', 'just a test', 1, 1, 1, '2020-09-13 03:31:50', '2020-09-13 03:31:50');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (16, 10, 'test', 'just a test', 123, 1, 1, '2020-09-13 04:12:44', '2020-09-13 04:12:44');
