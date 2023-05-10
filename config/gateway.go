package config

import (
	"github.com/gin-gonic/gin"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

func (cfg *Config) initGateway() serror.SError {

	cfg.Gin = gin.Default()

	return nil
}
