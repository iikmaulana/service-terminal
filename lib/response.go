package lib

import (
	"github.com/gin-gonic/gin"
	"github.com/iikmaulana/service-terminal/models"
)

func Response(response int, error string, appid string, svcid string, controller string, action string, data interface{}, c *gin.Context) {

	switch action {
	case "POST":
		action = "add"
	case "GET":
		action = "view"
	case "PUT":
		action = "edit"
	case "DELETE":
		action = "delete"
	default:
		action = action
	}

	res := struct {
		Response   int         `json:"response"`
		Error      string      `json:"error"`
		Appid      string      `json:"appid"`
		Svcid      string      `json:"svcid"`
		Controller string      `json:"controller"`
		Action     string      `json:"action"`
		Result     interface{} `json:"result"`
	}{
		Response:   response,
		Error:      error,
		Appid:      appid,
		Svcid:      svcid,
		Controller: controller,
		Action:     action,
		Result:     data,
	}
	c.JSON(response, res)

}

func Paginate(totalData int64, page int, limit int) (res models.FormMetaData, offset int) {

	offset = (page * limit) - limit
	cLimit := int64(limit)

	lastPage := totalData / cLimit
	if totalData%cLimit != 0 {
		lastPage += 1
	}

	res.TotalDataPerpage = 1
	res.TotalPage = int(lastPage)
	res.TotalData = totalData
	res.To = page + 1
	res.From = page - 1

	if res.From < res.TotalDataPerpage {
		res.From = 1
	}

	if res.To > res.TotalPage {
		res.To = res.TotalPage
	}

	return res, offset
}
