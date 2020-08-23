package util


import (
	"encoding/json"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/cache"
	"os"
)

const PAGELIMIT = 20

/**
 * 获取redis连接实例
 */
func GetRedis() (adapter cache.Cache, err error) {
	redisKey := beego.AppConfig.String("rediskey")
	redisAddr := beego.AppConfig.String("redisaddr")
	redisPort := beego.AppConfig.String("redisport")
	redisdbNum := beego.AppConfig.String("redisdbnum")

	redis_config_map := map[string] string {
		"key": redisKey,
		"conn": redisAddr + ":" + redisPort,
		"dbNum": redisdbNum,
	}
	redis_config, _ := json.Marshal(redis_config_map)

	cache_conn, err := cache.NewCache("redis", string(redis_config))
	if err != nil {
		return nil, err
	}

	return cache_conn, nil
}

/**
 * 判断当前path是否存在的工具方法
 */
