package config

import (
	"datawiz-aiservices/pkg/config"
)

func init() {

	config.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{

			"host":     config.Env("REDIS_HOST", "127.0.0.1"),
			"port":     config.Env("REDIS_PORT", "6379"),
			"password": config.Env("REDIS_PASSWORD", ""),

			// 业务类存储使用 1 (图片验证码、短信验证码、会话)
			"database": config.Env("REDIS_MAIN_DB", 1),

			// 缓存 cache 包使用 0 ，缓存清空理应当不影响业务
			"database_cache": config.Env("REDIS_CACHE_DB", 0),
		}
	})
}

// 启动用docker 命令
// $ docker run --name redisRJ -d -p 6379:6379 redis
// $ docker run --name redisRJ -d -p 6379:6379 redis redis-server --appendonly yes
// $ docker run --name redisRJ --link redisRJ:redis -d application-that-uses-redis
