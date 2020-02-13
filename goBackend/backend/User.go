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
	if u.Email == "admin" && u.Password == "password" {
		return u.Email, nil
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
