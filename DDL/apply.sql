-- auto-generated definition
create table apply
(
    updated_at timestamp default CURRENT_TIMESTAMP not null comment '更新日期',
    deleted_at timestamp default CURRENT_TIMESTAMP null comment '注销日期',
    created_at timestamp default CURRENT_TIMESTAMP not null comment '注册日期',
    user_id    int unsigned                        not null comment '申请人',
    job_id     int unsigned                        not null comment '岗位',
    id         int unsigned auto_increment comment '主键，没有实际业务含义'
        primary key,
    constraint apply_ibfk_1
        foreign key (job_id) references job (id),
    constraint apply_ibfk_2
        foreign key (user_id) references user (id)
);

create index job_id
    on apply (job_id);

create index user_id
    on apply (user_id);

