package global

import (
	"EdgeSys/pkg/config"
	"EdgeSys/pkg/events"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Log  *logrus.Logger // 日志
	Db   *gorm.DB       // gorm
	Conf *config.Config
)
var EventEmitter = events.EventEmitter{}
