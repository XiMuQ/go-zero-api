package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"go-zero-api/api/code/internal/config"
	"go-zero-api/api/code/internal/middleware"
)

type ServiceContext struct {
	Config config.Config
	Check  rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Check:  middleware.NewCheckMiddleware().Handle,
	}
}
