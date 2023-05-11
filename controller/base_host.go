package controller

import (
	"github.com/iikmaulana/service-terminal/models"
	"github.com/iikmaulana/service-terminal/service"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

type hostUsecase struct {
	hostRepo service.HostRepo
}

func NewHostUsageUsecase(hostRepo service.HostRepo) service.HostUsageUsecase {
	return hostUsecase{hostRepo: hostRepo}
}

func (u hostUsecase) HostCreateUsecase(form models.HostCreateRequest) (result string, serr serror.SError) {
	if form.HostType == "" {
		return result, serror.New("Host type tidak boleh kosong")
	}

	tmpResult, err := u.hostRepo.HostCreateRepo(form)
	if err != nil {
		return result, err
	}

	result = tmpResult
	return result, nil
}

func (u hostUsecase) HostUpdateUsecase(form models.HostUpdateRequest) (result string, serr serror.SError) {
	if form.HostType == "" {
		return result, serror.New("Host type tidak boleh kosong")
	}
	tmpResult, err := u.hostRepo.HostUpdateRepo(form)
	if err != nil {
		return result, err
	}

	result = tmpResult
	return result, nil
}

func (u hostUsecase) HostListUsecase(form models.FilterParams) (result models.HostListResult, serr serror.SError) {
	tmpResult, err := u.hostRepo.HostListRepo(form)
	if err != nil {
		return result, err
	}

	result = tmpResult
	return result, nil
}

func (u hostUsecase) HostViewUsecase(id string) (result models.HostResult, serr serror.SError) {
	tmpResult, err := u.hostRepo.HostViewRepo(id)
	if err != nil {
		return result, err
	}

	result = tmpResult
	return result, nil
}

func (u hostUsecase) HostDeleteUsecase(id string) (result string, serr serror.SError) {
	_, err := u.hostRepo.HostDeleteRepo(id)
	if err != nil {
		return result, err
	}

	return result, nil
}
