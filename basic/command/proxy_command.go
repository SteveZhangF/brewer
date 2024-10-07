package command

import (
	"encoding/json"
	"net/http"

	"github.com/SteveZhangF/brewer/basic"
)

type Proxy interface {
	RealCommand() Command
}

// Execute(*model.User) (interface{}, error)
// 	Valid() error
// 	Name() string
// 	CacheKey() string
// 	NeedPermission() string
// 	NeedLogin() bool

type ProxyCommand struct {
	// DumyCommand
	Command Command
	DefaultCommand
}

func (pc *ProxyCommand) CacheKey() string {
	return ""
}

// func (pc *ProxyCommand) IsProxy() bool {
// 	return true
// }

func (pc *ProxyCommand) NeedLogin() bool {
	return pc.Command.NeedLogin()
}

func (pc *ProxyCommand) NeedPermission() int {
	return pc.Command.NeedPermission()
}

func (pc *ProxyCommand) Valid() error {
	return pc.Command.Valid()
}

func (pc *ProxyCommand) Name() string {
	return pc.Command.Name()
}

func (pc *ProxyCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(pc.RealCommand())
}

func (pc *ProxyCommand) RealCommand() Command {
	if px, ok := pc.Command.(Proxy); ok {
		return px.RealCommand()
	}
	return pc.Command
}

func (dc *ProxyCommand) REST(action string, searchParameter *basic.CommonParameter, data map[string]interface{}, id uint, code string) {
	dc.RealCommand().REST(action, searchParameter, data, id, code)
}

func (dc *ProxyCommand) SetRequestResponse(w http.ResponseWriter, r *http.Request) {
	dc.RealCommand().SetRequestResponse(w, r)
	dc.DefaultCommand.SetRequestResponse(w, r)
}
