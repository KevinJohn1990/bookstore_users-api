package mysql_utils

import (
	"log"
	"strings"

	"github.com/KevinJohn1990/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)
const(
	noRowsInResultSet = "no rows in result set"

)

func ParseError(err error) *errors.RestErr{
	sqlErr, ok := err.(*mysql.MySQLError)
	log.Println("sql err: ", sqlErr, "; ok:", ok)
	if !ok{
		if strings.Contains(err.Error(), noRowsInResultSet){
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqlErr.Number{
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	
	return errors.NewInternalServerError("error processing request")
}