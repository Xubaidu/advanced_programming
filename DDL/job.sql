-- auto-generated definition
create table job
(
    updated_at timestamp default CURRENT_TIMESTAMP not null comment '更新日期',
    deleted_at timestamp default CURRENT_TIMESTAMP null comment '注销日期',
    created_at timestamp default CURRENT_TIMESTAMP not null comment '注册日期',
    user_id    int unsigned                        not null,
    applicants int unsigned                        not null comment '申请人数',
    content    varchar(1000)                       null,
    base       varchar(100)                        not null comment '岗位所在地',
    title      varchar(100)                        not null comment '岗位名称',
    company    varchar(100)                        not null comment '公司名',
    id         int unsigned auto_increment
        primary key,
    constraint job_ibfk_1
        foreign key (user_id) references user (id)
);

create index author_id
    on job (user_id);

