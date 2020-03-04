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
	if user.Email != "" && CheckHash(u.Password, user.Password) {
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

func getPostsByName(name string) ([]prisma.Post, error) {
	client := prisma.New(nil)
	ctx := context.TODO()

	posts, err := client.User(prisma.UserWhereUniqueInput{
		Name: &name,
	}).Posts(nil).Exec(ctx)

	return posts, err
}

func AddUser(name string, email string, password string) (*prisma.User, error) {
	client := prisma.New(nil)
	ctx := context.TODO()

	hashedPw, _ := Bhash(password)

	newUser, err := client.CreateUser(prisma.UserCreateInput{
		Name:     name,
		Email:    email,
		Password: hashedPw,
	}).Exec(ctx)

	return newUser, err

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

func addComment(email, comment, post, uniqueName string) error {
	client := prisma.New(nil)
	ctx := context.TODO()

	_, err := client.UpdatePost(prisma.PostUpdateParams{
		Where: prisma.PostWhereUniqueInput{
			ID: &post,
		},
		Data: prisma.PostUpdateInput{
			Comments: &prisma.CommentUpdateManyWithoutPostInput{
				Create: []prisma.CommentCreateWithoutPostInput{
					{
						UniqueName:  uniqueName,
						CommentText: comment,
						User: prisma.UserCreateOneInput{
							Connect: &prisma.UserWhereUniqueInput{
								Email: &email,
							},
						},
					},
				},
			},
		},
	}).Exec(ctx)

	return err
}


func fetchComments(post string) ([]prisma.Comment, error) {
	client := prisma.New(nil)
	ctx := context.TODO()

	orderBy := prisma.CommentOrderByInputCommentedAtAsc

	comments, err := client.Comments(&prisma.CommentsParams{
		Where: &prisma.CommentWhereInput{
			Post: &prisma.PostWhereInput{
				ID: &post,
			},
		},
		OrderBy: &orderBy,
	}).Exec(ctx)

	return comments, err
}

func deleteComment(uniqueName string) error {
	client := prisma.New(nil)
	ctx := context.TODO()

	_, err := client.DeleteComment(prisma.CommentWhereUniqueInput{UniqueName: &uniqueName}).Exec(ctx)

	return err
}

func updateComment(uniqueName, newComment string) error {
	client := prisma.New(nil)
	ctx := context.TODO()

	_, err := client.UpdateComment(prisma.CommentUpdateParams{
		Where: prisma.CommentWhereUniqueInput{UniqueName: &uniqueName},
		Data:  prisma.CommentUpdateInput{CommentText: &newComment},
	}).Exec(ctx)

	return err
}

func fetchUserFromComment(id string) (*prisma.User, error) {
	client := prisma.New(nil)
	ctx := context.TODO()

	user, err := client.Comment(prisma.CommentWhereUniqueInput{UniqueName: &id}).Post().Owner().Exec(ctx)

	return user, err
}
