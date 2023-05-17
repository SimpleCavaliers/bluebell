CREATE TABLE `user`(
                       `id` bigint(20) not null auto_increment,
                       `user_id` bigint(20) not null,
                       `username` varchar(64) collate utf8mb4_general_ci not null,
                       `password` varchar(64) collate utf8mb4_general_ci not null,
                       `email` varchar(64) collate utf8mb4_general_ci,
                       `gender` tinyint(4) not null default '0',
                       `create_time` timestamp null default CURRENT_TIMESTAMP,
                       `update_time` timestamp null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
                       primary key (`id`),
                       unique key `idx_username` (`username`) USING BTREE,
                       unique key `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


CREATE TABLE `community`(
                       `id` int(11) not null auto_increment,
                       `community_id` int(10) unsigned not null,
                       `community_name` varchar(128) collate utf8mb4_general_ci not null,
                       `introduction` varchar(256) collate utf8mb4_general_ci not null,
                       `create_time` timestamp null default CURRENT_TIMESTAMP,
                       `update_time` timestamp null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
                       primary key (`id`),
                       unique key `idx_username` (`community_id`),
                       unique key `idx_user_id` (`community_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `community` values ('1', '1', 'Go', 'Golang', '2016-11-01 08:10:10', '2016-11-01 08:10:10');
INSERT INTO `community` values ('2', '2', 'leetcode', '刷就硬刷', '2020-01-01 08:00:10', '2020-01-01 08:00:10');
INSERT INTO `community` values ('3', '3', 'CS:Go', 'Rush B', '2018-07-01 08:30:00', '2018-07-01 08:30:00');
INSERT INTO `community` values ('4', '4', 'LOL', '大家都来打飞机', '2016-01-01 08:00:00', '2016-01-01 08:00:00');

create table `post`(
    `id` bigint(20) not null auto_increment,
    `post_id` bigint(20) not null comment '帖子id',
    `title` varchar(128) collate utf8mb4_general_ci not null comment '标题',
    `content` varchar(128) collate utf8mb4_general_ci not null comment '内容',
    `author_id` bigint(20) not null,
    `community_id` bigint(20) not null,
    `status` tinyint(4) not null default '1',
    `create_time` timestamp null default CURRENT_TIMESTAMP,
    `update_time` timestamp null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    primary key (`id`),
    unique key `idx_post_id` (`post_id`),
    key `idx_authod_id` (`author_id`),
    key `idx_community_id` (`community_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

