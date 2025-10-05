/*
 Navicat Premium Data Transfer

 Source Server         : 本地MYSQL8.0
 Source Server Type    : MySQL
 Source Server Version : 80037
 Source Host           : localhost:3307
 Source Schema         : task4

 Target Server Type    : MySQL
 Target Server Version : 80037
 File Encoding         : 65001

 Date: 05/10/2025 15:21:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `post_id` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_comments_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_comments_user`(`user_id` ASC) USING BTREE,
  INDEX `fk_comments_post`(`post_id` ASC) USING BTREE,
  CONSTRAINT `fk_comments_post` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_comments_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES (1, '2025-10-05 15:12:38.909', '2025-10-05 15:12:38.909', NULL, '文章1的评论', 1, 1);
INSERT INTO `comments` VALUES (2, '2025-10-05 15:12:53.774', '2025-10-05 15:12:53.774', NULL, '文章2的评论', 1, 2);
INSERT INTO `comments` VALUES (3, '2025-10-05 15:13:01.516', '2025-10-05 15:13:01.516', NULL, '文章1的评论2', 1, 1);
INSERT INTO `comments` VALUES (4, '2025-10-05 15:14:49.797', '2025-10-05 15:14:49.797', NULL, '文章1的评论3', 1, 1);

-- ----------------------------
-- Table structure for posts
-- ----------------------------
DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_posts_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_posts_user`(`user_id` ASC) USING BTREE,
  CONSTRAINT `fk_posts_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of posts
-- ----------------------------
INSERT INTO `posts` VALUES (1, '2025-10-05 14:34:40.888', '2025-10-05 14:34:40.888', NULL, '这是标题1', '标题1文章的内容11111111', 1);
INSERT INTO `posts` VALUES (2, '2025-10-05 14:37:02.698', '2025-10-05 14:37:02.698', NULL, '这是标题2', '标题2文章的内容11111111', 1);
INSERT INTO `posts` VALUES (3, '2025-10-05 14:37:08.438', '2025-10-05 14:37:08.438', NULL, '这是标题3', '标题3文章的内容11111111', 1);
INSERT INTO `posts` VALUES (4, '2025-10-05 14:37:13.965', '2025-10-05 15:00:14.436', '2025-10-05 15:04:18.681', '这是标题44444', '标题4文章的内容11111111', 1);
INSERT INTO `posts` VALUES (5, '2025-10-05 14:37:28.787', '2025-10-05 14:37:28.787', NULL, '这是标题5', '标题5文章的内容11111111', 3);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_users_username`(`username` ASC) USING BTREE,
  UNIQUE INDEX `uni_users_email`(`email` ASC) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '2025-10-05 14:15:25.069', '2025-10-05 14:15:25.069', NULL, '张杰', '$2a$10$WiESaagDxxO78lEPw.psuuu8y1wokqhaBmudI/UozjmkWLoi0ZyVC', '1448282786@qq.com');
INSERT INTO `users` VALUES (3, '2025-10-05 14:59:07.907', '2025-10-05 14:59:07.907', NULL, '谢娜', '$2a$10$MaWtThSPeBWgO3tGdR7ex.O3Rm4i2D90VKljMMNosCBsIW5NK6QoK', '1448282787@qq.com');

SET FOREIGN_KEY_CHECKS = 1;
