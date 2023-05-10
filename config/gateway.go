package config

import (
	"github.com/gin-gonic/gin"
	"github.com/uzzeet/uzzeet-gateway/controller"
	"github.com/uzzeet/uzzeet-gateway/libs"
	"github.com/uzzeet/uzzeet-gateway/libs/helper"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/logger"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
	"os"
)

func (cfg *Config) initGateway() serror.SError {

	var err error

	cfg.Registry, err = controller.InitRegistry(controller.RegistryConfig{
		Address:  helper.Env(libs.AppRegistryAddr, "172.26.35.23:6379"),
		Password: helper.Env(libs.AppRegistryPwd, ""),
	})
	if err != nil {
		return serror.NewFromError(err)
	}

	cfg.Server, err = controller.NewServerHttp(controller.ServerConfig{
		Host:      helper.Env(libs.AppHost, "127.0.0.1"),
		Port:      int(helper.StringToInt(helper.Env(libs.AppPort, "6001"), 6001)),
		Key:       helper.Env(libs.AppKey, "fuel"),
		Name:      helper.Env(libs.AppName, "fuel"),
		Namespace: helper.Env(libs.AppNamespace, libs.NamespaceDefault),
		TypeConn:  "http",
	}, cfg.Registry)
	if err != nil {
		return serror.NewFromError(err)
	}

	cfg.Gateway = cfg.Server.AsGatewayService(os.Getenv("APP_BASEPOINT"))

	serr := cfg.Server.Write()
	if serr != nil {
		logger.Panic(serr)
	}

	cfg.Gin = gin.Default()

	return nil
}
