/*
 * @Descripttion: 后台模型 - Admin Model
 */
package AdminModel

import "time"

/**
 * @Descripttion: 视频组合模型 - Video group struct
 */
type VideoDelete struct {
	Vid int `json:"vid"`
}

/**
 * @Descripttion: 视频删除模型 - Video delete struct
 */
type VideoDeleteGroup struct {
	Group []VideoDelete `json:"group"`
}

/**
 * @Descripttion: 视频组合修改模型 - Video group edit struct
 */
type VideoBatchCid struct {
	Cid   int           `json:"cid"`
	Group []VideoDelete `json:"group"`
}

/**
 * @Descripttion: 分类删除模型 - Class delete struct
 */
type ClassDelete struct {
	Cid int `json:"cid"`
}

/**
 * @Descripttion: 分类组合删除模型 - Class delete group struct
 */
type ClassDeleteGroup struct {
	Group []ClassDelete `json:"group"`
}

/**
 * @Descripttion: 用户组合模型 - User group struct
 */
type UserDelete struct {
	Uid int `json:"uid"`
}

/**
 * @Descripttion: 用户删除组合模型 - User delete group struct
 */
type UserDeleteGroup struct {
	Group []UserDelete `json:"group"`
}

/**
 * @Descripttion: 用户管理员权限模型 - User admin authority Struct
 */
type UserIsadmin struct {
	Uid     int `json:"uid"`
	Isadmin int `json:"isadmin"`
}

/**
 * @Descripttion: 用户上传者权限模型 = User uploader authority Struct
 */
type UserIsuploader struct {
	Uid        int `json:"uid"`
	Isuploader int `json:"isuploader"`
}

/**
 * @Descripttion: 图片组合模型 - Photo group struct
 */
type PhotoDelete struct {
	Pid int `json:"pid"`
}

/**
 * @Descripttion: 图片删除组合模型 - Photo delete group struct
 */
type PhotoDeleteGroup struct {
	Group []PhotoDelete `json:"group"`
}

/**
 * @Descripttion: 后台数据模型 - Back-end data struct
 */
type AdminDashboard struct {
	Id      int `json:"id"`
	Re_Code int `json:"re_code"`
}

/**
 * @Descripttion: 后台文件列表 - Admin file list
 */
type AdminFileList struct {
	Id   int       `json:"id"`
	Name string    `json:"name"`
	Size int64     `json:"size"`
	Date time.Time `json:"date"`
}

/**
 * @Descripttion: 删除文件模型 - Delete struct
 */
type DeleteStruct struct {
	Name string `json:"name"`
}

/**
 * @Descripttion: 后台登录日志模型 - Admin login log struct
 */
type UserLog struct {
	Id          int       `gorm:"primaryKey" json:"id"`
	Uid         int       `json:"uid"`
	Account     string    `json:"account"`
	Username    string    `json:"username"`
	Ip          string    `json:"ip"`
	Isdelete    int       `json:"isdelete"`
	Create_Time time.Time `json:"create_time"`
	Update_Time time.Time `json:"update_time"`
}
