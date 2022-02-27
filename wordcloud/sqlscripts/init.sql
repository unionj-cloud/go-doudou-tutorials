CREATE SCHEMA `tutorial` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

create table tutorial.t_invalid_token
(
    id        int auto_increment
        primary key,
    token     varchar(120)                       not null comment 'revoked token',
    expire_at datetime                           null comment 'expire time',
    create_at datetime default CURRENT_TIMESTAMP null,
    update_at datetime default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    delete_at datetime                           null,
    constraint token_idx
        unique (token)
);

create table tutorial.t_user
(
    id        int auto_increment
        primary key,
    username  varchar(45)                        not null comment 'username',
    password  varchar(60)                        not null comment 'password',
    name      varchar(45)                        not null comment 'real name',
    phone     varchar(45)                        not null comment 'phone number',
    dept      varchar(45)                        not null comment 'department',
    create_at datetime default CURRENT_TIMESTAMP null,
    update_at datetime default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    delete_at datetime                           null,
    avatar    varchar(255)                       not null comment 'user avatar',
    constraint username_idx
        unique (username)
);

create table tutorial.t_word_cloud_task
(
    id        int auto_increment
        primary key,
    src_url   varchar(255)                       not null comment 'source file url 原文件链接',
    img_url   varchar(255)                       not null comment 'word cloud image url 词云图链接',
    create_at datetime default CURRENT_TIMESTAMP null,
    update_at datetime default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    delete_at datetime                           null,
    status    tinyint                            not null comment 'task status 任务状态, 0 waiting 未开始, 1 processing 进行中, 2 success 成功, 3 fail 失败',
    error     varchar(655)                       not null comment 'error message 错误信息',
    user_id   int                                not null,
    lang      varchar(45)                        not null comment 'text language 文本语种, zh中文, en英文',
    top       int                                not null comment 'take top frequent words into word cloud image 取词频前top的词 default is 0, mean no limit 默认值为0，表示不限制',
    constraint fk_user
        foreign key (user_id) references tutorial.t_user (id)
            on delete cascade
);
