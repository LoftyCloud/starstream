/*
 * @Descripttion: 视频观看控制层 - Watch Controller
 */
package WatchController

import (
	"VideoHubGo/caches/VideoCache"
	"VideoHubGo/models/WatchModel"
	"VideoHubGo/services/WatchServices"
	"VideoHubGo/utils/JsonUtils"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @Descripttion: 查询视频详细信息 - Query video details
 * @Param: vid (int)
 * @Return: Json
 */
func GetVideoDetail(ctx *gin.Context) {
	requestBody := WatchModel.WatchRequest{}
	err := ctx.BindJSON(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(600, "参数错误 - Parameter error", ""))
		return
	}
	vid := requestBody.Vid
	uid := ctx.MustGet("uid").(int)
	videoData := WatchServices.GetVideoDetail(vid)
	reId := WatchServices.IsRelation(uid, vid)
	rData := map[string]interface{}{"list": videoData, "relation": reId}
	ctx.JSON(http.StatusOK, JsonUtils.JsonResult(200, "200", rData))
}

/**
 * @Descripttion: 增加视频流量 - Increase video traffic
 * @Param: vid (int)
 * @Return: Json
 */
func PlusVideoWatch(ctx *gin.Context) {
	requestBody := WatchModel.WatchRequest{}
	err := ctx.BindJSON(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(600, "参数错误 - Parameter error", ""))
		return
	}
	vid := requestBody.Vid
	result := WatchServices.PlusVideoWatch(vid)
	if result != 1 {
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(201, "增加视频流量时产生了异常 - Exception occurred while increasing video traffic\n\n", ""))
	}
	VideoCache.VideoDeleteCaches()
	ctx.JSON(http.StatusOK, JsonUtils.JsonResult(200, "200", ""))
}
