package rediss

import "time"

type RedisAPI interface {
	//存储一个持久化值
	SetValue(key string, value interface{}) (err error)
	//获取一个持久化值，转换成指定格式
	GetValue(key string, valuePointer interface{}) (err error)
	//删除一个持久化值
	DelValue(keys ...interface{}) (err error)
}

func NewRedis(db int, host string, password string, defaultExpiration time.Duration) RedisAPI {
	return NewRedisPool(db, host, password, defaultExpiration)
}
