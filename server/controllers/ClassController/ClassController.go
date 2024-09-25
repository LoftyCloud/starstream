/*
 * @Descripttion: 视频分类控制器 - Video Class Controller
 */
package ClassController

import (
	"VideoHubGo/caches/ClassCache"
	"VideoHubGo/services/ClassServices"
	"VideoHubGo/utils/JsonUtils"
	"VideoHubGo/utils/LogUtils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

// 使用SingleFlight防止Redis缓存击穿
var sfg singleflight.Group

/**
 * @Descripttion: 获取所有的视频分类 - Get All Video Class List
 * @Return: 数据 - Data
 */
func GetClassList(ctx *gin.Context) {

	ret, err, _ := sfg.Do("getclass", func() (interface{}, error) {
		var classData = ClassCache.ClassGetListCache()
		if classData == nil {
			classData = ClassServices.FindAllClass()
			//存入Redis - Save In Redis
			ClassCache.ClassWriteCache(classData)
		}
		return classData, nil
	})
	if err != nil {
		LogUtils.Logger("[GIN ERROR]在处理getclass缓存击穿中异常：" + err.Error())
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(201, "未知的错误-击穿", ""))
		return
	}

	ctx.JSON(http.StatusOK, JsonUtils.JsonResult(200, "200", ret))
}
