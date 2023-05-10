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
	"github.com/uzzeet/uzzeet-gateway/libs/utils/uttime"
	"reflect"
	"strings"
)

type userRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) service.UserRepo {
	return userRepository{DB: db}
}

func (u userRepository) UserCreateRepo(form models.UserCreateRequest) (result string, serr serror.SError) {

	err := u.DB.QueryRow(query.QueryUserInsert, uuid.New().String(), form.Username, form.Password, uttime.Now().Format("2006-01-02 15:04:05"), form.Status, form.LastLogin).Scan(&result)
	if err != nil {
		return result, serror.New(err.Error())
	}
	return result, nil
}

func (u userRepository) UserUpdateRepo(form models.UserUpdateRequest) (result string, serr serror.SError) {
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
	queryUserID := prefix
	prefix = `UPDATE db_terminal.user SET `
	queryUpdateUser := prefix + strings.Join(dynamicUpdate, ",") + queryUserID + " returning id"

	err := u.DB.QueryRow(queryUpdateUser, form.Id).Scan(&result)

	if err != nil {
		return result, serror.New(err.Error())
	}

	return result, nil
}

func (u userRepository) UserListRepo(form models.FilterParams) (result models.UserListResult, serr serror.SError) {

	var totalData int64
	var offset int
	var paginate models.FormMetaData

	tmpQuery := query.QueryUserList

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

	tmpData := []models.UserViewResult{}
	for rows.Next() {
		tmpResult := models.UserViewResult{}
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

func (u userRepository) UserViewRepo(id string) (result models.UserResult, serr serror.SError) {

	tmpQuery := query.QueryUserView
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

func (u userRepository) UserDeleteRepo(id string) (result string, serr serror.SError) {

	tmpQuery := query.QueryUserDelete
	rows, err := u.DB.Queryx(tmpQuery, id)
	if err != nil {
		return result, serror.New(err.Error())
	}

	defer rows.Close()
	return result, nil
}

func (u userRepository) UserViewByUseranmeRepo(username string) (result models.UserResult, serr serror.SError) {
	tmpQuery := query.QueryUserLogin
	rows, err := u.DB.Queryx(tmpQuery, username)
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
