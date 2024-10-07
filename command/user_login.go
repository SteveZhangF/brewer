package command

import (
	"encoding/json"
	"net/http"

	"github.com/SteveZhangF/brewer/basic/command"
	"github.com/SteveZhangF/brewer/basic/model"
)

type UserLoginCommand struct {
	command.DefaultCommand
	Username string `json:"username"`
	Password string `json:"password"`
}

func (dc *UserLoginCommand) NeedPermission() int {
	return 0
}

func (dc *UserLoginCommand) NeedLogin() bool {
	return true
}

func (dc *UserLoginCommand) CacheKey() string {
	return ""
}

func (dc *UserLoginCommand) Valid() error {
	return nil
}

func (dc *UserLoginCommand) Name() string {
	return "UserLoginCommand"
}

func (dc *UserLoginCommand) Execute(u *model.User) (interface{}, error) {
	store := &model.User{}
	return store.Login(dc.Username, dc.Password)
}

func (dc *UserLoginCommand) Status(u *model.User) int64 {
	return 0
}

func (dc *UserLoginCommand) Region(u *model.User) string {
	return ""
}

func HTTPUserLogin(r *http.Request) command.Command {
	cmd := &UserLoginCommand{}
	json.NewDecoder(r.Body).Decode(cmd)
	return cmd
}
