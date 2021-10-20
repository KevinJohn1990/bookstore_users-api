package users

import (
	"log"

	"github.com/KevinJohn1990/bookstore_users-api/datasources/mysql/users_db"
	"github.com/KevinJohn1990/bookstore_users-api/utils/date_utils"
	"github.com/KevinJohn1990/bookstore_users-api/utils/errors"
	"github.com/KevinJohn1990/bookstore_users-api/utils/mysql_utils"
)

const (
	// emailUniqueIndex  = "email_unique"
	// noRowsInResultSet = "no rows in result set"

	queryInsertUser = "insert into user(first_name, last_name, email, date_created) values(?, ?, ?, STR_TO_DATE(?,'%Y-%m-%d %H:%i:%s'));"
	queryUpdateUser = "update user set first_name = ?, last_name = ?, email = ? where id = ?;"
	queryDeleteUser = "delete from user where id = ?;"
	
	queryGetUser    = "select id, first_name, last_name, email, date_created from user where id = ?;"
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	row := users_db.Client.QueryRow(queryGetUser, user.Id)
	// if multiple rows are returned, you will have to close it
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)
	if err != nil {
		return mysql_utils.ParseError(err)		
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertRes, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	userId, err := insertRes.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr{
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	log.Println("Update err: ", err)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	log.Println("Update err: ", err)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Delete() *errors.RestErr{
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	log.Println("delete err: ", err)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id)
	log.Println("delete err: ", err)
	if err != nil{
		return mysql_utils.ParseError(err)		
	}
	return nil
}