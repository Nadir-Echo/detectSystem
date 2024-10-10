package api

import (
	"Demo/models"
	"Demo/pkg/e"
	"Demo/pkg/util"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)" json:"username"`
	Password string `valid:"Required; MaxSize(50) " json:"password"`
	// Code     string `valid:"Required; MaxSize(50) " json:"code"`
}
type UserTable struct {
	UserID   int    `gorm:"AUTO_INCREMENT" json:"user_id"`
	UserName string `json:"user_name"`
	ProId    int    `json:"pro_id"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
}
type newUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	ProId    string  `json:"proID"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
}

// 获取token
func Login(c *gin.Context) {

	var user auth
	// user.Username = c.Query("username")
	// user.Password = c.Query("password")

	if err := c.ShouldBind(&user); err != nil {
		fmt.Println("BindJSON参数绑定错误!")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(user.Username)
	fmt.Println(user.Password)

	// //验证参数
	// valid := Validator.New()

	// err := valid.Struct(user)
	// if err != nil {
	// 	for _, err := range err.(validator.ValidationErrors) {
	// 		fmt.Println(err)
	// }
	// }

	// a := auth{Username: user.Username, Password: user.Password}
	// ok, err := valid.Valid(&a)
	// fmt.Println(ok)

	// if err != nil {
	// 	fmt.Println("参数验证失败!")
	// 	log.Panicln(err)
	// }
	code := e.INVALID_PARAMS
	data := make(map[string]interface{})
	// fmt.Print(ok)
	// ok := true

	//验证用户名密码是否正确
	isExit := models.CheckLogin(user.Username, user.Password)
	//判断用户名密码
	if isExit {
		//query role
		role := models.GetRole(user.Username)
		fmt.Println(role)
		token, err := util.GenerateToken(user.Username, role)
		fmt.Print(token)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			// data["token"] = token
			data["token"] = token
			print(token)
			//成功获取
			code = e.SUCCESS
		}
		// c.JSON(http.StatusOK, gin.H{
		// 	"status": e.SUCCESS,
		// 	"token":  token,
		// })
	} else {
		//用户名密码错误
		code = e.ERROR_AUTH
	}

	//success
	// c.JSON(http.StatusOK, gin.H{@QA vvv b
	// 	"status": code,
	// 	"msg":    e.GetMsg(code),
	// 	"data":   data,
	// })
	fmt.Println(code)

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    data,
	})

}

// 添加用户信息
func AddNewUser(c *gin.Context) {
	var uInfo newUser
	
	if err := c.ShouldBind(&uInfo); err != nil {
		fmt.Println("BindJSON参数绑定错误!")
		c.JSON(http.StatusOK, gin.H{"code":400,"error": err.Error()})
		return
	}
	IsExit := models.CheckUser(uInfo.UserName)
	if IsExit {
		code := e.ERROR_USER_EXIST
		res := gin.H{
			"code":    code,
			"message": e.GetMsg(code),
		}
		c.JSON(http.StatusOK, res)
		return

	}
	id,err := strconv.Atoi(uInfo.ProId)
	if err != nil {
		fmt.Println("proID int类型转换错误")
	}
	user := models.User{
		UserName: uInfo.UserName,
		Password: uInfo.Password,
		ProId:    id,
		Role:    uInfo.Role,
		Phone:    uInfo.Phone,
		Name:     uInfo.Name,
		State: 1,
	}
	// fmt.Println(user)
	IsSuccess := models.AddUser(user)
	if !IsSuccess {
		code := e.ERROR_ADD_USER
		res := gin.H{
			"code":    code,
			"message": e.GetMsg(code),
		}
		c.JSON(http.StatusOK, res)
		return
	}
	code := e.SUCCESS
	res := gin.H{
		"code":    code,
		"message": e.GetMsg(code),
	}
	c.JSON(http.StatusOK, res)

}

// 获取user info
func GetUserInfo(c *gin.Context) {
	user := models.GetAllUser()
	userTable := []UserTable{}
	for _, v := range user {
		form := UserTable{
			UserID:   v.UserID,
			UserName: v.UserName,
			ProId:    v.ProId,
			Role:     v.Role,
			Phone:    v.Phone,
			Name:     v.Name,
		}
		userTable = append(userTable, form)
	}
	res := gin.H{
		"code": 200,
		"data": userTable,
	}
	fmt.Println(res)
	c.JSON(http.StatusOK, res)

}

// 根据token获取用户信息和角色
func Getinfo(c *gin.Context) {
	// var token string
	//get token from Request URL: http://localhost:9528/user/info?token=admin
	raw_token := c.Query("token")
	fmt.Println(raw_token)
	// generate token
	token, err := util.ParseToken(raw_token)
	if err != nil {
		code := e.ERROR_PARSE_GET_ROLE
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": e.GetMsg(code),
		})
	}
	avatar := "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"

	//get user info
	fmt.Println(token)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"roles": []string{token.Role}, "introduction": "I am a super administrator", "avatar": avatar},
		"name": token.Username,
	})

}

// LogoutHandler can be used by clients to remove the jwt cookie (if set)
func Logout(c *gin.Context) {
	c.SetCookie("X-token", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": e.SUCCESS,
		"data":   "成功退出！",
	})
}
