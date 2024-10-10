package util

import (
	"Demo/pkg/e"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// auth admin function to check role is admin
func Auth_admin(c *gin.Context) {
	token := c.Request.Header.Get("X-Token")
	// parse token
	data, err := ParseToken(token)
	if err != nil {
		log.Println("parse token error:", err)
		code := e.ERROR_AUTH_CHECK_TOKEN_FAIL
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": code,
			"msg":  e.GetMsg(code)})

		return
	}
	role := data.Role
	if role != "admin" {
		code := e.ERROR_AUTH_CHECK_ROLE_FAIL
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": code,
			"msg":  e.GetMsg(code)})
		return
	}
	//success auth admin return 200
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
	})

}

// auth admin function to check role is detector
func Auth_detector(c *gin.Context) {
	token := c.Request.Header.Get("X-Token")
	// parse token
	data, err := ParseToken(token)
	if err != nil {
		log.Println("parse token error:", err)
		code := e.ERROR_AUTH_CHECK_TOKEN_FAIL

		response := gin.H{
            "code": code,
            "msg":  e.GetMsg(code),
        }
		jsonData, _ := json.Marshal(response)
        fmt.Println(string(jsonData))

		c.AbortWithStatusJSON(http.StatusOK, response)
		return
	}
	role := data.Role
	if role != "detector" {
		code := e.ERROR_AUTH_CHECK_ROLE_FAIL
		response := gin.H{
            "code": code,
            "msg":  e.GetMsg(code),
        }
		jsonData, _ := json.Marshal(response)
        fmt.Println(string(jsonData))

		c.AbortWithStatusJSON(http.StatusOK, response)
		return
	}
	//success auth detector return 200
	// c.JSON(http.StatusOK, gin.H{
	// 	"code": e.SUCCESS,
	// 	"msg":  e.GetMsg(e.SUCCESS),
	// })

	// Success auth detector return 200
        response := gin.H{
            "code": e.SUCCESS,
            "msg":  e.GetMsg(e.SUCCESS),
        }
        jsonData, _ := json.Marshal(response)
        fmt.Println(string(jsonData))
        c.JSON(http.StatusOK, response)


}
