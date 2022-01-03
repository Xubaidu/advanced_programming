-- auto-generated definition
create table blog
(
    id         int unsigned auto_increment comment '帖子 id'
        primary key,
    title      varchar(100)                           not null comment '帖子标题',
    created_at timestamp    default CURRENT_TIMESTAMP null,
    deleted_at timestamp                              null,
    updated_at timestamp    default CURRENT_TIMESTAMP null,
    content    varchar(1000)                          not null comment '正文',
    user_id    int unsigned                           not null,
    likes      int unsigned default '0'               not null,
    constraint blog_ibfk_1
        foreign key (user_id) references user (id)
);

create index author_id
    on blog (user_id);

