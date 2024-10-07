package command

import (
	"encoding/json"
	model "kenneth/backend/model"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

type AuditCommand struct {
	ProxyCommand
	Request *http.Request
}

func (au *AuditCommand) Execute(u *model.User) (interface{}, error) {
	logrus.Info("LogCommand")
	data, err := json.Marshal(au.Command)
	if err != nil {
		return nil, err
	}
	store := model.Audit{}
	audit := &model.Audit{
		UserId: u.ID,
		Data:   string(data),
		Name:   au.Command.Name(),
		Method: au.Request.Method,
		URI:    au.Request.RequestURI,
	}

	if isRestCommand, ok := au.RealCommand().(IsRestCommand); ok {
		if strings.ToLower(isRestCommand.RESTCommandAction()) == "search" || strings.ToLower(isRestCommand.RESTCommandAction()) == "get" {
			return au.Command.Execute(u)
		}
	}

	go store.Create(audit)
	return au.Command.Execute(u)
}

func WithAudit(c Command, r *http.Request) Command {
	au := &AuditCommand{}
	au.Request = r
	au.Command = c
	return au
}
