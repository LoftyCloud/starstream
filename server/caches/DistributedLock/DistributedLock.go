/*
 * @Descripttion: 分布式锁 - Distributed Lock
 */
package DistributedLock

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
 * @Descripttion: 设置分布式锁 - Set Distributed Lock
 * @Param: lock (int)
 * @Return: Result (int)
 */
func LockFunc(key string, lock int) int {
	err := conn.Set(ctx, key, lock, time.Hour*2).Err()
	if err != nil {
		LogUtils.Logger("[Redis报错] 在设置锁时出现异常：" + err.Error())
		return 0
	}
	conn.Persist(ctx, key)
	return 1
}

/**
 * @Descripttion: 获取分布式锁 - Get Distributed Lock
 * @Return: Result (int)
 */
func GetLock(key string) int {
	count, err := conn.Get(ctx, key).Result()
	if err != nil {
		return 0
	}
	tCount, _ := strconv.Atoi(count)
	return tCount
}

/**
 * @Descripttion: 删除分布式锁 - Delete Distributed Lock
 */
func UnLockFunc(key string) {
	_, err := conn.Del(ctx, key).Result()
	if err != nil {
		LogUtils.Logger("[Redis操作]删除lockfunc时异常：" + err.Error())
	}
}
