//go:build wireinject

package wire

import (
	"github.com/Guohuixixi/go-easy/bootstrap"
	"github.com/google/wire"
)

func InitApp() (bootstrap.Application, error) {
	wire.Build(BaseProvider, WebProvider)
	return bootstrap.Application{}, nil
}
