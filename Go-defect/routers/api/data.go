package api

import (
	"Demo/models"
	"Demo/pkg/e"
	"Demo/pkg/util"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type dataForm struct {
	Pid   int    `json:"pid"`
	Class string `json:"class"`
	ProId int    `json:"proId"`
	Time  string `json:"time"`
}
type ChartData struct {
	Name string `json:"name"`
	Value int64 `json:"value"`
}

//get chart result
func GetChartData(c *gin.Context){
	token := c.Request.Header.Get("X-Token")
	// parse token
	data, err := util.ParseToken(token)
	if err != nil {
		fmt.Println("parse token error:", err)
		code := e.ERROR_AUTH_PARSE
		fmt.Println(e.GetMsg(code))
	}
	name := data.Username
	result,err := models.GetChartResult(name)
	if err != nil {
		fmt.Println("get chart result error:",err)
		log.Fatal(err)
	}
	var chartDataList []ChartData
	fmt.Println(result)
	for _,v := range result{
		tmp := ChartData{
			Name: v.DefectResult,
			Value: v.Total,
		}
		chartDataList = append(chartDataList,tmp)
	}


	res := gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": chartDataList,
	}
	fmt.Println(res)
	c.JSON(http.StatusOK, res)

}

// get form data
func GetDataForm(c *gin.Context) {
	token := c.Request.Header.Get("X-Token")
	// parse token
	data, err := util.ParseToken(token)
	if err != nil {
		fmt.Println("parse token error:", err)
		code := e.ERROR_AUTH_PARSE
		fmt.Println(e.GetMsg(code))
	}
	role := data.Role
	var formData []dataForm
	if role == "admin" {
		formData = getForm(0)

	} else {
		username := data.Username
		proId := models.GetProID(username)
		formData = getForm(proId)

	}
	res := gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": formData,
	}
	fmt.Println(res)
	c.JSON(200, res)
}

// get ori_url and det_url according to pid
func GetDataImg(c *gin.Context) {
	pid := c.Query("pid")
	// convert pid to int
	// pid,err := strconv.Atoi(c.Param("pid"))
	id, err := strconv.Atoi(pid)
	if err != nil {
		log.Println("get pid error:", err)
		panic(err)
	}
	ori_url, det_url, err := models.GetUrl(id)
	if err != nil {
		log.Println("get url error:", err)
		panic(err)
	}
	res := gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": gin.H{
			"ori": ori_url,
			"det": det_url,
		},
	}
	c.JSON(200, res)
}

// get dafect data form database
func getForm(proId int) []dataForm {
	var result []dataForm
	var data []models.Prod
	var err error
	if proId == 0 {
		data, err = models.GetAllProResult(proId)
		if err != nil {
			log.Println("get data from pro  error:", err)
		}
	} else {
		data, err = models.GetProResult(proId)
		if err != nil {
			log.Println("get data from pro  error:", err)
		}
	}

	for _, v := range data {

		// result[i].Pid = v.Pid
		// result[i].ProId = v.ProId
		// result[i].Time = v.CreTime
		tmp, err := models.GetResultByPidString(v.Pid)
		if err != nil {
			log.Println("get data from result  error:", err)
		}
		// result[i].Class = tmp
		form := dataForm{
			Pid:   v.Pid,
			Class: tmp,
			ProId: v.ProId,
			Time:  v.CreTime.Format("2006-01-02"),
		}
		result = append(result, form)

	}
	return result
}
