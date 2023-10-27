use flicka_frame;

CREATE TABLE `user` (
                        `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                        `created_at` datetime(3) DEFAULT NULL,
                        `updated_at` datetime(3) DEFAULT NULL,
                        `deleted_at` datetime(3) DEFAULT NULL,
                        `nick_name` longtext,
                        `avtar_url` longtext,
                        `gender` enum('0','1','2') COLLATE utf8mb4_bin DEFAULT NULL,
                        `age` bigint DEFAULT NULL,
                        `tik_tok_id` varchar(100) DEFAULT NULL,
                        `slogan` longtext,
                        `password` longtext,
                        `phone` varchar(100) DEFAULT NULL,
                        `following_count` bigint DEFAULT NULL,
                        `follower_count` bigint DEFAULT NULL,
                        PRIMARY KEY (`id`),
                        KEY `idx_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

INSERT INTO flicka_frame.user (id, created_at, updated_at, deleted_at, nick_name, avatar_url, gender, age, tik_tok_id, slogan, password, phone, following_count, follower_count) VALUES (1, '2023-10-26 11:12:14.000', '2023-10-26 11:12:24.000', null, '一只小熊阿萨', 'avatar/mosaic-legacy_2b1fb000159504ee55b9f~c5_300x300.jpeg', null, 22, '26748570', null, null, '12345678901', null, null);
INSERT INTO flicka_frame.user (id, created_at, updated_at, deleted_at, nick_name, avatar_url, gender, age, tik_tok_id, slogan, password, phone, following_count, follower_count) VALUES (2, '2023-10-26 11:12:15.000', '2023-10-26 11:12:26.000', null, '奶油喵', 'avatar/tos-cn-avt-0015_e8b5775c7755b93c16c8fb5dc358a5fe~c5_300x300.jpeg', null, null, '39109084', '168/52 旅游｜拍照｜穿搭 记录生活的美好日常', null, '11234567890', null, null);
INSERT INTO flicka_frame.user (id, created_at, updated_at, deleted_at, nick_name, avatar_url, gender, age, tik_tok_id, slogan, password, phone, following_count, follower_count) VALUES (3, '2023-10-26 11:12:16.000', '2023-10-26 11:12:29.000', null, '赣', 'avatar/tos-cn-i-0813_ocIALK3BYEfAmA7eHSQAB8DBoen7pSAh2booiG~c5_300x300.jpeg', null, null, '77105162676', '', null, '12345672901', null, null);
INSERT INTO flicka_frame.user (id, created_at, updated_at, deleted_at, nick_name, avatar_url, gender, age, tik_tok_id, slogan, password, phone, following_count, follower_count) VALUES (4, '2023-10-26 11:12:17.000', '2023-10-26 11:12:27.000', null, 'ZouHaoHao', 'avatar/tos-cn-avt-0015_c6b9be7c006f7d613e0b6f44f1006fc1~c5_300x300.jpeg', null, null, 'ZouHaoHao', '无任何海外社交平台 注意辨别 切勿被骗 合：Hao94499（备注品牌）
5级粉丝牌进群加：a46190236', null, '12345678931', null, null);
INSERT INTO flicka_frame.user (id, created_at, updated_at, deleted_at, nick_name, avatar_url, gender, age, tik_tok_id, slogan, password, phone, following_count, follower_count) VALUES (5, '2023-10-26 11:12:18.000', '2023-10-26 11:12:30.000', null, 'soso 11', 'avatar/tos-cn-avt-0015_79eb2b3a555bd1baf1d301d6151a5afa~c5_300x300.jpeg', null, null, 'soso77609', '看世界也找自己', null, '12345478901', null, null);
INSERT INTO flicka_frame.user (id, created_at, updated_at, deleted_at, nick_name, avatar_url, gender, age, tik_tok_id, slogan, password, phone, following_count, follower_count) VALUES (6, '2023-10-26 11:12:20.000', '2023-10-26 11:12:31.000', null, '菠萝头娱乐', 'avatar/tos-cn-avt-0015_96aec97f6e1cb6485ed376c7845331bb~c5_300x300.jpeg', null, null, 'woshiboluotou', '一只爱冲浪刷剧追星吃瓜的打工菠萝头。 ?商务合作洽谈? ?...', null, '12345678201', null, null);
INSERT INTO flicka_frame.user (id, created_at, updated_at, deleted_at, nick_name, avatar_url, gender, age, tik_tok_id, slogan, password, phone, following_count, follower_count) VALUES (7, '2023-10-26 11:12:21.000', '2023-10-26 11:12:32.000', null, '一个板栗栗', 'avatar/tos-cn-avt-0015_85f786622ae61a031342cad02dc2b3e5~c5_300x300.jpeg', '1', 21, '49187982260', '留在我身边', null, '12345678921', null, null);
INSERT INTO flicka_frame.user (id, created_at, updated_at, deleted_at, nick_name, avatar_url, gender, age, tik_tok_id, slogan, password, phone, following_count, follower_count) VALUES (8, '2023-10-26 11:12:22.000', '2023-10-26 11:12:33.000', null, '李一桐Q', 'avatar/tos-cn-avt-0015_97cef8728febb4c5509662b2f9c7ee7d~c5_300x300.jpeg', null, null, 'liyitong0906', '', null, '12645678901', null, null);
INSERT INTO flicka_frame.user (id, created_at, updated_at, deleted_at, nick_name, avatar_url, gender, age, tik_tok_id, slogan, password, phone, following_count, follower_count) VALUES (9, '2023-10-26 11:12:23.000', '2023-10-26 11:12:34.000', null, '星球音乐', 'avatar/tos-cn-avt-0015_a72474791c379a819bb12328ccc2552f~c5_300x300.jpeg', null, null, 'xryyzn512', '', null, '15345678901', null, null);


