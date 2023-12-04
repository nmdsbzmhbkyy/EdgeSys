package cache

import (
	"time"

	"mod.miligc.com/edge-common/CommonKit/cache"
	"mod.miligc.com/edge-common/CommonKit/rediscli"
)

var RedisDb *rediscli.RedisDB
var PanelCache = cache.NewTimedCache(cache.NoExpiration, 600*time.Second)
