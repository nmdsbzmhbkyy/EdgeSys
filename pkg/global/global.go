package global

import (
	"EdgeSys/pkg/config"
	"EdgeSys/pkg/events"
	"gorm.io/gorm"
)

var (
	Db     *gorm.DB // gorm
	Conf   *config.Config
	MiliDb *gorm.DB
)
var EventEmitter = events.EventEmitter{}
