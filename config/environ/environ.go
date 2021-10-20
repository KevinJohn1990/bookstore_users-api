package environ

import (
	"fmt"
	"os"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host"
	mysql_users_port     = "mysql_users_port"
	mysql_users_schema   = "mysql_users_schema"
)

func SetConfig(showVariables bool) {
	err := os.Setenv(mysql_users_username, "root")
	handleErrWithPanic(err)

	os.Setenv(mysql_users_password, "")
	handleErrWithPanic(err)

	os.Setenv(mysql_users_host, "localhost")
	handleErrWithPanic(err)

	os.Setenv(mysql_users_port, "3306")
	handleErrWithPanic(err)

	os.Setenv(mysql_users_schema, "users_db")
	handleErrWithPanic(err)

	if showVariables {
		DisplayConfig()
	}
}

func DisplayConfig() {
	fmt.Println(mysql_users_username, " : ", os.Getenv(mysql_users_username))
	fmt.Println(mysql_users_password, " : ", os.Getenv(mysql_users_password))
	fmt.Println(mysql_users_host, " : ", os.Getenv(mysql_users_host))
	fmt.Println(mysql_users_port, " : ", os.Getenv(mysql_users_port))
	fmt.Println(mysql_users_schema, " : ", os.Getenv(mysql_users_schema))
}

func handleErrWithPanic(err error) {
	if err != nil {
		panic(err)
	}
}
