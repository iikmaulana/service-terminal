package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/iikmaulana/service-terminal/lib"
	"github.com/iikmaulana/service-terminal/models"
	"github.com/uzzeet/uzzeet-gateway/libs/helper"
	"net/http"
	"strings"
)

func (ox gatewayHandler) UserCreateRequest(ctx *gin.Context) {
	_, serr := lib.ClaimToken(ctx.Request.Header["Authorization"])
	if serr != nil {
		lib.Response(http.StatusUnauthorized, serr.Error(), "", "", serr.File(), "", "", ctx)
		return
	}

	form := models.UserCreateRequest{
		Username: ctx.PostForm("username"),
		Password: ctx.PostForm("password"),
		Status:   int(helper.StringToInt(ctx.PostForm("status"), 0)),
	}

	tmpResult, err := ox.userUsecase.UserCreateUsecase(form)
	if err != nil {
		lib.Response(http.StatusInternalServerError, err.Error(), "", "", "UserCreateUsecase", "add", "", ctx)
		return
	}

	lib.Response(http.StatusOK, "", "", "", "UserCreateUsecase", "add", tmpResult, ctx)
	return
}

func (ox gatewayHandler) UserUpdateRequest(ctx *gin.Context) {
	_, serr := lib.ClaimToken(ctx.Request.Header["Authorization"])
	if serr != nil {
		lib.Response(http.StatusUnauthorized, serr.Error(), "", "", serr.File(), "", "", ctx)
		return
	}

	form := models.UserUpdateRequest{
		Id:       ctx.PostForm("id"),
		Username: ctx.PostForm("username"),
		Password: ctx.PostForm("password"),
		Status:   int(helper.StringToInt(ctx.PostForm("status"), 0)),
	}

	tmpResult, err := ox.userUsecase.UserUpdateUsecase(form)
	if err != nil {
		lib.Response(http.StatusInternalServerError, err.Error(), "", "", "UserUpdateUsecase", "add", "", ctx)
		return
	}

	lib.Response(http.StatusOK, "", "", "", "UserUpdateUsecase", "add", tmpResult, ctx)
	return
}

func (ox gatewayHandler) UserListRequest(ctx *gin.Context) {
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

	tmpResult, err := ox.userUsecase.UserListUsecase(form)
	if err != nil {
		lib.Response(http.StatusInternalServerError, err.Error(), "", "", "UserListUsecase", "list", "", ctx)
		return
	}

	lib.Response(http.StatusOK, "", "", "", "UserListUsecase", "list", tmpResult, ctx)
	return
}

func (ox gatewayHandler) UserViewRequest(ctx *gin.Context) {
	_, serr := lib.ClaimToken(ctx.Request.Header["Authorization"])
	if serr != nil {
		lib.Response(http.StatusUnauthorized, serr.Error(), "", "", serr.File(), "", "", ctx)
		return
	}

	id := ctx.Query("id")

	tmpResult, err := ox.userUsecase.UserViewUsecase(id)
	if err != nil {
		lib.Response(http.StatusInternalServerError, err.Error(), "", "", "UserViewUsecase", "view", "", ctx)
		return
	}

	if tmpResult.Id != "" {
		lib.Response(http.StatusOK, "", "", "", "UserViewUsecase", "view", tmpResult, ctx)
		return
	} else {
		lib.Response(http.StatusOK, "", "", "", "UserViewUsecase", "view", "", ctx)
		return
	}
}

func (ox gatewayHandler) UserLoginRequest(ctx *gin.Context) {

	tmpUsername := ctx.Query("username")
	tmpPassword := ctx.Query("password")

	tmpResult, err := ox.userUsecase.UserLoginUsecase(tmpUsername, tmpPassword)
	if err != nil {
		lib.Response(http.StatusInternalServerError, err.Error(), "", "", "UserLoginUsecase", "get", "", ctx)
		return
	}

	lib.Response(http.StatusOK, "", "", "", "UserLoginUsecase", "get", tmpResult, ctx)
	return
}

func (ox gatewayHandler) UserDeleteRequest(ctx *gin.Context) {
	_, serr := lib.ClaimToken(ctx.Request.Header["Authorization"])
	if serr != nil {
		lib.Response(http.StatusUnauthorized, serr.Error(), "", "", serr.File(), "", "", ctx)
		return
	}

	id := ctx.Query("id")

	tmpResult, err := ox.userUsecase.UserDeleteUsecase(id)
	if err != nil {
		lib.Response(http.StatusInternalServerError, err.Error(), "", "", "UserDeleteUsecase", "delete", "", ctx)
		return
	}

	lib.Response(http.StatusOK, "", "", "", "UserDeleteUsecase", "delete", tmpResult, ctx)
	return
}
