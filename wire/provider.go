package wire

import (
	"github.com/Guohuixixi/go-easy/bootstrap/internal"
	"github.com/Guohuixixi/go-easy/core"
	"github.com/google/wire"
)

var BaseProvider = wire.NewSet(
	internal.NewViper,
	internal.NewConfig,
	internal.NewMysql,
	internal.NewRedis,
	internal.NewZap,
	core.NewApplication,
)

var WebProvider = wire.NewSet(
	internal.NewServer,
)
