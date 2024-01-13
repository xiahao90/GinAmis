package nbcache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var cacheAdapter *cache.Cache

func init() {
	// 创建一个默认过期时间为5分钟的缓存适配器
	// 每60清除一次过期的项目
	cacheAdapter = cache.New(5*time.Minute, 60*time.Second)
}

func SetCahce(k string, x interface{}, d time.Duration) {
	cacheAdapter.Set(k, x, d)
}

func GetCache(k string) (interface{}, bool) {
	return cacheAdapter.Get(k)
}

//设置cache 无时间参数
func SetDefaultCahce(k string, x interface{}) {
	cacheAdapter.SetDefault(k, x)
}
//删除 cache
func DeleteCache(k string) {
	cacheAdapter.Delete(k)
}

// Add() 加入缓存
func AddCache(k string, x interface{}, d time.Duration) {
	cacheAdapter.Add(k, x, d)
}

// IncrementInt() 对已存在的key 值自增n
func IncrementIntCahce(k string, n int) (num int, err error){
	return cacheAdapter.IncrementInt(k, n)
}
