/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80034
 Source Host           : localhost:3306
 Source Schema         : collabsearch

 Target Server Type    : MySQL
 Target Server Version : 80034
 File Encoding         : 65001

 Date: 07/09/2023 13:06:13
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `uid` bigint(64) UNSIGNED ZEROFILL NOT NULL,
  `email` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `username` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `profile` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT 'Introduce yourself to others.',
  `avatar` mediumtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  PRIMARY KEY (`uid`) USING BTREE,
  UNIQUE INDEX `uid_UNIQUE`(`uid` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = 'user information' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_workspace
-- ----------------------------
DROP TABLE IF EXISTS `user_workspace`;
CREATE TABLE `user_workspace`  (
  `uid` bigint(64) UNSIGNED ZEROFILL NOT NULL,
  `wid` int NOT NULL,
  PRIMARY KEY (`uid`, `wid`) USING BTREE,
  INDEX `wid`(`wid` ASC) USING BTREE,
  CONSTRAINT `user_workspace_ibfk_2` FOREIGN KEY (`uid`) REFERENCES `user` (`uid`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `user_workspace_ibfk_3` FOREIGN KEY (`wid`) REFERENCES `workspace` (`wid`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for workspace
-- ----------------------------
DROP TABLE IF EXISTS `workspace`;
CREATE TABLE `workspace`  (
  `wid` int NOT NULL AUTO_INCREMENT,
  `topic` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `description` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `handler` bigint(64) UNSIGNED ZEROFILL NOT NULL,
  `due` datetime NOT NULL DEFAULT '9999-12-31 23:59:59',
  PRIMARY KEY (`wid`) USING BTREE,
  UNIQUE INDEX `wid_UNIQUE`(`wid` ASC) USING BTREE,
  INDEX `handler`(`handler` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = 'workspace status, informations and relative entity' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
