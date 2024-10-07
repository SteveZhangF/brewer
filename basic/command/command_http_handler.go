package command

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	// "strings"
	"time"

	"github.com/SteveZhangF/brewer/basic"
	"github.com/SteveZhangF/brewer/basic/errors"
	"github.com/SteveZhangF/brewer/basic/model"
	"github.com/gorilla/mux"
	// "github.com/gurufocus/gfmicro-common/plugins/user/cli"
)

type CommandHTTPHandlerOptions struct {
	RequestWriter bool
	Auth          bool
	RateLimit     int
	UserLimit     string
	Log           bool
	Valid         bool
	// Translate     bool
	Cache   time.Duration
	Methods []string
	Path    string
	REST    string
	Audit   bool

	DirectWrite bool
}

type CommandHTTPHandler struct {
	CommandCreator func(*http.Request) Command
	CommandHTTPHandlerOptions
}

func (handler *CommandHTTPHandler) Register(router *mux.Router) {
	if handler.REST != "" {
		sub := router.PathPrefix("/" + handler.REST).Subrouter()
		sub.Handle("/search", handler).Methods("POST", "OPTIONS")
		sub.Handle("/{id}", handler).Methods("GET", "OPTIONS")
		sub.Handle("", handler).Methods("POST", "OPTIONS")
		sub.Handle("/{id}", handler).Methods("PUT", "OPTIONS")
		sub.Handle("/{id}", handler).Methods("DELETE", "OPTIONS")
	} else {
		router.Handle(handler.Path, handler).Methods(handler.Methods...)
	}
}

func (handler *CommandHTTPHandler) NewCommand(w http.ResponseWriter, r *http.Request) *CommandWrapper {
	cmd := NewHttpCommand(r, handler.CommandCreator)
	if handler.Valid {
		cmd.Valid()
	}
	if handler.Auth {
		cmd.Auth()
	}
	if handler.Log {
		cmd.Log()
	}

	if handler.Audit {
		cmd.Audit(r)
	}

	if handler.REST != "" {
		var action string
		searchParameter := &basic.CommonParameter{}
		switch r.Method {
		case "POST":
			if r.URL.Path == "/"+handler.REST+"/search" {
				action = "search"
				json.NewDecoder(r.Body).Decode(searchParameter)
				break
			}
			if r.URL.Path == "/"+handler.REST {
				action = "create"
			}
		case "GET":
			action = "get"
		case "PUT":
			action = "update"
		case "DELETE":
			action = "delete"
		}
		idStr := mux.Vars(r)["id"]
		idInt, _ := strconv.ParseInt(idStr, 10, 64)
		data := map[string]interface{}{}
		json.NewDecoder(r.Body).Decode(&data)

		code := mux.Vars(r)["id"]

		cmd.REST(action, searchParameter, data, uint(idInt), code)
	}
	return cmd
}

func (handler *CommandHTTPHandler) userFromRequest(r *http.Request) *model.User {
	token := r.Header.Get("X-Token")
	if token == "" {
		return &model.User{
			IsGuest: true,
		}
	}
	user, err := model.UserByToken(token)
	if err != nil {
		fmt.Println(err)
		return &model.User{
			IsGuest: true,
		}
	}
	return user
}

func (handler *CommandHTTPHandler) doServe(w http.ResponseWriter, r *http.Request) {
	// defer func() {
	// 	r := recover()
	// 	if r != nil {
	// 		Error(w, errors.InternalErrorHappened.Error(fmt.Errorf("error while handling request %v ", r)))
	// 	}
	// }()
	// u := gfcontext.User(r.Context())
	// ctx := r.Context()

	user := handler.userFromRequest(r)

	cmd := handler.NewCommand(w, r)
	cmd.SetRequestResponse(w, r)
	result, e := cmd.Execute(user)
	if handler.DirectWrite {
		if e != nil {
			w.Write([]byte(fmt.Sprintf("%v", e)))
			// } else {
			// 	w.Write([]byte(fmt.Sprintf("%v", result)))
		}
		return
	}
	if e != nil {
		Error(w, errors.Error(e))
		return
	}
	Success(w, result)
}

func (handler *CommandHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var httpHandler = http.HandlerFunc(handler.doServe)
	httpHandler.ServeHTTP(w, r)
}

func NewHttpHandler(creator func(*http.Request) Command, option CommandHTTPHandlerOptions) *CommandHTTPHandler {
	handler := &CommandHTTPHandler{
		CommandCreator: creator,
	}
	handler.CommandHTTPHandlerOptions = option
	return handler
}
