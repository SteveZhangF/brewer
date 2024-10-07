module kenneth/backend

go 1.14

replace kenneth/backend => ../backend

require (
	github.com/go-sql-driver/mysql v1.8.1
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/mux v1.8.1
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/mozillazg/go-pinyin v0.20.0 // indirect
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646 // indirect
	github.com/pkg/errors v0.9.1
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/sirupsen/logrus v1.9.3
	golang.org/x/crypto v0.25.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.10
)
