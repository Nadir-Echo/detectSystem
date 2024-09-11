package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"Demo/models"
)

type DefectData struct {
	Class    string  `json:"class"`
	Conf     float64 `json:"conf"`
	Position []int   `json:"position"`
}

// Open the file with the given filename from the multipart.FileHeader
// func convertFileHeaderToOSFile(fh *multipart.FileHeader) (*os.File, error) {
// 	file, err := fh.Open()
// 	if err != nil {
// 		log.Println(err)
// 		println("文件打开失败！")
// 		return nil, err
// 	}

// 	// Create a temporary file to write the contents of the multipart file using os

// 	tmpfile, err := os.CreateTemp("", "")
// 	if err != nil {
// 		log.Println(err)
// 		println("创建临时文件失败！")
// 		return nil, err
// 	}

// 	// Write the contents of the multipart file to the temporary file
// 	_, err = io.Copy(tmpfile, file)
// 	if err != nil {
// 		log.Println(err)
// 		println("写入临时文件失败！")
// 		return nil, err
// 	}

// 	// Seek the temporary file back to the beginning
// 	_, err = tmpfile.Seek(0, 0)
// 	if err != nil {
// 		log.Println(err)
// 		println("临时文件指针重置失败！")
// 		return nil, err
// 	}

// 	return tmpfile, nil
// }

// detecttoyolov5 is a function to receive image from front end and send it to defect detection model
func detecttoyolov5(imageFile *os.File, filename string) (result *http.Response, err error) {

	// Create a new multipart form
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// Add the imageFile to the form
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("创建表单失败！")
		return
	}

	// Copy the imageFile to the form
	_, err = io.Copy(part, imageFile)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("复制文件到表单失败！")
		return
	}

	// close the multipart writer
	err = writer.Close()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("关闭表单失败！")
		return
	}
	// Create a new POST request with the form data
	req, err := http.NewRequest("POST", "http://127.0.0.1:8000/steel/defectimg", body)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("创建POST请求失败！")
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept", "*/*")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("发送POST请求失败")
	}
	// defer resp.Body.Close()

	return resp, err

}

// []int convert to string
func IntArrayToString(a []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ",", -1), "[]")
}

// Defectimg is a function to receive image from front end and send it to defect detection model
func Defectimg(c *gin.Context) {

	// receive image from context
	//  It returns a *multipart.FileHeader and an error.
	file, err := c.FormFile("file")

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("从前端获取图片失败！")
		return
	}
	// use file.Header to get file information
	fmt.Printf("Uploaded file: %s (size: %d bytes, content type: %s)\n",
		file.Filename, file.Size, file.Header.Get("Content-Type"))
	// filename := file.Filename
	//save image to local
	c.SaveUploadedFile(file, "1.jpeg")
	filename := "1.jpeg"

	//upload image to Tencent COS
	img_upload, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("打开文件失败！")
		return
	}
	originalMode := "orignal"

	ori_url, err := UploadImageToTencentCloudCos(img_upload, originalMode)
	fmt.Println("ori_url: ", ori_url)
	if err != nil {
		fmt.Println(err.Error())
	}
	img_upload.Close()

	// convert file to os.file
	// imageFile, err := convertFileHeaderToOSFile(file)

	// open image file
	imageFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(err)
		fmt.Println("打开文件失败！")
		return
	}

	// use funtion to detect defect
	// print(filename)
	resp, err := detecttoyolov5(imageFile, filename)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("调用模型失败！")
		return
	}

	imageFile.Close()

	// Read the response body
	fmt.Println("resp:", resp)
	respBody, err := io.ReadAll(resp.Body)
	// var respBody http.Response
	// err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("读取响应失败！")

	}
	fmt.Println("Length:", len(respBody))
	// defect img url
	var result_url string
	// defect exist or not
	var err_result int
	// defect data
	var defectData []DefectData
	// if no defect, set flag to 0 else set flag to 1 and todo next
	if len(respBody) < 3 {
		fmt.Println("没有检测到缺陷！")
		err_result = 0
	} else {
		fmt.Println("检测到缺陷！")
		err_result = 1
		if err := json.Unmarshal(respBody, &defectData); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			fmt.Println("解析JSON失败！")
		}

		for _, defect := range defectData {
			fmt.Printf("Class: %s, Conf: %f, Position: %v\n", defect.Class, defect.Conf, defect.Position)
			fmt.Printf("position type: %T\n", defect.Position)
		}

		// Draw squares on the image using the function of drawSquares
		// Decode the image
		imageDraw, err := os.Open(filename)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("打开文件失败！")
			return
		}
		img, _, err := image.Decode(imageDraw)
		if err != nil {
			println("decode image error")
			log.Println(err)
		}

		//use the func of drawSquare
		result_url = DrawSquare(img, defectData, color.RGBA{255, 0, 0, 255})
		imageDraw.Close()
		fmt.Println("result_url: ", result_url)

	}


	username, IsExist := c.Get("username")
	if !IsExist {
		fmt.Println("username is not exist")
	}
	fmt.Println("username: ", username)
	ProId := models.GetProID(username.(string))
	fmt.Println("ProId: ", ProId)

	//save to Prod table
	var data models.Prod
	data.OriUrl = ori_url
	data.DetUrl = result_url
	data.ProId = ProId
	data.ErrResult = err_result
	data.CreTime = time.Now()

	Pid, err := models.AddProd(&data)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("保存到数据库失败！")
	}

	// create a result struct

	for _, defect := range defectData {
		result := models.Result{}
		// fmt.Printf("Class: %s, Conf: %f, Position: %v\n", defect.Class, defect.Conf, defect.Position)
		result.DefectResult = defect.Class
		result.Conf = defect.Conf
		result.Position = IntArrayToString(defect.Position)
		result.Pid = Pid
		models.AddResult(&result)

	}
	c.JSON(http.StatusOK, gin.H{
		"code":       200,
		"msg":        "success",
		"ori_url":    ori_url,
		"result_url": result_url,
	})

}
