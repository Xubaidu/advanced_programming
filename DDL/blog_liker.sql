-- auto-generated definition
create table blog_liker
(
    updated_at timestamp default CURRENT_TIMESTAMP not null comment '更新日期',
    deleted_at timestamp default CURRENT_TIMESTAMP null comment '注销日期',
    created_at timestamp default CURRENT_TIMESTAMP not null comment '注册日期',
    user_id    int unsigned                        not null comment '点赞人 user id',
    blog_id    int unsigned                        not null comment '帖子所属的 blog id',
    id         int unsigned auto_increment comment '主键，没有实际业务含义'
        primary key,
    constraint blog_liker_ibfk_1
        foreign key (blog_id) references blog (id),
    constraint blog_liker_ibfk_2
        foreign key (user_id) references user (id)
);

create index blog_id
    on blog_liker (blog_id);

create index user_id
    on blog_liker (user_id);

