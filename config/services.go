package config

import (
	"github.com/iikmaulana/service-terminal/controller"
	"github.com/iikmaulana/service-terminal/service/handler"
	"github.com/iikmaulana/service-terminal/service/repository/impl"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

func (cfg Config) InitService() serror.SError {

	userRepo := impl.NewUserRepository(cfg.DB)
	hostRepo := impl.NewHostRepository(cfg.DB)

	userUsecase := controller.NewUserUsageUsecase(userRepo)
	hostUsecase := controller.NewHostUsageUsecase(hostRepo)

	handler.NewHandler(cfg.Gin, userUsecase, hostUsecase)
	return nil
}
