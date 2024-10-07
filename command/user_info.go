package command

import (
	"net/http"

	"github.com/SteveZhangF/brewer/basic/command"
	"github.com/SteveZhangF/brewer/basic/model"
)

type UserInfoCommand struct {
	command.DefaultCommand
	Token string
}

func (dc *UserInfoCommand) NeedPermission() int {
	return 0
}

func (dc *UserInfoCommand) NeedLogin() bool {
	return true
}

func (dc *UserInfoCommand) CacheKey() string {
	return ""
}

func (dc *UserInfoCommand) Valid() error {
	return nil
}

func (dc *UserInfoCommand) Name() string {
	return "UserInfoCommand"
}

func (dc *UserInfoCommand) Execute(u *model.User) (interface{}, error) {
	return model.UserByToken(dc.Token)
}

func (dc *UserInfoCommand) Status(u *model.User) int64 {
	return 0
}

func (dc *UserInfoCommand) Region(u *model.User) string {
	return ""
}

func HTTPUserInfo(r *http.Request) command.Command {
	cmd := &UserInfoCommand{}
	token := r.URL.Query().Get("token")
	cmd.Token = token
	return cmd
}
