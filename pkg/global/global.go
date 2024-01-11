package global

import (
	"EdgeSys/pkg/config"
	"gorm.io/gorm"
)

var (
	Conf   *config.Config
	MiliDb *gorm.DB
)
