package command

import (
	model "kenneth/backend/model"

	"github.com/sirupsen/logrus"
)

// "github.com/gurufocus/gfmicro-common/errors"
// model "github.com/gurufocus/gfmicro/basic/http"

type LogCommand struct {
	ProxyCommand
}

func (au *LogCommand) Execute(u *model.User) (interface{}, error) {
	logrus.Info("LogCommand")
	return au.Command.Execute(u)
}

func WithLog(c Command) Command {
	au := &LogCommand{}
	au.Command = c
	return au
}
