package service

import (
	"github.com/iikmaulana/service-terminal/models"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

type UserRepo interface {
	UserCreateRepo(form models.UserCreateRequest) (result string, serr serror.SError)
	UserUpdateRepo(form models.UserUpdateRequest) (result string, serr serror.SError)
	UserListRepo(form models.FilterParams) (result models.UserListResult, serr serror.SError)
	UserViewRepo(id string) (result models.UserResult, serr serror.SError)
	UserDeleteRepo(id string) (result string, serr serror.SError)
	UserViewByUseranmeRepo(username string) (result models.UserResult, serr serror.SError)
}
type HostRepo interface {
	HostCreateRepo(form models.HostCreateRequest) (result string, serr serror.SError)
	HostUpdateRepo(form models.HostUpdateRequest) (result string, serr serror.SError)
	HostListRepo(form models.FilterParams) (result models.HostListResult, serr serror.SError)
	HostViewRepo(id string) (result models.HostResult, serr serror.SError)
	HostDeleteRepo(id string) (result string, serr serror.SError)
}
