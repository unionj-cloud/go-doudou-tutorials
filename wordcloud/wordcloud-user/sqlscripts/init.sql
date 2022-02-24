CREATE SCHEMA `tutorial` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

CREATE TABLE `tutorial`.`t_user`
(
    `id`        INT         NOT NULL AUTO_INCREMENT,
    `username`  VARCHAR(45) NOT NULL COMMENT 'username',
    `password`  VARCHAR(60) NOT NULL COMMENT 'password',
    `name`      VARCHAR(45) NOT NULL COMMENT 'real name',
    `phone`     VARCHAR(45) NOT NULL COMMENT 'phone number',
    `dept`      VARCHAR(45) NOT NULL COMMENT 'department',
    `create_at` DATETIME    NULL DEFAULT current_timestamp,
    `update_at` DATETIME    NULL DEFAULT current_timestamp on update current_timestamp,
    `delete_at` DATETIME    NULL,
    PRIMARY KEY (`id`)
);