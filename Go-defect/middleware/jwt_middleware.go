package middleware

import (
	"Demo/pkg/e"
	"Demo/pkg/util"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// jwt中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		//从请求头中获取token
		token := c.Request.Header.Get("X-Token")
		//验证token是否为空
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			checkToken := strings.Split(token, " ")
			if len(checkToken) == 0 {
				code = e.INVALID_PARAMS
			} else {
				//验证token是否有效
				claims, err := util.ParseToken(checkToken[1])
				if err != nil {
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				} else if !claims.VerifyExpiresAt(time.Now(), false) {
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				}
			}
		}
		//验证token是否有效
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			//验证token是否有效
			c.Abort()
			return
		}
		c.Next()
	}
}

// check role is admin,if not return ERROR_AUTH_CHECK_ROLE_FAIL
func JWT_admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = e.SUCCESS

		//从请求头中获取token
		token := c.Request.Header.Get("X-Token")
		// parse token
		data, err := util.ParseToken(token)
		if err != nil {
			log.Println("parse token error:", err)
		}

		role := data.Role
		if role != "admin" {
			code = e.ERROR_AUTH_CHECK_ROLE_FAIL
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": code,
			"msg":  e.GetMsg(code),})
			// c.JSON(http.StatusUnauthorized, gin.H{
			// 	"code": code,
			// 	"msg":  e.GetMsg(code),
			// })
		}
		//create a new context
		// ctx := context.WithValue(c.Request.Context(),"username",data.Username)
		// ctx = context.WithValue(ctx,"role",role)
		// c.Set("context",ctx)
		c.Set("username", data.Username)
		c.Set("role", role)
		//判断角色是否为admin

		c.Next()
	}
}

// check role is detector,if not return ERROR_AUTH_CHECK_ROLE_FAIL
func JWT_detector() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = e.SUCCESS
		//从请求头中获取token
		token := c.Request.Header.Get("X-Token")
		// parse token
		data, err := util.ParseToken(token)
		if err != nil {
			log.Println("parse token error:", err)
		}

		role := data.Role
		//判断角色是否为detector
		fmt.Println("role:", role)
		fmt.Println("code", code)
		if role != "detector" {
			code = e.ERROR_AUTH_CHECK_ROLE_FAIL
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": 50009,
				"msg":  e.GetMsg(code),
			})
		}

		c.Set("username", data.Username)
		c.Set("role:", role)
		fmt.Println("成功创建context!")
		c.Next()
	}
}
