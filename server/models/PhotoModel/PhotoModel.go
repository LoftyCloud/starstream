/*
 * @Descripttion: 图片模型 - Photo Model
 */
package PhotoModel

import "time"

/**
 * @Descripttion: 图片基本模型 - Photo base model
 */
type PhotoModel struct {
	Pid         int    `gorm:"primaryKey" json:"pid"`
	Plast       string `json:"plast"`
	Isdelete    int    `json:"isdelete"`
	Create_Time time.Time
	Update_Time time.Time
}

/**
 * @Descripttion: 图片数据返回模型 - Photo data return model
 */
type PhotoModelRe struct {
	Pid         int       `json:"pid"`
	Plast       string    `json:"plast"`
	Isdelete    int       `json:"isdelete"`
	Create_Time time.Time `json:"create_time"`
}

/**
 * @Descripttion: 图片搜索模型 - Photo search model
 */
type PhotoSearch struct {
	Page int `json:"page"`
	Size int `json:"size"`
}
