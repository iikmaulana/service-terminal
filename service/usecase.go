package service

import (
	"github.com/iikmaulana/service-terminal/models"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

type UserUsageUsecase interface {
	UserCreateUsecase(form models.UserCreateRequest) (result string, serr serror.SError)
	UserUpdateUsecase(form models.UserUpdateRequest) (result string, serr serror.SError)
	UserListUsecase(form models.FilterParams) (result models.UserListResult, serr serror.SError)
	UserViewUsecase(id string) (result models.UserViewResult, serr serror.SError)
	UserDeleteUsecase(id string) (result string, serr serror.SError)
	UserLoginUsecase(username string, password string) (result string, serr serror.SError)
}

type HostUsageUsecase interface {
	HostCreateUsecase(form models.HostCreateRequest) (result string, serr serror.SError)
	HostUpdateUsecase(form models.HostUpdateRequest) (result string, serr serror.SError)
	HostListUsecase(form models.FilterParams) (result models.HostListResult, serr serror.SError)
	HostViewUsecase(id string) (result models.HostResult, serr serror.SError)
	HostDeleteUsecase(id string) (result string, serr serror.SError)
}
