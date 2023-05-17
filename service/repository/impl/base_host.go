package impl

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/iikmaulana/service-terminal/lib"
	"github.com/iikmaulana/service-terminal/models"
	"github.com/iikmaulana/service-terminal/service"
	"github.com/iikmaulana/service-terminal/service/repository/query"
	"github.com/jmoiron/sqlx"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
	"reflect"
	"strings"
)

type hostRepository struct {
	DB *sqlx.DB
}

func NewHostRepository(db *sqlx.DB) service.HostRepo {
	return hostRepository{DB: db}
}

func (u hostRepository) HostCreateRepo(form models.HostCreateRequest) (result string, serr serror.SError) {
	err := u.DB.QueryRow(query.QueryHostInsert, uuid.New().String(), form.Name, form.Url, form.Port, form.HostUsername, form.HostPassword, form.HostClientId, form.HostType, form.Topic).Scan(&result)
	if err != nil {
		return result, serror.New(err.Error())
	}
	return result, nil
}

func (u hostRepository) HostUpdateRepo(form models.HostUpdateRequest) (result string, serr serror.SError) {
	var dynamicUpdate []string
	n := 0
	x := reflect.ValueOf(form)
	num := x.NumField()
	if num <= 0 {
		return result, serr
	}
	for i := 0; i < num; i++ {
		coloumn := x.Type().Field(i).Tag.Get("db")
		t := x.Type().Field(i).Type
		exist := x.Field(i).Interface() != reflect.Zero(t).Interface()
		if exist {
			v := fmt.Sprint(x.Field(i).Interface())
			n = n + 1
			q := coloumn + ` = ` + `'` + v + `'`

			dynamicUpdate = append(dynamicUpdate, q)
		}
	}

	prefix := ` WHERE id = $1`
	queryHostID := prefix
	prefix = `UPDATE db_terminal.host SET `
	queryUpdateHost := prefix + strings.Join(dynamicUpdate, ",") + queryHostID + " returning id"

	err := u.DB.QueryRow(queryUpdateHost, form.Id).Scan(&result)

	if err != nil {
		return result, serror.New(err.Error())
	}

	return result, nil
}

func (u hostRepository) HostListRepo(form models.FilterParams) (result models.HostListResult, serr serror.SError) {

	var totalData int64
	var offset int
	var paginate models.FormMetaData

	tmpQuery := query.QueryHostList

	queryTotalData := fmt.Sprintf(`with x as (%v) select count(*) as total_data from x`, tmpQuery)
	if err := u.DB.QueryRow(queryTotalData).Scan(&totalData); err != nil {
		return result, serror.NewFromError(err)
	}

	paginate, offset = lib.Paginate(totalData, int(form.Page), int(form.Limit))

	if form.SortType != "" {
		tmpQuery = tmpQuery + fmt.Sprintf(" order by %v limit %v offset %v", form.SortType, form.Limit, offset)
	} else {
		tmpQuery = tmpQuery + fmt.Sprintf(" limit %v offset %v ", form.Limit, offset)
	}

	rows, err := u.DB.Queryx(tmpQuery)
	if err != nil {
		return result, serror.New(err.Error())
	}

	defer rows.Close()

	tmpData := []models.HostResult{}
	for rows.Next() {
		tmpResult := models.HostResult{}
		if err := rows.StructScan(&tmpResult); err != nil {
			return result, serror.New("Failed scan for struct models")
		}
		tmpData = append(tmpData, tmpResult)
	}

	result.Data = tmpData
	result.TotalData = paginate.TotalData
	result.TotalDataPerpage = int64(paginate.TotalDataPerpage)
	result.TotalPage = int64(paginate.TotalPage)
	result.From = int64(paginate.From)
	result.To = int64(paginate.To)

	return result, nil
}

func (u hostRepository) HostViewRepo(id string) (result models.HostResult, serr serror.SError) {

	tmpQuery := query.QueryHostView
	rows, err := u.DB.Queryx(tmpQuery, id)
	if err != nil {
		return result, serror.New(err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		if err := rows.StructScan(&result); err != nil {
			return result, serror.New("Failed scan for struct models")
		}
	}

	return result, nil
}

func (u hostRepository) HostDeleteRepo(id string) (result string, serr serror.SError) {

	tmpQuery := query.QueryHostDelete
	rows, err := u.DB.Queryx(tmpQuery, id)
	if err != nil {
		return result, serror.New(err.Error())
	}

	defer rows.Close()
	return result, nil
}
