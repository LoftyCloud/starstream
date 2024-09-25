/*
 * @Descripttion: 无效路由拦截器 - No Route Middlewares
 */
package NoRouteMiddleware

import (
	"VideoHubGo/utils/JsonUtils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoRouteHttp(ctx *gin.Context) {
	//返回信息 - Return Message
	ctx.JSON(http.StatusOK, JsonUtils.JsonResult(400, "无效访问 - Invalid access", ""))
	return
}

func NoMethodHttp(ctx *gin.Context) {
	//返回信息 - Return Message
	ctx.JSON(http.StatusOK, JsonUtils.JsonResult(400, "不允许的方法 - Method Not Allowed", ""))
	return
}
