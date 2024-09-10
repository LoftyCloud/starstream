/*
 * @Descripttion: 用户服务层 - User Services
 */
package UserServices

import (
	"VideoHubGo/models/UserModel"
	"VideoHubGo/utils/DataBaseUtils"
	"VideoHubGo/utils/EncryptionUtils"
	"VideoHubGo/utils/FileUtils"
	"VideoHubGo/utils/LogUtils"
	"VideoHubGo/utils/UploadUtils"
	"time"

	uuid "github.com/satori/go.uuid"
)

var db = DataBaseUtils.GoDB()

/**
 * @Descripttion: 用户登录 - User Login
 * @Param: Account (string)
 * @Param: Password (string)
 * @Return: UserModel User
 */
func LoginUser(account string, password string) UserModel.User {
	userSalt := FindUserSalt(account)
	realPwd := EncryptionUtils.ReversePassword(password, userSalt)
	userData := UserModel.User{}
	if err := db.Table("userdata").Where("account = ? and password = ?", account, realPwd).Take(&userData).Error; err != nil {
		return userData
	}
	return userData
}

/**
 * @Descripttion:
 * @Param: Account (string)
 * @Param: Password (string)
 * @Return: Result (int)
 */
func RegisterUser(account string, password string, username string) int {
	//先检查是否有人已注册了账号
	isAcAlive, isUnAlive := FindUserAlive(account), FindUserNameAlive(username)
	if isAcAlive {
		return 2
	} else if isUnAlive {
		return 3
	}
	userData := UserModel.UserRegister{}
	pwd, salt := EncryptionUtils.EncPassword(account, password)
	userData.Password = pwd
	userData.Salt = salt
	userData.Username = username
	userData.Account = account
	userData.Isdelete = 0
	userData.Isuploader = 0
	userData.Isadmin = 0
	userData.Avatar = ""
	userData.Create_Time = time.Now()
	isDeleteUid := FindIsDeleteUser()
	if isDeleteUid != 0 {
		userData.Uid = isDeleteUid
		if err := db.Table("userdata").Save(&userData).Error; err != nil {
			LogUtils.Logger("[数据库错误 SQL Error]在操作用户注册更新时产生异常：" + err.Error())
			return 4
		}
	} else {
		if err := db.Table("userdata").Create(&userData).Error; err != nil {
			LogUtils.Logger("[数据库错误 SQL Error]在操作用户注册创建时产生异常：" + err.Error())
			return 0
		}
	}
	return 1
}

/**
 * @Descripttion: 更新用户密码 - Update User Password
 * @Param: account (string)
 * @Param: oldPassword (string)
 * @Param: newPassword (string)
 * @Return: Result (int)
 */
func UpdatePassword(account string, oldPassword string, newPassword string) int {
	userData := LoginUser(account, oldPassword)
	if userData.Uid == 0 {
		return 0
	}
	encPassword := EncryptionUtils.ReversePassword(newPassword, userData.Salt)
	db.Table("userdata").Where("uid = ?", userData.Uid).Update("password", encPassword)
	return 1
}

/**
 * @Descripttion: 用户头像保存到数据库 - User Avatar Save In DataBase
 * @Param: userId (int)
 * @Param: fileTpye (string)
 * @Return: fileName (string)
 */
func UserUploadAvatar(userId int, fileTpye string) string {
	uuidKey := uuid.NewV4()
	nowTime := time.Now().Format("2006-1-2 15:04:05.000")
	uuidValue := uuid.NewV5(uuidKey, nowTime).String()
	userData := UserModel.User{}
	db.Table("userdata").Where("uid = ?", userId).First(&userData)
	FileUtils.DeleteFile(UploadUtils.GetFilePath("user.userAvatar") + "/" + userData.Avatar)
	db.Table("userdata").Where("uid = ?", userId).Update("avatar", uuidValue+fileTpye)
	return uuidValue + fileTpye
}

/**
 * @Descripttion: 查询账号是否存在 - Query whether the account exists
 * @Param: Account (string)
 * @Return: Result (int)
 */
func FindUserAlive(account string) bool {
	tempData := UserModel.User{}
	db.Select("uid").Table("userdata").First(&tempData, "account = ?", account)
	return tempData.Uid != 0
}

/**
 * @Descripttion: 查询账号的盐值 - Query the salt value of the account
 * @Param: Account (string)
 * @Return: Salt (string)
 */
func FindUserSalt(account string) string {
	userModel := UserModel.User{}
	if err := db.Table("userdata").Select("salt").Where("account = ?", account).Find(&userModel).Error; err != nil {
		return ""
	}
	return userModel.Salt
}

/**
 * @Descripttion: 查询用户名是否被占用 - Query whether the user name is occupied
 * @Param: Username (string)
 * @Return: Result (int)
 */
func FindUserNameAlive(username string) bool {
	tempData := UserModel.User{}
	db.Select("uid").Table("userdata").First(&tempData, "username = ? ", username)
	return tempData.Uid != 0
}

/**
 * @Descripttion: 查找已经被注销删除的用户 - Find Is Deleted User Info
 * @Return: uid (int)
 */
func FindIsDeleteUser() int {
	userData := UserModel.User{}
	db.Select("uid").Table("userdata").First(&userData, "isdelete = ?", 1)
	return userData.Uid
}

/**
 * @Descripttion: 根据用户查询是否被锁定 - Whether it is locked according to user query
 * @Param: Account (string)
 * @Return: isdelete (int)
 */
func FindIsDeleteUserLogin(account string) int {
	userData := UserModel.User{}
	db.Select("uid").Table("userdata").First(&userData, "account = ?", account)
	return userData.Isdelete
}
