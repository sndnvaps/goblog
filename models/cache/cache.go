package cache

import (
	"github.com/astaxie/beego/cache"
)

var _instance cache.Cache
var _init = false

func init() {
	cache.Register("LocalCache", NewCache())
}

type LocalCache struct {
	data map[string]interface{}
}

func (this *LocalCache) Get(key string) interface{} {
	if v, ok := this.data[key]; ok {
		return v
	}
	return ""
}

func (this *LocalCache) Put(key string, val interface{}, timeout int64) error {
	this.data[key] = val
	return nil
}

func (this *LocalCache) Delete(key string) error {
	delete(this.data, key)
	return nil
}

func (this *LocalCache) Incr(key string) error {
	return nil
}

func (this *LocalCache) Decr(key string) error {
	return nil
}

func (this *LocalCache) IsExist(key string) bool {
	if _, ok := this.data[key]; ok {
		return true
	}
	return false
}

func (this *LocalCache) ClearAll() error {
	return nil
}

func (this *LocalCache) StartAndGC(config string) error {
	return nil
}

func NewCache() cache.Cache {
	obj := &LocalCache{}
	obj.data = make(map[string]interface{})
	return obj
}

func Instance() cache.Cache {
	if !_init {
		_instance, _ = cache.NewCache("LocalCache", "")
		_init = true
	}
	return _instance
}