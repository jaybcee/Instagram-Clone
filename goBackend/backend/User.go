package main

import (
	"Nicolas-MacBeth/main/backend/generated/prisma-client"
	"context"

	jwt "github.com/appleboy/gin-jwt/v2"
)

//AuthenticateUser function to authenticate user
func AuthenticateUser(u struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}) (string, error) {
	user,err := GetUser(u.Email)

	// Verify if email exists in DB and check if password is correct
	if(user.Email != "" && user.Password == u.Password){
		return user.Email, nil
	}

	if err != nil {
		panic(err)
	}
	
	return "", jwt.ErrFailedAuthentication
}

//GetUser function to get userinfo by email from db
//This is basically an example of how to query the database.
func GetUser(email string) (*prisma.User, error) {
	client := prisma.New(nil)
	ctx := context.TODO()

	user, err := client.User(prisma.UserWhereUniqueInput{
		Email: &email,
	}).Exec(ctx)


	return user, err

}

func AddUser(name string,email string, password string) (*prisma.User, error) {
	client := prisma.New(nil)
	ctx := context.TODO()

	newUser, err := client.CreateUser(prisma.UserCreateInput{
		Name: name,
		Email: email,
		Password : password,
	  }).Exec(ctx)


	return newUser, err

}
