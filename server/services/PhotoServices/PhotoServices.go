/*
 * @Descripttion: 图片服务层 - Photo Services
 */
package PhotoServices

import (
	"VideoHubGo/models/PhotoModel"
	"VideoHubGo/utils/DataBaseUtils"
	"time"
)

var db = DataBaseUtils.GoDB()

/**
 * @Descripttion: 获取图片集 - Get photo list
 * @Param: size (int)
 * @Param: page (int)
 * @Return: photoData (PhotoModel.PhotoModelRe), count (int)
 */
func GetPhotoList(size int, offset int) []PhotoModel.PhotoModelRe {
	var photoData []PhotoModel.PhotoModelRe
	db.Table("picturedata").Where("isdelete = ?", 0).Order("pid DESC").Limit(size).Offset(offset).Find(&photoData)
	return photoData
}

/**
 * @Descripttion: 获取图片总数 - Get Photo Count
 * @Return: count (int)
 */
func GetCountPhotoList() int {
	var count int64
	db.Select("pid").Table("picturedata").Where("isdelete = ?", 0).Count(&count)
	return int(count)
}

/**
 * @Descripttion: 存储图片到数据库 - Save photo to sql
 * @Param: plast (string)
 * @Return: Result (int)
 */
func SavePhoto(plast string) int {
	photoData := PhotoModel.PhotoModel{}
	photoData.Plast = plast
	photoData.Create_Time = time.Now()
	photoData.Update_Time = time.Now()
	delPid := FindIsdeletePhotoPid()
	if delPid != 0 {
		photoData.Pid = delPid
		photoData.Isdelete = 0
		db.Table("picturedata").Save(&photoData)
	} else {
		db.Table("picturedata").Create(&photoData)
	}
	return photoData.Pid
}

/**
 * @Descripttion: 查询被删除的数据 - Find isdelete data
 * @Return: Result (int)
 */
func FindIsdeletePhotoPid() int {
	photoData := PhotoModel.PhotoModel{}
	db.Table("picturedata").Where("isdelete = ?", 1).First(&photoData)
	return photoData.Pid
}
