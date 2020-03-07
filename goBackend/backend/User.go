package main

import (
	"Nicolas-MacBeth/main/backend/generated/prisma-client"
	"context"
	"sort"

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

func unfollowAUser(unfollower string, unfollowee string) (*prisma.User, error, error) {
	client := prisma.New(nil)
	ctx := context.TODO()

	// remove the unfollower from followers for the unfollowee
	unfolloweeUser, err1 := client.UpdateUser(prisma.UserUpdateParams{
		Where: prisma.UserWhereUniqueInput{
			Name: &unfollowee,
		},
		Data: prisma.UserUpdateInput{
			Followers: &prisma.UserUpdateManyWithoutFollowingInput{
				Disconnect: []prisma.UserWhereUniqueInput{
					{Name: &unfollower},
				},
			},
		},
	}).Exec(ctx)

	// remove the unfollowee from following for the unfollower
	_, err2 := client.UpdateUser(prisma.UserUpdateParams{
		Where: prisma.UserWhereUniqueInput{
			Name: &unfollower,
		},
		Data: prisma.UserUpdateInput{
			Following: &prisma.UserUpdateManyWithoutFollowersInput{
				Disconnect: []prisma.UserWhereUniqueInput{
					{Name: &unfollowee},
				},
			},
		},
	}).Exec(ctx)

	return unfolloweeUser, err1, err2
}

func followAUser(follower string, followee string) (*prisma.User, error, error) {
	client := prisma.New(nil)
	ctx := context.TODO()

	// Add the follower as a follower for the followee
	followeeUser, err1 := client.UpdateUser(prisma.UserUpdateParams{
		Where: prisma.UserWhereUniqueInput{
			Name: &followee,
		},
		Data: prisma.UserUpdateInput{
			Followers: &prisma.UserUpdateManyWithoutFollowingInput{
				Connect: []prisma.UserWhereUniqueInput{
					{Name: &follower},
				},
			},
		},
	}).Exec(ctx)

	//Add the followee as following for follower
	_, err2 := client.UpdateUser(prisma.UserUpdateParams{
		Where: prisma.UserWhereUniqueInput{
			Name: &follower,
		},
		Data: prisma.UserUpdateInput{
			Following: &prisma.UserUpdateManyWithoutFollowersInput{
				Connect: []prisma.UserWhereUniqueInput{
					{Name: &followee},
				},
			},
		},
	}).Exec(ctx)

	return followeeUser, err1, err2
}

func getPostsByName(name string) ([]prisma.Post, error) {
	client := prisma.New(nil)
	ctx := context.TODO()

	posts, err := client.User(prisma.UserWhereUniqueInput{
		Name: &name,
	}).Posts(nil).Exec(ctx)

	//Sort all the posts
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].PostedAt > posts[j].PostedAt
	})

	return posts, err
}

func getAlreadyFollowing(name string, selfName string) (bool, error) {
	client := prisma.New(nil)
	ctx := context.TODO()

	users, err := client.Users(&prisma.UsersParams{
		Where: &prisma.UserWhereInput{
			FollowersSome: &prisma.UserWhereInput{
				Name: &selfName,
			},
		}}).Exec(ctx)

	alreadyFollowing := false
	for _, element := range users {
		if element.Name == name {
			alreadyFollowing = true
		}
	}

	return alreadyFollowing, err
}

func getPersonalizedPosts(name string) ([]prisma.Post, error) {
	client := prisma.New(nil)
	ctx := context.TODO()

	users, err := client.Users(&prisma.UsersParams{
		Where: &prisma.UserWhereInput{
			FollowersSome: &prisma.UserWhereInput{
				Name: &name,
			},
		}}).Exec(ctx)

	var posts []prisma.Post

	for _, element := range users {
		tempPosts, _ := client.User(prisma.UserWhereUniqueInput{
			Name: &element.Name,
		}).Posts(nil).Exec(ctx)
		posts = append(posts, tempPosts...)
	}

	//Sort all the posts
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].PostedAt > posts[j].PostedAt
	})

	return posts, err
}

func getFollowersByName(name string) ([]prisma.User, error) {
	client := prisma.New(nil)
	ctx := context.TODO()

	followers, err := client.User(prisma.UserWhereUniqueInput{
		Name: &name,
	}).Followers(nil).Exec(ctx)

	return followers, err
}

func getFollowingByName(name string) ([]prisma.User, error) {
	client := prisma.New(nil)
	ctx := context.TODO()

	following, err := client.User(prisma.UserWhereUniqueInput{
		Name: &name,
	}).Following(nil).Exec(ctx)

	return following, err
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

	user, err := client.Comment(prisma.CommentWhereUniqueInput{UniqueName: &id}).User().Exec(ctx)

	return user, err
}

func fetchUserFromPost(id string) (*prisma.User, error) {
	client := prisma.New(nil)
	ctx := context.TODO()

	user, err := client.Post(prisma.PostWhereUniqueInput{ID: &id}).Owner().Exec(ctx)

	return user, err
}
