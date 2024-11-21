package wire

import (
	"github.com/Guohuixixi/go-easy/core"
	"github.com/Guohuixixi/go-easy/core/bootstrap"
	"github.com/Guohuixixi/go-easy/core/logger"
	"github.com/google/wire"
)

var BaseProvider = wire.NewSet(
	bootstrap.NewViper,
	bootstrap.NewConfig,
	bootstrap.NewMysql,
	bootstrap.NewRedis,
	bootstrap.NewZap,
	logger.NewZapLogger,
	core.NewApplication,
)

var WebProvider = wire.NewSet(
	bootstrap.NewServer,
)
