create table follow
(
    id           bigint auto_increment
        primary key,

    user_id    bigint                              not null comment '用户id',
    follower_id bigint                              not null comment 'follower if',
    cancel       tinyint   default 1                 not null comment '是否被取消',
    create_time  timestamp default CURRENT_TIMESTAMP null comment '创建时间',
    update_time  timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间'
)
    collate = utf8mb4_general_ci;

create index idx_user_id
    on follow (user_id);

create index idx_follower_id
    on follow (follower_id);

INSERT INTO bluebell.follow (id, follower_id, user_id, cancel,create_time, update_time) VALUES (1, 1, 471876904828272641
                                                                                               , 0, '2022-08-09 09:58:39', '2022-08-09 09:58:39');
INSERT INTO bluebell.follow (id, follower_id, user_id, cancel,create_time, update_time) VALUES (2, 2, 471876904828272641
                                                                                               , 0, '2022-08-09 09:58:39', '2022-08-09 09:58:39');                                                                                               , 0, '2022-08-09 09:58:39', '2022-08-09 09:58:39');
