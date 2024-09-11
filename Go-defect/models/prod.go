package models

import (
	"log"
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

type Prod struct {
    //pid gorm 定义自增字段
    Pid       int    `json:"pid" gorm:"primary_key;AUTO_INCREMENT"`
    OriUrl    string `json:"ori_url"`
    DetUrl    string `json:"det_url"`
    ProId     int    `json:"pro_id"`
    ErrResult int    `json:"err_result"`
    CreTime   time.Time `json:"cre_time"`
}

// save the image info(include ori_url,det_url,pro_id,pid是自增字段)  to prod table
func AddProd(data *Prod) (id int, err error) {
    db := GetDB()
    err = db.Create(&data).Error
    return data.Pid, err

}
//GetProResult 
func GetProResult(pro_id int) (prod []Prod, err error) {

    db := GetDB()
    err = db.Table("prod").Where("pro_id = ? AND err_result=?", pro_id,1).Find(&prod).Error
    return prod, err
}

func GetAllProResult(pro_id int) (prod []Prod, err error) {
  if pro_id!=0 {
    log.Println("pro_id is not 0!not allow use it!")
    return nil ,err
  }
  db := GetDB()
  err = db.Table("prod").Where("err_result=?",1).Find(&prod).Error
  return prod, err
}
// GetUrlByPid
func GetUrl(pid int) (ori_url string,det_url string, err error) {
    var prod Prod
    db := GetDB()
    err = db.Table("prod").Where("pid = ?", pid).First(&prod).Error
    return prod.OriUrl, prod.DetUrl,err
}