/*
 * @Descripttion: 用户关系模型 - Relation Model
 */
package RelationModel

/**
 * @Descripttion: 用户关系模型 - Relation Model
 */
type Relation struct {
	Id       int `gorm:"primaryKey" json:"id"`
	Uid      int `json:"uid"`
	Vid      int `json:"vid"`
	Isdelete int
}

/**
 * @Descripttion: Relation普通类型请求体 - Relation common type request struct
 */
type RelationRequest struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

/**
 * @Descripttion: Relation普通类型分类请求体 - Relation common type request Class struct
 */
type RelationRequestClass struct {
	Cid  int `json:"cid"`
	Page int `json:"page"`
	Size int `json:"size"`
}

/**
 * @Descripttion: Relation普通类型搜索请求体 - Relation common type request Saerch struct
 */
type RelationRequestSearch struct {
	Cid  int    `json:"cid"`
	Key  string `json:"key"`
	Page int    `json:"page"`
	Size int    `json:"size"`
}

/**
 * @Descripttion: Relation添加删除请求体 - Relation Add Delete Request Struct
 */
type RelationRequestBody struct {
	Uid int `json:"uid"`
	Vid int `json:"vid"`
}
