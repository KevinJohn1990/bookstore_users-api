package services

import (
	"fmt"
	"log"

	"github.com/KevinJohn1990/bookstore_users-api/domain/users"

	"github.com/KevinJohn1990/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return &result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUser(userId int64) *errors.RestErr{
	user := &users.User{Id: userId}
	return user.Delete()
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	log.Println("services - UpdateUser")
	current, err := GetUser(user.Id)
	if err != nil{
		return nil, err
	}	
	if isPartial{
		if user.FirstName != ""{
			current.FirstName = user.FirstName
		}
		if user.LastName != ""{
			current.LastName = user.LastName
		}
		if user.Email != ""{
			current.Email = user.Email
		}
	} else{
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	fmt.Println("User:", current)

	if err := current.Update(); err != nil{
		return nil, err
	}
	return current, nil
}