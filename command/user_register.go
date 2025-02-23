package command

import (
	"encoding/json"
	"net/http"

	"github.com/SteveZhangF/brewer/basic/command"
	"github.com/SteveZhangF/brewer/basic/model"
)

type UserRegisterCommand struct {
	command.DefaultCommand
	Data model.User
}

func (dc *UserRegisterCommand) NeedPermission() int {
	return 0
}

func (dc *UserRegisterCommand) NeedLogin() bool {
	return true
}

func (dc *UserRegisterCommand) CacheKey() string {
	return ""
}

func (dc *UserRegisterCommand) Valid() error {
	return nil
}

func (dc *UserRegisterCommand) Name() string {
	return "UserRegisterCommand"
}

func (dc *UserRegisterCommand) Execute(u *model.User) (interface{}, error) {
	dc.Data.HashPassword(dc.Data.Password)
	err := dc.Data.Create()
	if err != nil {
		return nil, err
	}
	return dc.Data, nil
}

func (dc *UserRegisterCommand) Status(u *model.User) int64 {
	return 0
}

func (dc *UserRegisterCommand) Region(u *model.User) string {
	return ""
}

func HTTPUserRegister(r *http.Request) command.Command {
	cmd := &UserRegisterCommand{}
	data := model.User{}
	json.NewDecoder(r.Body).Decode(&data)
	return cmd
}
