package models

import (
	"fmt"
	"testing"
)

// test  GetProID function
func TestGetProID(t *testing.T) {
	pro_id := GetProID("test")
	fmt.Println(pro_id)
}

// test GetRole function
func TestGetRole(t *testing.T) {
	role := GetRole("test")
	fmt.Println(role)
}

// test GetAllUser function
func TestGetAllUser(t *testing.T) {
	user := User{
		UserName: "test111",
		Password: "test233",
		ProId:    3,
		State:    1,
		Role:     "detector",
		Phone:    "123456789",
		Name:     "刘九三",
	}
	tmp := AddUser(user)
	fmt.Println(tmp)
}
