package command

import (
	"encoding/json"
	"kenneth/backend/basic/command"
	"kenneth/backend/model"
	"net/http"
)

type UserLogoutCommand struct {
	command.DefaultCommand
	Username string `json:"username"`
	Password string `json:"password"`
}

func (dc *UserLogoutCommand) NeedPermission() int {
	return 0
}

func (dc *UserLogoutCommand) NeedLogin() bool {
	return true
}

func (dc *UserLogoutCommand) CacheKey() string {
	return ""
}

func (dc *UserLogoutCommand) Valid() error {
	return nil
}

func (dc *UserLogoutCommand) Name() string {
	return "UserLogoutCommand"
}

func (dc *UserLogoutCommand) Execute(u *model.User) (interface{}, error) {
	return nil, u.Logout()
}

func (dc *UserLogoutCommand) Status(u *model.User) int64 {
	return 0
}

func (dc *UserLogoutCommand) Region(u *model.User) string {
	return ""
}

func HTTPUserLogout(r *http.Request) command.Command {
	cmd := &UserLogoutCommand{}
	json.NewDecoder(r.Body).Decode(cmd)
	return cmd
}
