/*
 * @Descripttion: 用户模型 - User Model
 */
package UserModel

import (
	"time"
)

/**
 * @Descripttion: 用户模型 - User Model
 */
type User struct {
	Uid         int    `gorm:"primaryKey" json:"uid"`
	Account     string `json:"account"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Salt        string
	Avatar      string
	Isadmin     int
	Isuploader  int
	Isdelete    int
	Create_Time time.Time
	Update_Time time.Time
}

/**
 * @Descripttion: 用户注册模型 - User Register Model
 */
type UserRegister struct {
	Uid         int    `gorm:"primaryKey" json:"uid"`
	Account     string `json:"account"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Salt        string
	Isadmin     int
	Isuploader  int
	Isdelete    int
	Avatar      string
	Create_Time time.Time
}

/**
 * @Descripttion: 用户更新密码模型 - User Update Password Model
 */
type UserUpdatePassword struct {
	Account     string `json:"account"`
	Password    string `json:"password"`
	RePassword  string `json:"repassword"`
	NewPassword string `json:"newpassword"`
}

/**
 * @Descripttion: 用户编辑模型 - User Edit Struct
 */
type EditUser struct {
	Uid        int    `gorm:"primaryKey" json:"uid"`
	Account    string `json:"account"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Salt       string
	Avatar     string `json:"avatar"`
	Isadmin    int    `json:"isadmin"`
	Isuploader int    `json:"isuploader"`
	Isdelete   int    `json:"isdelete"`
}

/**
 * @Descripttion: 后台搜索用户模型 - Admin search user struct
 */
type UserAdminSearch struct {
	Uid      int    `json:"uid"`
	Account  string `json:"account"`
	Username string `json:"username"`
	Page     int    `json:"page"`
	Size     int    `json:"size"`
}

/**
 * @Descripttion: 后台用户模型 - Admin user struct
 */
type UserAdminRe struct {
	Uid         int       `json:"uid"`
	Account     string    `json:"account"`
	Username    string    `json:"username"`
	Avatar      string    `json:"avatar"`
	Isadmin     int       `json:"isadmin"`
	Isuploader  int       `json:"isuploader"`
	Isdelete    int       `json:"isdelete"`
	Create_Time time.Time `json:"create_time"`
}
