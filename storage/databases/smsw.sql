/*
 Navicat Premium Data Transfer

 Source Server         : 本地虚拟机ub20
 Source Server Type    : MySQL
 Source Server Version : 80031 (8.0.31)
 Source Host           : 192.168.1.13:3306
 Source Schema         : smsw

 Target Server Type    : MySQL
 Target Server Version : 80031 (8.0.31)
 File Encoding         : 65001

 Date: 20/04/2023 01:35:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for autotask
-- ----------------------------
DROP TABLE IF EXISTS `autotask`;
CREATE TABLE `autotask`  (
                             `id` int NOT NULL AUTO_INCREMENT,
                             `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '自动化名称',
                             `dotime` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '执行时间',
                             `deviceId` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '关联的设备id',
                             `deviceAc` int NULL DEFAULT NULL COMMENT '设备动作',
                             `show` int NULL DEFAULT NULL COMMENT '是否打开',
                             `createdAt` datetime NULL DEFAULT NULL COMMENT '创建时间',
                             `updatedAt` datetime NULL DEFAULT NULL COMMENT '更新时间',
                             PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of autotask
-- ----------------------------
INSERT INTO `autotask` VALUES (1, '自动打开照明', '15:54', '103', 1, 1, '2023-04-19 15:54:53', '2023-04-19 15:54:55');
INSERT INTO `autotask` VALUES (2, '自动打开消毒', '15:55', '105', 1, 1, '2023-04-19 16:19:21', '2023-04-19 15:55:20');
INSERT INTO `autotask` VALUES (14, 'aaaa', '01:34', '105', 1, 1, '2023-04-19 22:50:56', '2023-04-19 22:50:56');

-- ----------------------------
-- Table structure for deviceconfig
-- ----------------------------
DROP TABLE IF EXISTS `deviceconfig`;
CREATE TABLE `deviceconfig`  (
                                 `id` bigint NOT NULL AUTO_INCREMENT,
                                 `deviceId` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                                 `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                                 `iconUrl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                                 `status` int NULL DEFAULT NULL,
                                 `createdAt` datetime NULL DEFAULT NULL,
                                 `updatedAt` datetime NULL DEFAULT NULL,
                                 PRIMARY KEY (`id`) USING BTREE,
                                 UNIQUE INDEX `id`(`id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of deviceconfig
-- ----------------------------
INSERT INTO `deviceconfig` VALUES (1, '103', '照明', 'https://smartwardrobe.oss-cn-shenzhen.aliyuncs.com/icon/zhaoming.png', 0, '2023-04-19 17:27:33', '2023-04-02 17:27:35');
INSERT INTO `deviceconfig` VALUES (2, '105', '消毒', 'https://smartwardrobe.oss-cn-shenzhen.aliyuncs.com/icon/zhaoming.png', 0, '2023-04-02 17:27:46', '2023-04-02 17:27:48');
INSERT INTO `deviceconfig` VALUES (3, '106', '门', 'https://smartwardrobe.oss-cn-shenzhen.aliyuncs.com/icon/zhaoming.png', 0, '2023-04-02 17:28:05', '2023-04-02 17:28:07');
INSERT INTO `deviceconfig` VALUES (4, '107', '除湿', 'https://smartwardrobe.oss-cn-shenzhen.aliyuncs.com/icon/zhaoming.png', 0, '2023-04-16 17:28:35', '2023-04-16 17:28:37');

-- ----------------------------
-- Table structure for devicevalue
-- ----------------------------
DROP TABLE IF EXISTS `devicevalue`;
CREATE TABLE `devicevalue`  (
                                `id` bigint NOT NULL AUTO_INCREMENT,
                                `deviceId` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                                `value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                                `unit` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                                `createdAt` datetime NULL DEFAULT NULL,
                                `updatedAt` datetime NULL DEFAULT NULL,
                                PRIMARY KEY (`id`) USING BTREE,
                                UNIQUE INDEX `id`(`id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 32 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of devicevalue
-- ----------------------------
INSERT INTO `devicevalue` VALUES (21, '101', '28', '℃', '2023-04-02 16:55:29', '2023-04-16 16:55:29');
INSERT INTO `devicevalue` VALUES (22, '102', '68', '%', '2023-04-02 16:55:29', '2023-04-16 16:55:29');
INSERT INTO `devicevalue` VALUES (23, '101', '27', '℃', '2023-04-02 17:13:05', '2023-04-16 17:13:05');
INSERT INTO `devicevalue` VALUES (24, '102', '70', '%', '2023-04-16 17:13:06', '2023-04-16 17:13:06');
INSERT INTO `devicevalue` VALUES (25, '102', '69', '%', '2023-04-20 00:23:37', '2023-04-20 00:23:40');
INSERT INTO `devicevalue` VALUES (26, '102', '68', '%', '2023-04-20 00:23:53', '2023-04-20 00:23:55');
INSERT INTO `devicevalue` VALUES (27, '102', '67', '%', '2023-04-20 00:24:11', '2023-04-20 00:24:13');
INSERT INTO `devicevalue` VALUES (28, '102', '66', '%', '2023-04-20 00:24:31', '2023-04-20 00:24:34');
INSERT INTO `devicevalue` VALUES (29, '102', '65', '%', '2023-04-20 00:25:04', '2023-04-20 00:25:06');
INSERT INTO `devicevalue` VALUES (30, '101', '26', '℃', '2023-04-20 00:45:25', '2023-04-20 00:45:27');
INSERT INTO `devicevalue` VALUES (31, '101', '29', '℃', '2023-04-20 00:45:38', '2023-04-20 00:45:41');

SET FOREIGN_KEY_CHECKS = 1;
