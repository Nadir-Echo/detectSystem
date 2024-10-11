package models

import (
	"fmt"
	"testing"
	"time"
)

/*
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for prod
-- ----------------------------
DROP TABLE IF EXISTS `prod`;
CREATE TABLE `prod`  (
  `pid` int NOT NULL AUTO_INCREMENT COMMENT '检测产品唯一标识',
  `ori_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '原图片',
  `det_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '检测图',
  `pro_id` int NULL DEFAULT NULL COMMENT '检测对应的生产线',
  `err_result` int NULL DEFAULT NULL COMMENT '0 是无缺陷，1是有缺陷',
  `cre_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`pid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for result
-- ----------------------------
DROP TABLE IF EXISTS `result`;
CREATE TABLE `result`  (
  `defect_id` int NOT NULL AUTO_INCREMENT COMMENT '缺陷编号',
  `defect_result` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '检测缺陷种类',
  `pid` int NULL DEFAULT NULL COMMENT '对应的产品编号',
  `conf` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '置信度',
  `position` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '缺陷坐标',
  PRIMARY KEY (`defect_id`) USING BTREE,
  INDEX `pid`(`pid` ASC) USING BTREE,
  CONSTRAINT `pid` FOREIGN KEY (`pid`) REFERENCES `prod` (`pid`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `user_id` int NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `user_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户名',
  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT ' 密码',
  `pro_id` int NULL DEFAULT NULL COMMENT '负责生产线的编号',
  `state` int NULL DEFAULT NULL COMMENT ' 用户状态（0：启用，1：禁用',
  `role` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT ' 用户类型（用于用户分组，比如管理员，生产线负责人）',
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '电话号码',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '真实姓名',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;


*/

// test AddPro function
func TestAddPro(t *testing.T) {

	//cteate a new prod called data
	var data Prod
	data.OriUrl = "https://joplin-1305917469.cos.ap-nanjing.myqcloud.com/Snipaste_2023-04-17_16-23-26.png"
	data.DetUrl = ""
	data.ProId = 1
	data.ErrResult = 0
	data.CreTime = time.Now()
	// data.Pid = 0

	proId, err := AddProd(&data)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(proId)
}
// Test GetProResult function
func TestGetProResult(t *testing.T) {
  data,err := GetProResult(1)
  if err != nil {
    t.Error(err)
  }
  fmt.Println(data[1])
}

// Test GetALlProResult function
func TestGetAllProResult(t *testing.T) {
  data,err := GetAllProResult(0)
  if err != nil {
    t.Error(err)
  }
  fmt.Println(data)
}
