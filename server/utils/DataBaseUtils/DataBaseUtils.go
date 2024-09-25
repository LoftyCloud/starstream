/*
 * @Descripttion: 数据库工具 - DataBase Utils
 */
package DataBaseUtils

import (
	"VideoHubGo/utils/LogUtils"
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB

/**
 * @Descripttion: 数据库连接 - DataBase Connection
 */
func init() {
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

	host := config.GetString("database.host")
	port := config.GetString("database.port")
	databasename := config.GetString("database.databasename")
	username := config.GetString("database.username")
	password := config.GetString("database.password")

	//拼接MySQL连接地址 - Splicing Mysql Connection Address
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, databasename)
	// fmt.Println(url)
	var errdb error

	_db, errdb = gorm.Open(mysql.Open(url), &gorm.Config{})
	if errdb != nil {
		LogUtils.Logger(errdb.Error())
	}

	sqlDB, _ := _db.DB()
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	// sqlDB.SetMaxIdleConns(100) //设置最大连接数 - Set Max SQL Connection
	// sqlDB.SetMaxIdleConns(20)  //设置最大空闲连接数 - Set Max Free Connection

	// 自动迁移模式
	// 把models文件夹中定义的模型进行迁移
	// _db.AutoMigrate(&UserModel.User{}, &UserModel.EditUser{}, &UserModel.UserAdminRe{}, &UserModel.UserAdminSearch{}, &UserModel.UserRegister{}, &UserModel.UserUpdatePassword{})
	// _db.AutoMigrate(&SystemModel.UsageStat{})
	// _db.AutoMigrate(&RelationModel.Relation{}, &RelationModel.RelationRequest{}, &RelationModel.RelationRequestBody{}, &RelationModel.RelationRequestClass{}, &RelationModel.RelationRequestSearch{})
	// _db.AutoMigrate(&PhotoModel.PhotoModel{}, &PhotoModel.PhotoSearch{}, &PhotoModel.PhotoModelRe{})
	// _db.AutoMigrate(&ClassModel.Class{}, &ClassModel.ClassAdminRe{}, &ClassModel.ClassAdminSearch{}, &ClassModel.ClassRe{})
	// _db.AutoMigrate(&AdminModel.AdminDashboard{}, &AdminModel.AdminFileList{}, &AdminModel.ClassDelete{}, &AdminModel.ClassDeleteGroup{}, &AdminModel.DeleteStruct{}, &AdminModel.PhotoDelete{}, &AdminModel.PhotoDeleteGroup{}, &AdminModel.UserDelete{}, &AdminModel.UserDeleteGroup{}, &AdminModel.UserIsadmin{}, &AdminModel.UserIsuploader{}, &AdminModel.UserLog{}, &AdminModel.VideoBatchCid{}, &AdminModel.VideoDelete{}, &AdminModel.VideoDeleteGroup{})
	// _db.AutoMigrate(&VideoModel.Video{}, &VideoModel.VideoAdminRe{}, &VideoModel.VideoAdminSearch{}, &VideoModel.VideoRe{}, &VideoModel.VideoEdit{}, &VideoModel.VideoRequest{}, &VideoModel.VideoRequestClass{}, &VideoModel.VideoRequestSearch{})

}

/**
 * @Descripttion: 数据库连接对象 - DataBase Connection Object
 */
func GoDB() *gorm.DB {
	return _db
}
