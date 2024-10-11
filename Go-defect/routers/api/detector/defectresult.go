package detecotr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/types"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"Demo/models"
	"Demo/pkg/e"
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

// post Protobuf to 8080 Post and receive result
func Defectimg(c *gin.Context) {
	// receive image from context
	file, err := c.FormFile("file")
	fmt.Printf("%T\n", file)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		fmt.Println("图片上传失败！")
		return
	}
	log.Println(file.Filename)
	// c.SaveUploadedFile(file,"./img")

	// read the file into memory
	imageFile, err := file.Open()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		fmt.Println("读取图片到内存失败！")
		return
	}

	defer imageFile.Close()

	imageData, err := ioutil.ReadAll(imageFile)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		fmt.Println("读取文件内容失败！")
		return
	}

	// create post request
	reqBody := bytes.NewBuffer(imageData)

	// post request
	resp, err := http.Post("http://localhost:8000/steel/defectimg", "application/octet-stream", reqBody)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		fmt.Println("图片检测请求发送失败！")
		return
	}
	defer resp.Body.Close()

	// read response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		fmt.Println("api接口读取失败！")
		return
	}
	fmt.Println(string(respBody))

	// parse response,Extract "class", "conf", "position"
	type DefectResult struct {
		Class    string      `json:"class"`
		Conf     float64     `json:"conf"`
		Position types.Slice `json:"position"`
	}


	var result DefectResult

	err = json.Unmarshal(respBody, &result)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		fmt.Println("json解析失败！")
		return
	}
	fmt.Println(result)
	// fmt.Println(result["class"])
	// fmt.Println(result["conf"])
	// fmt.Println(result["position"])

	

	c.JSON(http.StatusOK, gin.H{
		// "code": code,
		// "msg":  e.GetMsg(code),
		// "data": data,
		"code": "200",
		"msg":  e.GetMsg(200),
	})
}
// curl -X 127.0.0.1:3306/api/v1/defectimg -F"file=C:\Users\NadirEcho\Downloads\Compressed\NEU-DET\IMAGES\crazing_1.jpg" -H "Content-Type:multipart/form-data"

