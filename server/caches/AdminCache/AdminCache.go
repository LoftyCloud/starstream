/*
 * @Descripttion: 后台缓冲层 - Admin Cache
 */
package AdminCache

import (
	"VideoHubGo/utils/LogUtils"
	"VideoHubGo/utils/RedisUtils"
	"context"
	"strconv"
	"time"
)

var conn = RedisUtils.RedisClient
var ctx = context.Background()

/**
 * @Descripttion: 存储后台二次编码状态 - Save back-end recode status
 */
func SaveReCodeCountList(code int) {
	err := conn.Set(ctx, "recode", code, time.Hour*2).Err()
	conn.Persist(ctx, "recode")
	if err != nil {
		LogUtils.Logger("[Redis报错] 在存储后台二次编码时出错：" + err.Error())
	}

}

/**
 * @Descripttion: 获取后台二次编码状态 - get back-end recode status
 */
func GetReCodeCache() int {
	res, err := conn.Get(ctx, "recode").Result()
	if err != nil {
		return 0
	}
	tRes, _ := strconv.Atoi(res)
	return tRes
}

/**
 * @Descripttion: 删除后台二次编码状态缓存 - Delete back-end recode status cache
 */
func ReCodeDeleteCaches() {
	_, err := conn.Del(ctx, "recode").Result()
	if err != nil {
		LogUtils.Logger("[Redis操作]删除后台二次编码时异常：" + err.Error())
	}
}
