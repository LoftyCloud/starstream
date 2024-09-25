/*
 * @Descripttion: 视频类型模型 - Video Class Model
 */
package ClassModel

import (
	"time"
)

/**
 * @Descripttion: 视频类型模型 - Video Class Model
 */
type Class struct {
	Cid         int    `gorm:"primaryKey" json:"cid"`
	Classname   string `json:"classname"`
	Isdelete    int    `json:"isdelete"`
	Create_Time time.Time
	Update_Time time.Time
}

/**
 * @Descripttion: 回传类型数据 - Return type data
 */
type ClassRe struct {
	Cid       int    `json:"cid"`
	Classname string `json:"classname"`
}

/**
 * @Descripttion: 后台视频分类模型 - Admin video class struct
 */
type ClassAdminSearch struct {
	Cid       int    `json:"cid"`
	Classname string `json:"classname"`
	Page      int    `json:"page"`
	Size      int    `json:"size"`
}

/**
 * @Descripttion: 后台分类模型 - Admin class struct
 */
type ClassAdminRe struct {
	Cid         int       `json:"cid"`
	Classname   string    `json:"classname"`
	Isdelete    int       `json:"isdelete"`
	Create_Time time.Time `json:"create_time"`
}
