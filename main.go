package main

import (
	"fmt"
	_ "kenneth/backend/basic"
	basicCommand "kenneth/backend/basic/command"
	"kenneth/backend/command"
	_ "kenneth/backend/model"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, *, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {

	// defer basic.Shutdown()

	// 初始化配置
	r := mux.NewRouter()

	// basicCommand.NewHttpHandler(command.HTTPOrdersPayment, basicCommand.CommandHTTPHandlerOptions{
	// 	Path:    "/orders/payment",
	// 	Log:     true,
	// 	Valid:   true,
	// 	Auth:    true,
	// 	Methods: []string{"POST", "OPTIONS"},
	// }).Register(r)

	// basicCommand.NewHttpHandler(command.HTTPRESTPayment, basicCommand.CommandHTTPHandlerOptions{
	// 	REST:  "payments",
	// 	Log:   true,
	// 	Valid: true,
	// 	Auth:    true,
	// }).Register(r)

	// basicCommand.NewHttpHandler(command.HTTPPaymentPay, basicCommand.CommandHTTPHandlerOptions{
	// 	Path:    "/payments/{id}/pay",
	// 	Log:     true,
	// 	Valid:   true,
	// 	Auth:    true,
	// 	Methods: []string{"POST", "OPTIONS"},
	// }).Register(r)

	basicCommand.NewHttpHandler(command.HTTPUserLogin, basicCommand.CommandHTTPHandlerOptions{
		Path:    "/login",
		Log:     true,
		Audit:   true,
		Valid:   true,
		Auth:    false,
		Methods: []string{"POST", "OPTIONS"},
	}).Register(r)

	basicCommand.NewHttpHandler(command.HTTPUserLogout, basicCommand.CommandHTTPHandlerOptions{
		Path:    "/user/logout",
		Log:     true,
		Valid:   true,
		Auth:    true,
		Audit:   true,
		Methods: []string{"POST", "OPTIONS"},
	}).Register(r)

	basicCommand.NewHttpHandler(command.HTTPUserInfo, basicCommand.CommandHTTPHandlerOptions{
		Path:    "/user/info",
		Log:     true,
		Valid:   true,
		Audit:   true,
		Auth:    true,
		Methods: []string{"GET", "OPTIONS"},
	}).Register(r)

	basicCommand.NewHttpHandler(command.HTTPRESTAudit, basicCommand.CommandHTTPHandlerOptions{
		REST:  "audits",
		Log:   true,
		Valid: true,
		Auth:  false,
	}).Register(r)

	// ADD MORE ROUTERS HERE DO NOT DELETE THIS LINE

	r.Use(CorsMiddleware)

	// r.Use(middleware.Timeout)
	// )
	l, err := net.Listen("tcp", ":9991")
	if err != nil {
		fmt.Println(err)
	}
	defer l.Close()
	// l = netutil.LimitListener(l, max)
	srv := &http.Server{Handler: r, ReadTimeout: 30 * time.Second, WriteTimeout: 60 * time.Second}
	srv.Serve(l)
}
