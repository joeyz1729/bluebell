create table member
(
    id           bigint auto_increment
        primary key,
    community_id      int unsigned                              not null comment '社区id',
    user_id    bigint                              not null comment '用户id',
    cancel       tinyint   default 1                 not null comment '是否被取消',
    create_time  timestamp default CURRENT_TIMESTAMP null comment '创建时间',
    update_time  timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间'
)
    collate = utf8mb4_general_ci;

create index idx_community_id
    on member (community_id);

create index idx_user_id
    on member (user_id);

INSERT INTO bluebell.member (id, community_id, user_id, cancel,create_time, update_time) VALUES (1, 1, 469896035695591425, 0, '2022-08-09 09:58:39', '2022-08-09 09:58:39');
INSERT INTO bluebell.member (id, community_id, user_id, cancel,create_time, update_time) VALUES (2, 2, 469896035695591425, 0, '2022-08-09 09:58:39', '2022-08-09 09:58:39');
INSERT INTO bluebell.member (id, community_id, user_id, cancel,create_time, update_time) VALUES (3, 3, 469896035695591425, 0, '2022-08-09 09:58:39', '2022-08-09 09:58:39');
