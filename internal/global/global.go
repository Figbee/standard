package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/gobuffalo/packr/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Conf    Configuration
	ConfBox *packr.Box
	Logger  *zap.SugaredLogger
	Orm     *gorm.DB
	Casbin  *casbin.Enforcer
)
