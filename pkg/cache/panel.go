package cache

import (
	"time"

	"mod.miligc.com/edge-common/CommonKit/cache"
)

var PanelCache = cache.NewTimedCache(cache.NoExpiration, 600*time.Second)
