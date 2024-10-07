package command

import (
	"encoding/json"
	"net/http"

	"github.com/SteveZhangF/brewer/basic/command"
	"github.com/SteveZhangF/brewer/basic/model"
)

type UserSetPassword struct {
	command.DefaultCommand
	Token       string `json:"token"`
	OldPassword string `json:"old_password"`
	Password    string `json:"password"`
}

func (dc *UserSetPassword) NeedPermission() int {
	return 0
}

func (dc *UserSetPassword) NeedLogin() bool {
	return true
}

func (dc *UserSetPassword) CacheKey() string {
	return ""
}

func (dc *UserSetPassword) Valid() error {
	return nil
}

func (dc *UserSetPassword) Name() string {
	return "UserSetPassword"
}

func (dc *UserSetPassword) Execute(u *model.User) (interface{}, error) {
	err := u.SetPassword(dc.OldPassword, dc.Password)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (dc *UserSetPassword) Status(u *model.User) int64 {
	return 0
}

func (dc *UserSetPassword) Region(u *model.User) string {
	return ""
}

func HTTPUserSetPassword(r *http.Request) command.Command {
	cmd := &UserSetPassword{}
	json.NewDecoder(r.Body).Decode(cmd)
	return cmd
}
