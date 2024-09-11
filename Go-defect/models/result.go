package models

import (
	"fmt"
	"log"
	"strings"
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

type Result struct {
	//AUTO_INCREMENT
	DefectId     int    `gorm:"column:defect_id;primary_key;AUTO_INCREMENT" json:"defect_id"`
	DefectResult string `gorm:"column:defect_result" json:"defect_result"`
	Conf         float64 `gorm:"column:conf" json:"conf"`
	Position     string  `gorm:"column:position" json:"position"`
	Pid          int    `gorm:"column:pid" json:"pid"`
}
type ChartResult struct{
  DefectResult string
	Total        int64
}

// add defect result  to result table
func AddResult(result *Result) error {
  db := GetDB()
  db = db.Table("result")
	if err := db.Create(&result).Error; err != nil {
		log.Println(err)
		return err
	}
	fmt.Print("add result success!")
	fmt.Print(result)
	return nil
}
//不重复的defect_result
func GetResultByPidString(pid int) (string, error) {
  db := GetDB()
  db = db.Table("result")
  var result []Result
  if err := db.Select("DISTINCT defect_result").Where("pid = ?", pid).Find(&result).Error; err != nil {
    log.Println(err)
    return "", err
  }
  defectResults := make([]string, len(result))
	for i, res := range result {
		defectResults[i] = res.DefectResult
	}

	resultString := strings.Join(defectResults, ",")
	return resultString, nil
}

//GetChartResultByusername
func GetChartResult(username string) ([]ChartResult, error) {
  var data []ChartResult
  db := GetDB()

  id := GetProID(username)
  if id == 0 {
    err := db.Table("result").Select("defect_result, COUNT(*) AS total").Group("defect_result").Find(&data).Error
    if err != nil {
      log.Fatal(err)
    }
  } else{
    var pid []int
    db.Table("prod").Select("pid").Where("pro_id = ?", id).Pluck("pid", &pid)
    err := db.Table("result").Select("defect_result, COUNT(*) AS total").Where("pid IN ?", pid).Group("defect_result").Find(&data).Error
    if err != nil {
      log.Fatal(err)
    }
  }

  fmt.Println(data)
  return data, nil
}