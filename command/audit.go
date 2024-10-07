package command

import (
	"net/http"

	"github.com/SteveZhangF/brewer/basic/command"
	"github.com/SteveZhangF/brewer/basic/model"
)

type AuditCommand struct {
	command.RESTCommand
}

func (dc *AuditCommand) NeedPermission() int {
	return 0
}

func (dc *AuditCommand) NeedLogin() bool {
	return true
}

func (dc *AuditCommand) CacheKey() string {
	return ""
}

func (dc *AuditCommand) Valid() error {
	return nil
}

func (dc *AuditCommand) Name() string {
	return "FindCommand"
}

func (dc *AuditCommand) Execute(u *model.User) (interface{}, error) {
	store := &model.Audit{}
	switch dc.Action {
	case "search":
		return store.Search(dc.SearchParameter)
	case "update":
		return store.Update(dc.ID, dc.Data)
	case "delete":
		return nil, store.Delete(dc.ID)
	}
	return store.Find(dc.ID)
}

func (dc *AuditCommand) Status(u *model.User) int64 {
	return 0
}

func (dc *AuditCommand) Region(u *model.User) string {
	return ""
}

func HTTPRESTAudit(r *http.Request) command.Command {
	cmd := &AuditCommand{}
	return cmd
}
