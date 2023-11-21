package cache

import (
	"time"

	"github.com/PandaXGO/PandaKit/cache"
	"github.com/PandaXGO/PandaKit/rediscli"
)

var RedisDb *rediscli.RedisDB
var PanelCache = cache.NewTimedCache(cache.NoExpiration, 600*time.Second)
