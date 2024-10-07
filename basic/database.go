package basic

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var locker sync.RWMutex

var TABLE_PREFIX = "prefix_"

func init() {
	locker.Lock()
	defer locker.Unlock()
	env, _ := godotenv.Read(".env")
	url := env["MYSQL_DNS"]
	if url == "" {
		url = "forge:forge@tcp(127.0.0.1:3306)/database?charset=utf8&parseTime=True&loc=Local"
	}
	TABLE_PREFIX = env["TABLE_PREFIX"]

	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       url,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: TABLE_PREFIX, // table name prefix, table for `User` would be `t_users`
			// SingularTable: true,                              // use singular table name, table for `User` would be `user` with this option enabled
			// NoLowerCase:   true,                              // skip the snake_casing of names
			// NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
	})
	if err != nil {
		panic(err)
	}
	// db = db.Debug()
}

func GetDatabase() *gorm.DB {
	locker.RLock()
	defer locker.RUnlock()
	return db
}
