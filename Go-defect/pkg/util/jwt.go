package util

//jwt 工具包

import (
	"Demo/pkg/setting"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}



//生成token
func GenerateToken(username, role string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		username,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime), //过期时间
			Issuer:    "superadmin",                   //签发人
		},
	}
	//生成签名对象,指定签名方法和声明
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err

}

//解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
	if tokenClaims != nil {
		fmt.Println(tokenClaims)
		//fmt.Println(tokenClaims.Claims)
		//fmt.Println(tokenClaims.Claims.(*Claims))
		//fmt.Println(tokenClaims.Claims.(*Claims).Username)
		//fmt.Println(tokenClaims.Claims.(*Claims).Password)
		//fmt.Println(tokenClaims.Claims.(*Claims).Role)
		//fmt.Println(tokenClaims.Claims.(*Claims).StandardClaims)
		//fmt.Println(tokenClaims.Claims.(*Claims).StandardClaims.ExpiresAt)
		//fmt.Println(tokenClaims.Claims.(*Claims).StandardClaims.Issuer)
		//fmt.Println(tokenClaims.Claims.(*Claims).StandardClaims.IssuedAt)
		//fmt.Println(tokenClaims.Claims.(*Claims).StandardClaims.NotBefore)
		//fmt.Println(tokenClaims.Claims.(*Claims).StandardClaims.Subject)
		//fmt.Println(tokenClaims.Claims.(*Claims).StandardClaims.Id)
		//fmt.Println(tokenClaims.Claims.(*Claims).StandardClaims.Audience)
		//fmt.Println(tokenClaims.Claims.(*Claims).StandardClaims)
		if conditions, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return conditions, nil
		}

	}
	return nil, err
}
