package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/iikmaulana/service-terminal/service"
)

type gatewayHandler struct {
	service     *gin.Engine
	userUsecase service.UserUsageUsecase
	hostUsecase service.HostUsageUsecase
}

func NewHandler(svc *gin.Engine,
	userUsecase service.UserUsageUsecase,
	hostUsecase service.HostUsageUsecase,
) {
	h := gatewayHandler{
		service:     svc,
		userUsecase: userUsecase,
		hostUsecase: hostUsecase,
	}

	h.initRouteUsage()

}
