/*
 * @Descripttion: 用户收藏控制层 - Relation Controller
 */
package RelationController

import (
	"VideoHubGo/models/RelationModel"
	"VideoHubGo/services/RelationServices"
	"VideoHubGo/utils/JsonUtils"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @Descripttion: 获取用户收藏信息 - Get user favorite information
 * @Param: 分页 - page (int)
 * @Param: 数据条数 - size (int)
 * @Return: Json
 */
func GetRelationList(ctx *gin.Context) {
	requestBody := RelationModel.RelationRequest{}
	err := ctx.BindJSON(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(600, "参数错误 - Parameter error", ""))
		return
	}

	page := requestBody.Page
	size := requestBody.Size
	if page < 1 {
		page = 1
	}
	if size < 20 {
		size = 20
	}
	if size > 40 {
		size = 20
	}
	offset := size * (page - 1)
	uid := ctx.MustGet("uid").(int)
	videoData, count := RelationServices.FindRelationByVideoList(uid, size, offset)
	rData := map[string]interface{}{"list": videoData, "count": count}
	ctx.JSON(http.StatusOK, JsonUtils.JsonResult(200, "200", rData))
}

/**
 * @Descripttion: 搜索用户收藏信息 - Search User Favorite information
 * @Param: 分类 - cid (int)
 * @Param: 分页 - page (int)
 * @Param: 数据条数 - size (int)
 * @Return:
 */
func FindRelationClassList(ctx *gin.Context) {
	requestBody := RelationModel.RelationRequestClass{}
	err := ctx.BindJSON(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(600, "参数错误 - Parameter error", ""))
		return
	}

	page := requestBody.Page
	size := requestBody.Size
	cid := requestBody.Cid
	if page < 1 {
		page = 1
	}
	if size < 20 {
		size = 20
	}
	if size > 40 {
		size = 20
	}
	if cid < 0 {
		cid = 0
	}
	offset := size * (page - 1)
	uid := ctx.MustGet("uid").(int)
	videoData, count := RelationServices.FindRelationByVideoList_Class(uid, cid, size, offset)
	rData := map[string]interface{}{"list": videoData, "count": count}
	ctx.JSON(http.StatusOK, JsonUtils.JsonResult(200, "200", rData))
}

/**
 * @Descripttion: 查询收藏中的视频 - Query videos in your collection
 * @Param: 分类 - cid (int)
 * @Param: 关键字 - key (string)
 * @Param: 分页 - page (int)
 * @Param: 数据条数 - size (int)
 * @Return: Json
 */
func SearchRelationClassList(ctx *gin.Context) {
	requestBody := RelationModel.RelationRequestSearch{}
	err := ctx.BindJSON(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(600, "参数错误 - Parameter error", ""))
		return
	}

	page := requestBody.Page
	size := requestBody.Size
	cid := requestBody.Cid
	key := requestBody.Key
	if page < 1 {
		page = 1
	}
	if size < 20 {
		size = 20
	}
	if size > 40 {
		size = 20
	}
	if cid < 0 {
		cid = 0
	}
	offset := size * (page - 1)
	uid := ctx.MustGet("uid").(int)
	videoData, count := RelationServices.SearchRelationByVideoList_Class(uid, cid, key, size, offset)
	rData := map[string]interface{}{"list": videoData, "count": count}
	ctx.JSON(http.StatusOK, JsonUtils.JsonResult(200, "200", rData))
}

/**
 * @Descripttion: 添加收藏 - Add Relation
 * @Param: 用户ID - uid (int)
 * @Param: 视频ID - vid (int)
 * @Return: Json
 */
func RelationVideo(ctx *gin.Context) {
	requestBody := RelationModel.RelationRequestBody{}
	err := ctx.BindJSON(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(600, "参数错误 - Parameter error", ""))
		return
	}
	uid := ctx.MustGet("uid").(int)
	vid := requestBody.Vid
	rows := RelationServices.IsRelation(uid, vid)
	if rows == 0 {
		result := RelationServices.AddRelation(uid, vid)
		if result == 0 {
			ctx.JSON(http.StatusOK, JsonUtils.JsonResult(202, "在添加收藏时产生了未知的异常 - An unknown exception occurred while adding a collection", ""))
		}
		rData := map[string]interface{}{"relation": result}
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(200, "200", rData))
	} else {
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(201, "你已经收藏过该视频 - You have already collected this video", ""))
	}
}

/**
 * @Descripttion: 取消收藏 - Delete Relation
 * @Param: 用户ID - uid (int)
 * @Param: 视频ID - vid (int)
 * @Return: Json
 */
func RemoveRelation(ctx *gin.Context) {
	requestBody := RelationModel.RelationRequestBody{}
	err := ctx.BindJSON(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(600, "参数错误 - Parameter error", ""))
		return
	}
	uid := ctx.MustGet("uid").(int)
	vid := requestBody.Vid
	result := RelationServices.DeleteRelation(uid, vid)
	if result != 1 {
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(201, "在删除收藏时产生了未知的异常 - An unknown exception occurred while deleting a collection", ""))
	}
	ctx.JSON(http.StatusOK, JsonUtils.JsonResult(200, "200", ""))
}

/**
 * @Descripttion: 检查用户是否已经收藏 - Check User Is Already Relation
 * @Param: 用户ID - uid (int)
 * @Param: 视频ID - vid (int)
 * @Return: Json
 */
func IsRelation(ctx *gin.Context) {
	requestBody := RelationModel.RelationRequestBody{}
	err := ctx.BindJSON(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(600, "参数错误 - Parameter error", ""))
		return
	}
	uid := ctx.MustGet("uid").(int)
	vid := requestBody.Vid
	rows := RelationServices.IsRelation(uid, vid)
	if rows == 0 {
		rData := map[string]interface{}{"relation": rows}
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(200, "你没有收藏该视频 - You not already collected this video", rData))
	} else {
		rData := map[string]interface{}{"relation": rows}
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(201, "你已经收藏过该视频 - You have already collected this video", rData))
	}
}
