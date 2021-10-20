package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/KevinJohn1990/bookstore_users-api/config/environ"
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host"
	mysql_users_port     = "mysql_users_port"
	mysql_users_schema   = "mysql_users_schema"
)

var (
	Client *sql.DB
	err    error

	username string
	password string
	host     string
	port     string
	schema   string
)

func setVariables() {
	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host = os.Getenv(mysql_users_host)
	port = os.Getenv(mysql_users_port)
	schema = os.Getenv(mysql_users_schema)
}

func init() {
	fmt.Println("Inside users_db init()")
	if username == "" {
		environ.SetConfig(true)
		setVariables()
	}

	drivername := "mysql"
	// user:password@tcp(host:port)/dbname?charset=utf8
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		username, // "root",      // user
		password, // "",          // password
		host,     // "localhost", // host
		port,     // "3306",      // port
		schema,   // "users_db",  // dbname
	)

	// fmt.Println("Drivername : ", drivername)
	// fmt.Println("Datasourcename : ", datasourceName)

	Client, err = sql.Open(drivername, datasourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")

}
