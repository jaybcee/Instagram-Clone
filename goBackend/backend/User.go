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
	user, err := GetUser(u.Email)

	// Verify if email exists in DB and check if password is correct
<<<<<<< HEAD
	if user.Email != "" && user.Password == u.Password {
=======
	if user.Email != "" && CheckHash(u.Password, user.Password) {
>>>>>>> added pw hashing
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

func postPicture(email string, fileName string, caption string) error {
	client := prisma.New(nil)
	ctx := context.TODO()

	_, err := client.UpdateUser(prisma.UserUpdateParams{
		Where: prisma.UserWhereUniqueInput{
			Email: &email,
		},
		Data: prisma.UserUpdateInput{
			Posts: &prisma.PostUpdateManyWithoutOwnerInput{
				Create: []prisma.PostCreateWithoutOwnerInput{
					{
						FileName: fileName,
						Caption:  caption,
					},
				},
			},
		},
	}).Exec(ctx)

	return err
}
func AddUser(name string, email string, password string) (*prisma.User, error) {
	client := prisma.New(nil)
	ctx := context.TODO()
func AddUser(name string, email string, password string) (*prisma.User, error) {
	client := prisma.New(nil)
	ctx := context.TODO()

	hashedPw, _ := Bhash(password)

	newUser, err := client.CreateUser(prisma.UserCreateInput{
		Name:     name,
		Email:    email,
<<<<<<< HEAD
		Password: password,
=======
		Password: hashedPw,
>>>>>>> added pw hashing
	}).Exec(ctx)

	return newUser, err

}
