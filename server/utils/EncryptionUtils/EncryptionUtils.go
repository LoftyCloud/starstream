/*
 * @Descripttion: 加密工具 - Encryption Utils
 * @LastEditors: William Wu
 * @LastEditTime: 2022/05/28 上午 11:42
 */

package EncryptionUtils

import (
	"VideoHubGo/utils/LogUtils"
	"crypto/md5"
	"encoding/hex"
	"os"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

/**
 * @Descripttion: 加密用户密码 - Encryption User Password
 * @Param: Account (string)
 * @Param: Password (string)
 * @Return: EncryptionPassword , Salt
 */
func EncPassword(account string, password string) (string, string) {
	//读取配置文件 - Read The Configuration File
	path, err := os.Getwd()
	if err != nil {
		LogUtils.Logger(err.Error())
	}
	config := viper.New()
	config.AddConfigPath(path + "/configs")
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	//尝试进行配置读取 - Try Reading Configuration
	if err := config.ReadInConfig(); err != nil {
		LogUtils.Logger(err.Error())
	}

	keySalt := config.GetString("salt.value")
	keyUuid := uuid.NewV4()
	uuidValue := uuid.NewV5(keyUuid, account+keySalt).String()
	uuidMd5 := md5.Sum([]byte(uuidValue))
	salt := hex.EncodeToString(uuidMd5[:])
	obscurePwd := md5.Sum([]byte(password + salt))
	userPassword := hex.EncodeToString(obscurePwd[:])
	return userPassword, salt
}

/**
 * @Descripttion: 逆向密码 - Reverse Password
 * @Param: Password (string)
 * @Param: salt (string)
 * @Return: User Password
 */
func ReversePassword(password string, salt string) string {
	userPassword := md5.Sum([]byte(password + salt))
	result := hex.EncodeToString(userPassword[:])
	return result
}
