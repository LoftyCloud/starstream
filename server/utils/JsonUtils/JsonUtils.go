/*
 * @Descripttion: Json工具 -Json Utils
 */
package JsonUtils

import (
	"github.com/gin-gonic/gin"
)

func JsonResult(code int, msg string, data any) gin.H {
	h := gin.H{"code": code, "msg": msg, "data": data}
	return h
}
