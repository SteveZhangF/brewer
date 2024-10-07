package command

import (
	"github.com/SteveZhangF/brewer/basic/errors"
	"github.com/SteveZhangF/brewer/basic/model"
)

// "github.com/gurufocus/gfmicro-common/errors"
// model "github.com/gurufocus/gfmicro/basic/http"

type AuthCommand struct {
	ProxyCommand
}

func (au *AuthCommand) checkPermission(u *model.User) error {
	if au.RealCommand().NeedLogin() && u.IsGuest {
		return errors.AuthNeedLogin.Error()
	}

	needPermission := au.Command.NeedPermission()
	if needPermission == 0 {
		return nil
	}
	if u.Level < needPermission {
		return errors.NoPermission.Error()
	}

	return nil
}

func (au *AuthCommand) Execute(u *model.User) (interface{}, error) {
	if permissionError := au.checkPermission(u); permissionError != nil {
		return nil, permissionError
	}
	return au.Command.Execute(u)
}

func WithAuth(c Command) Command {
	au := &AuthCommand{}
	au.Command = c
	return au
}
