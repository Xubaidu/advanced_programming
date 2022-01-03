-- auto-generated definition
create table comment
(
    id         int unsigned auto_increment comment '主键，没有实际业务含义'
        primary key,
    blog_id    int unsigned                        not null comment '帖子所属的 blog id',
    created_at timestamp default CURRENT_TIMESTAMP null,
    deleted_at timestamp                           null,
    updated_at timestamp default CURRENT_TIMESTAMP null,
    author_id  int unsigned                        not null comment '评论人 user id',
    constraint comment_ibfk_1
        foreign key (blog_id) references blog (id),
    constraint comment_ibfk_2
        foreign key (author_id) references user (id)
);

create index blog_id
    on comment (blog_id);

create index user_id
    on comment (author_id);

