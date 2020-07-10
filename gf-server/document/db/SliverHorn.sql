-- admins Model
create table admins
(
    id           int unsigned auto_increment comment '自增ID'
        primary key,
    create_at    timestamp                                                       null comment '创建时间',
    update_at    timestamp                                                       null comment '更新时间',
    delete_at    timestamp                                                       null comment '删除时间',
    uuid         varbinary(255)                                                  null comment '用户唯一标识UUID',
    nickname     varchar(255) default 'QMPlusUser'                               null comment '用户昵称',
    header_img   varchar(255) default 'http://www.henrongyi.top/avatar/lufu.jpg' null comment '用户头像',
    authority_id double       default 888                                        null comment '用户角色ID',
    username     varchar(255)                                                    null comment '用户名',
    password     varchar(255)                                                    null comment '用户登录密码'
)
    charset = utf8;

create index idx_admins_deleted_at
    on admins (delete_at);

