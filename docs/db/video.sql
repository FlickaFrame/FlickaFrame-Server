use flicka_frame;

CREATE TABLE video (
                        id INT AUTO_INCREMENT PRIMARY KEY,
                        title VARCHAR(255),
                        play_url VARCHAR(255),
                        cover_url VARCHAR(255),
                        favorite_count INT default 0,
                        comment_count INT default 0,
                        author_id INT,
                        created_at TIMESTAMP,
                        updated_at TIMESTAMP,
                        index idx_authorid(author_id)
                        );

INSERT INTO flicka_frame.video (id, created_at, updated_at, deleted_at, title, e_tag, play_url, cover_url, favorite_count, comment_count, author_id, category_id) VALUES (3, '2023-10-26 11:00:52.000', '2023-10-26 11:01:00.000', null, null, null, 'video/8eb165b7216487286363d264063f6ea5_raw.mp4', null, null, null, 1, null);
INSERT INTO flicka_frame.video (id, created_at, updated_at, deleted_at, title, e_tag, play_url, cover_url, favorite_count, comment_count, author_id, category_id) VALUES (4, '2023-10-26 11:00:54.000', '2023-10-26 11:01:01.000', null, null, null, 'video/915a67b735b5bd4013be31c31674a92c_raw.mp4', null, null, null, 2, null);
INSERT INTO flicka_frame.video (id, created_at, updated_at, deleted_at, title, e_tag, play_url, cover_url, favorite_count, comment_count, author_id, category_id) VALUES (5, '2023-10-26 11:00:56.000', '2023-10-26 11:01:30.000', null, null, null, 'video/7ddb68bf06913626392c373354462894_raw.mp4', null, null, null, 2, null);
INSERT INTO flicka_frame.video (id, created_at, updated_at, deleted_at, title, e_tag, play_url, cover_url, favorite_count, comment_count, author_id, category_id) VALUES (6, '2023-10-26 11:00:57.000', '2023-10-26 11:01:32.000', null, null, null, 'video/07554c4e2991a8b0a5c7093263ed72cd_raw.mp4', null, null, null, 7, null);
INSERT INTO flicka_frame.video (id, created_at, updated_at, deleted_at, title, e_tag, play_url, cover_url, favorite_count, comment_count, author_id, category_id) VALUES (7, '2023-10-26 11:00:58.000', '2023-10-26 11:01:33.000', null, null, null, 'video/29cb3cd69ef19ed33549806d8f6b96b6_raw.mp4', null, null, null, 7, null);
INSERT INTO flicka_frame.video (id, created_at, updated_at, deleted_at, title, e_tag, play_url, cover_url, favorite_count, comment_count, author_id, category_id) VALUES (8, '2023-10-26 11:01:06.000', '2023-10-26 11:01:33.000', null, null, null, 'video/53e51db6dd95692de07db424412db009_raw.mp4', null, null, null, 7, null);
INSERT INTO flicka_frame.video (id, created_at, updated_at, deleted_at, title, e_tag, play_url, cover_url, favorite_count, comment_count, author_id, category_id) VALUES (9, '2023-10-26 11:01:06.000', '2023-10-26 11:01:36.000', null, null, null, 'video/67dcc76e39da126aee0950bace0f8d0e_raw.mp4', null, null, null, 7, null);
INSERT INTO flicka_frame.video (id, created_at, updated_at, deleted_at, title, e_tag, play_url, cover_url, favorite_count, comment_count, author_id, category_id) VALUES (10, '2023-10-26 11:01:08.000', '2023-10-26 11:01:37.000', null, null, null, 'video/9890e40bf432fbf3d341fca0f7514b1e_raw.mp4', null, null, null, 7, null);
INSERT INTO flicka_frame.video (id, created_at, updated_at, deleted_at, title, e_tag, play_url, cover_url, favorite_count, comment_count, author_id, category_id) VALUES (11, '2023-10-26 11:01:09.000', '2023-10-26 11:01:35.000', null, null, null, 'video/c129ad5fbd1595c46cdd4c5f4aadf161_raw.mp4', null, null, null, 3, null);
INSERT INTO flicka_frame.video (id, created_at, updated_at, deleted_at, title, e_tag, play_url, cover_url, favorite_count, comment_count, author_id, category_id) VALUES (12, '2023-10-26 11:01:05.000', '2023-10-26 11:01:39.000', null, null, null, 'video/a41dab29078223f12c03a5b81f76d4d1_raw.mp4', null, null, null, 4, null);
INSERT INTO flicka_frame.video (id, created_at, updated_at, deleted_at, title, e_tag, play_url, cover_url, favorite_count, comment_count, author_id, category_id) VALUES (13, '2023-10-26 11:01:10.000', '2023-10-26 11:01:40.000', null, null, null, 'video/3bae3f1a5b0a3e0512c04f5b31b1338e_raw.mp4', null, null, null, 8, null);
INSERT INTO flicka_frame.video (id, created_at, updated_at, deleted_at, title, e_tag, play_url, cover_url, favorite_count, comment_count, author_id, category_id) VALUES (14, '2023-10-26 11:01:13.000', '2023-10-26 11:01:41.000', null, null, null, 'video/6705eece7d6130901a6154be0b78b234_raw.mp4', null, null, null, 8, null);
INSERT INTO flicka_frame.video (id, created_at, updated_at, deleted_at, title, e_tag, play_url, cover_url, favorite_count, comment_count, author_id, category_id) VALUES (15, '2023-10-26 11:01:11.000', '2023-10-26 11:01:38.000', null, null, null, 'video/6d6c521d39d75d11122be3487982e20c_raw.mp4', null, null, null, 8, null);
INSERT INTO flicka_frame.video (id, created_at, updated_at, deleted_at, title, e_tag, play_url, cover_url, favorite_count, comment_count, author_id, category_id) VALUES (16, '2023-10-26 11:01:14.000', '2023-10-26 11:01:43.000', null, null, null, 'video/91a149f39295982d8a158762858aa7dd_raw.mp4', null, null, null, 8, null);
INSERT INTO flicka_frame.video (id, created_at, updated_at, deleted_at, title, e_tag, play_url, cover_url, favorite_count, comment_count, author_id, category_id) VALUES (17, '2023-10-26 11:01:15.000', '2023-10-26 11:01:45.000', null, null, null, 'video/db5f297bf3985247499da0f7c6153c6f_raw.mp4', null, null, null, 5, null);
INSERT INTO flicka_frame.video (id, created_at, updated_at, deleted_at, title, e_tag, play_url, cover_url, favorite_count, comment_count, author_id, category_id) VALUES (18, '2023-10-26 11:01:17.000', '2023-10-26 11:01:44.000', null, null, null, 'video/74c4d0a8a73e527593246b48a4df8e02_raw.mp4', null, null, null, 6, null);
INSERT INTO flicka_frame.video (id, created_at, updated_at, deleted_at, title, e_tag, play_url, cover_url, favorite_count, comment_count, author_id, category_id) VALUES (19, '2023-10-26 11:01:16.000', '2023-10-26 11:01:47.000', null, null, null, 'video/c3e24d2691db7dd030b7a65b5666ef3a_raw.mp4', null, null, null, 6, null);
INSERT INTO flicka_frame.video (id, created_at, updated_at, deleted_at, title, e_tag, play_url, cover_url, favorite_count, comment_count, author_id, category_id) VALUES (20, '2023-10-26 11:01:04.000', '2023-10-26 11:01:48.000', null, null, null, 'video/35d75c989c5c5eba47b9f45d37a660e6_raw.mp4', null, null, null, 9, null);
INSERT INTO flicka_frame.video (id, created_at, updated_at, deleted_at, title, e_tag, play_url, cover_url, favorite_count, comment_count, author_id, category_id) VALUES (21, '2023-10-26 11:01:19.000', '2023-10-26 11:01:50.000', null, null, null, 'video/9dd904c497717165f034a5c2bc46aecc_raw.mp4', null, null, null, 4, null);
