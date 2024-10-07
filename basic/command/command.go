package command

import (
	"kenneth/backend/basic"
	model "kenneth/backend/model"
	"net/http"
)

type Command interface {
	Execute(*model.User) (interface{}, error)
	Valid() error
	Name() string
	CacheKey() string
	NeedPermission() int
	NeedLogin() bool
	REST(action string, searchParameter *basic.CommonParameter, data map[string]interface{}, id uint, code string)
	SetRequestResponse(w http.ResponseWriter, r *http.Request)
}

type CodeRESTCommand struct {
	Action          string
	Data            map[string]interface{}
	Code            string
	SearchParameter *basic.CommonParameter
	DefaultCommand
}

func (restCommand *CodeRESTCommand) RESTCommandAction() string {
	return restCommand.Action
}

type IsRestCommand interface {
	RESTCommandAction() string
}

func (rc *CodeRESTCommand) REST(action string, searchParameter *basic.CommonParameter, data map[string]interface{}, id uint, code string) {
	rc.Action = action
	rc.SearchParameter = searchParameter
	rc.Data = data

	rc.Code = code
}

type RESTCommand struct {
	Action          string
	Data            map[string]interface{}
	SearchParameter *basic.CommonParameter
	ID              uint
	DefaultCommand
}

func (restCommand *RESTCommand) RESTCommandAction() string {
	return restCommand.Action
}

func (rc *RESTCommand) REST(action string, searchParameter *basic.CommonParameter, data map[string]interface{}, id uint, code string) {
	rc.Action = action
	rc.SearchParameter = searchParameter
	rc.Data = data
	rc.ID = id
}

type DefaultCommand struct {
	Request  *http.Request       `json:"-"`
	Response http.ResponseWriter `json:"-"`
}

func (dc *DefaultCommand) SetRequestResponse(w http.ResponseWriter, r *http.Request) {
	dc.Request = r
	dc.Response = w
}

func (dc *DefaultCommand) Valid() error {
	return nil
}

func (dc *DefaultCommand) REST(string, *basic.CommonParameter, map[string]interface{}, uint, string) {

}

func (dc *DefaultCommand) NeedPermission() int {
	return 0
}

func (dc *DefaultCommand) NeedLogin() bool {
	return false
}
