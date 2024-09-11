package v1

import (
	"Demo/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Result struct {
	ProdId      int       `json:"pro_id"`
	DefectId    int       `json:"defect_id" `
	PicOri      string    `json:"pic_ori"`
	PicDet      string    `json:"pic_det"`
	ResultId    int       `json:"result_id" gorm:"primaryKey"`
	CreatedTime time.Time `json:"created_time"`
}

// 获取检测结果
func GetResult(c *gin.Context) {
	prodId := 1
	db := models.GetDB()
	u := []Result{}
	db.Where("prod_id=?", prodId).Find(&u)

	fmt.Printf("%T\n", u)
	fmt.Println("test")
	fmt.Println(u)

}

// 新增检测结果
func AddResult(c *gin.Context) {

}
