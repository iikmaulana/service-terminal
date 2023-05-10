package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/iikmaulana/service-terminal/lib"
	"github.com/iikmaulana/service-terminal/models"
	"github.com/uzzeet/uzzeet-gateway/libs/helper"
	"net/http"
	"strings"
)

func (ox gatewayHandler) HostCreateRequest(ctx *gin.Context) {
	_, serr := lib.ClaimToken(ctx.Request.Header["Authorization"])
	if serr != nil {
		lib.Response(http.StatusUnauthorized, serr.Error(), "", "", serr.File(), "", "", ctx)
		return
	}

	form := models.HostCreateRequest{
		Name: ctx.PostForm("name"),
		Url:  ctx.PostForm("url"),
		Port: ctx.PostForm("port"),
	}

	tmpResult, err := ox.hostUsecase.HostCreateUsecase(form)
	if err != nil {
		lib.Response(http.StatusInternalServerError, err.Error(), "", "", "HostCreateUsecase", "add", "", ctx)
		return
	}

	lib.Response(http.StatusOK, "", "", "", "HostCreateUsecase", "add", tmpResult, ctx)
	return
}

func (ox gatewayHandler) HostUpdateRequest(ctx *gin.Context) {
	_, serr := lib.ClaimToken(ctx.Request.Header["Authorization"])
	if serr != nil {
		lib.Response(http.StatusUnauthorized, serr.Error(), "", "", serr.File(), "", "", ctx)
		return
	}

	form := models.HostUpdateRequest{
		Id:   ctx.PostForm("id"),
		Name: ctx.PostForm("name"),
		Url:  ctx.PostForm("url"),
		Port: ctx.PostForm("port"),
	}

	tmpResult, err := ox.hostUsecase.HostUpdateUsecase(form)
	if err != nil {
		lib.Response(http.StatusInternalServerError, err.Error(), "", "", "HostUpdateUsecase", "add", "", ctx)
		return
	}

	lib.Response(http.StatusOK, "", "", "", "HostUpdateUsecase", "add", tmpResult, ctx)
	return
}

func (ox gatewayHandler) HostListRequest(ctx *gin.Context) {
	_, serr := lib.ClaimToken(ctx.Request.Header["Authorization"])
	if serr != nil {
		lib.Response(http.StatusUnauthorized, serr.Error(), "", "", serr.File(), "", "", ctx)
		return
	}

	tmpFilter := strings.Split(ctx.Query("filter"), ",")
	tmpSearch := strings.Split(ctx.Query("search"), ",")

	form := models.FilterParams{
		Page:     helper.StringToInt(ctx.Query("page"), 1),
		Limit:    helper.StringToInt(ctx.Query("limit"), 10),
		SortType: ctx.Query("sort_by"),
		Filter:   tmpFilter,
		Search:   tmpSearch,
	}

	tmpResult, err := ox.hostUsecase.HostListUsecase(form)
	if err != nil {
		lib.Response(http.StatusInternalServerError, err.Error(), "", "", "HostListUsecase", "list", "", ctx)
		return
	}

	lib.Response(http.StatusOK, "", "", "", "HostListUsecase", "list", tmpResult, ctx)
	return
}

func (ox gatewayHandler) HostViewRequest(ctx *gin.Context) {
	_, serr := lib.ClaimToken(ctx.Request.Header["Authorization"])
	if serr != nil {
		lib.Response(http.StatusUnauthorized, serr.Error(), "", "", serr.File(), "", "", ctx)
		return
	}

	id := ctx.Query("id")

	tmpResult, err := ox.hostUsecase.HostViewUsecase(id)
	if err != nil {
		lib.Response(http.StatusInternalServerError, err.Error(), "", "", "HostViewUsecase", "view", "", ctx)
		return
	}

	if tmpResult.Id != "" {
		lib.Response(http.StatusOK, "", "", "", "HostViewUsecase", "view", tmpResult, ctx)
		return
	} else {
		lib.Response(http.StatusOK, "", "", "", "HostViewUsecase", "view", "", ctx)
		return
	}
}

func (ox gatewayHandler) HostDeleteRequest(ctx *gin.Context) {
	_, serr := lib.ClaimToken(ctx.Request.Header["Authorization"])
	if serr != nil {
		lib.Response(http.StatusUnauthorized, serr.Error(), "", "", serr.File(), "", "", ctx)
		return
	}

	id := ctx.Query("id")

	tmpResult, err := ox.hostUsecase.HostDeleteUsecase(id)
	if err != nil {
		lib.Response(http.StatusInternalServerError, err.Error(), "", "", "HostDeleteUsecase", "delete", "", ctx)
		return
	}

	lib.Response(http.StatusOK, "", "", "", "HostDeleteUsecase", "delete", tmpResult, ctx)
	return
}
