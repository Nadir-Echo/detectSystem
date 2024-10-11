package models

import (
	"fmt"
	"log"
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
//struct定义
type User struct {
	UserID   int    `gorm:"AUTO_INCREMENT" json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	ProId    int    `json:"pro_id"`
	State    int    `json:"state"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
}

// //验证登录
func CheckLogin(username, password string) bool {
	var user User
	db := GetDB()
	result := db.Where("user_name = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		// 处理查询错误
		fmt.Println(result.Error)
		return false
	}
	if result.RowsAffected == 0 {
		// 没有匹配的记录
		fmt.Println("没有匹配的记录")
		return false
	} else {
		// 存在匹配的记录
		fmt.Println("存在匹配的记录")
	}
	fmt.Println("user:", user)
	return user.UserName != ""
}

// according username Get role
func GetRole(username string) string {
	var user User
	db := GetDB()
	db.Select("role").Where("user_name = ?", username).First(&user)
	return user.Role
}
//check user is EXISTS
func CheckUser(username string) bool {
	var user User
	db := GetDB()
	result := db.Where("user_name = ?", username).First(&user)
	if result.Error != nil {
		// 处理查询错误
		fmt.Println(result.Error)
		return false
	}
	if result.RowsAffected == 0 {
		// 没有匹配的记录
		fmt.Println("没有匹配的记录")
		return false
	} else {
		// 存在匹配的记录
		fmt.Println("存在匹配的记录")
		return true
	}
}

// Get all user info
func GetAllUser() []User {
	var users []User
	db := GetDB()
	result := db.Where("state = ?", 1).Find(&users)
	if result.Error != nil {
		// 处理查询错误
		fmt.Println("GetAllUser error")
		log.Println(result.Error)
	}
	return users
}

// according user_name find pro_id
func GetProID(user_name string) int {
	var user User
	db := GetDB()
	result := db.Select("pro_id").Where("user_name = ?", user_name).First(&user)
	if result.Error != nil {
		// 处理查询错误
		fmt.Println("GetProID error")
		log.Println(result.Error)
	}
	return user.ProId
}

// add new user to db
func AddUser(user User) bool {
	db := GetDB()
	result := db.Create(&user)
	if result.Error != nil {
		// 处理查询错误
		fmt.Println("AddUser error")
		log.Println(result.Error)
		return false
	}
	return true
}

// func CheckAuth(username, password string) bool {
//     var auth Auth
//     db.Select("id").Where(Auth{Username : username, Password : password}).First(&auth)
//     if auth.ID > 0 {
//         return true
//     }

//     return false
// }

// //根据username查询用户角色role
// func GetRole(username string) string {
// 	var role string
// 	db.Select("role").Where("username = ?", username).First(&role)
// 	return role
// }
