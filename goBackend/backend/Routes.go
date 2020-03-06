package main

import (
	"Nicolas-MacBeth/main/backend/generated/prisma-client"
	"io"
	"os"
	"path/filepath"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func signupRoute(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := c.BindJSON(&req)
	check(err)

	_, addingUserError := AddUser(req.Name, req.Email, req.Password)

	check(addingUserError)

	c.String(200, "Success")
}

func getHome(c *gin.Context) {
	var req struct {
		User string `json:"user"`
	}

	err := c.BindJSON(&req)
	check(err)

	posts, err1 := getPersonalizedPosts(req.User)

	if err != nil || err1 != nil {
		c.String(500, "Posts are unretrievable")
		return
	}

	c.JSON(200, struct {
		Posts []prisma.Post `json:"posts"`
	}{posts})
}

func userRoute(c *gin.Context) {
	username := c.Param("id")
	selfUsername := c.Param("id2")

	posts, err1 := getPostsByName(username)
	followers, err2 := getFollowersByName(username)
	following, err3 := getFollowingByName(username)
	alreadyFollowing, err4 := getAlreadyFollowing(username, selfUsername)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		c.JSON(404, struct {
			UserNotFound bool `json:"userNotFound"`
		}{true})
		return
	}

	c.JSON(200, struct {
		Posts            []prisma.Post `json:"posts"`
		Followers        int           `json:"followers"`
		Following        int           `json:"following"`
		UserNotFound     bool          `json:"userNotFound"`
		AlreadyFollowing bool          `json:"alreadyFollowing"`
	}{posts, len(followers), len(following), false, alreadyFollowing})
}

func testRoute(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
	}

	err := c.BindJSON(&req)
	check(err)

	var recipient string
	if req.Name != "" {
		recipient = req.Name
	} else {
		recipient = "World"
	}

	c.JSON(200, struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{200, "Hello " + recipient})
}

func protectedTestRoute(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	email := claims["email"].(string)

	if email == "" {
		c.String(400, "Token not valid.")
		return
	}

	var req struct {
		Name string `json:"name"`
	}

	err := c.BindJSON(&req)
	check(err)

	var recipient string
	if req.Name != "" {
		recipient = req.Name
	} else {
		recipient = "World"
	}

	c.JSON(200, struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{"This route is protected", "Hello " + recipient})
}

type postBody struct {
	Photo   string `json:"photo"`
	Caption string `json:"caption"`
}

func postPhoto(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	email := claims["email"].(string)

	if email == "" {
		c.String(400, "Token not valid.")
		return
	}

	file, header, err := c.Request.FormFile("file")
	filename := header.Filename
	extension := filepath.Ext(filename)
	caption := c.PostForm("caption")
	filterBnW := c.PostForm("filterBnW")
	filterSurprise := c.PostForm("filterSurprise")

	randomID := uuid.New().String()

	newFilename := "./photos/" + randomID + extension

	out, err := os.Create(newFilename)
	check(err)
	defer out.Close()
	_, err = io.Copy(out, file)

	if filterBnW == "true" {
		image, _ := loadImage(newFilename)
		_, arr := rgbaToGrayArray(image)
		grayImage := arrayToGray(arr)
		saveImage(newFilename, grayImage)
	}

	if filterSurprise == "true" {
		image, _ := loadImage(newFilename)
		_, arr := rgbaToGrayArray(image)

		kernel := createGaussianKernel(5, 1)

		res := conv2d(arr, kernel)
		kx := conv2d(res, scharrKx)
		ky := conv2d(res, scharrKy)

		sobelMagnitude, _ := magnitudeOfGradient(kx, ky)

		finalImage := arrayToGray(sobelMagnitude)

		saveImage(newFilename, finalImage)

	}

	err = postPicture(email, randomID+extension, caption)
	check(err)

	c.JSON(200, struct {
		Message string `json:"message"`
	}{"Thank you photo"})
}

func postComment(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	email := claims["email"].(string)

	if email == "" {
		c.String(400, "Token not valid.")
		return
	}

	var req struct {
		Comment string `json:"comment"`
		PostID  string `json:"postID"`
	}

	err := c.BindJSON(&req)
	check(err)

	randomID := uuid.New().String()

	err = addComment(email, req.Comment, req.PostID, randomID)

	if err != nil {
		c.JSON(500, struct {
			Message string `json:"message"`
		}{"The comment could not be added"})
		return
	}

	c.JSON(200, struct {
		Message string `json:"message"`
	}{"Thank you comment"})
}

func followOrUnfollowUser(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	email := claims["email"].(string)

	if email == "" {
		c.JSON(403, struct {
			Authorized bool `json:"Authorized"`
		}{false})
		return
	}

	var req struct {
		Follower string `json:"follower"`
		Followee string `json:"followee"`
		Follow   bool   `json:"follow"`
	}

	err := c.BindJSON(&req)
	check(err)

	var (
		err2 error
		err3 error
	)

	// for follows
	if req.Follow == true {
		_, err2, err3 = followAUser(req.Follower, req.Followee)
	} else if req.Follow == false { // for unfollows
		_, err2, err3 = unfollowAUser(req.Follower, req.Followee)
	}

	if err == nil && err2 == nil && err3 == nil {
		c.String(200, "All good in the hood")
	} else {
		c.String(500, "Unable to follow/unfollow")
		return
	}
}

func getUserFromEmail(c *gin.Context) {

	claims := jwt.ExtractClaims(c)
	email := claims["email"].(string)

	user, err := GetUser(email)

	if err != nil {
		c.JSON(404, struct {
			UserNotFound bool `json:"userNotFound"`
		}{true})
		return
	}
	c.JSON(200, struct {
		Name string `json:"name"`
	}{user.Name})
}

func getComments(c *gin.Context) {

	//remove if we decided on not keeping secure

	// claims := jwt.ExtractClaims(c)

	// email := claims["email"].(string)

	// if email == "" {
	// 	c.String(400, "Token not valid.")
	// 	return
	// }

	post := c.Param("id")

	print(post)

	comments, err := fetchComments(post)

	if err != nil {
		c.JSON(404, struct {
			Message string `json:"message"`
		}{"No comments found for this post."})
		return
	}

	c.JSON(200, struct {
		Comments []prisma.Comment `json:"comments"`
	}{comments})
}

func getUserFromComment(c *gin.Context) {

	commentID := c.Param("id")

	user, err := fetchUserFromComment(commentID)
	check(err)

	c.JSON(200, struct {
		Username string `json:"username"`
	}{user.Name})
}

func updateUserComment(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	email := claims["email"].(string)

	if email == "" {
		c.String(400, "Token not valid.")
		return
	}

	var req struct {
		NewComment        string `json:"newComment"`
		CommentUniqueName string `json:"CommentUniqueName"`
	}

	err := c.BindJSON(&req)
	check(err)

	user, err := fetchUserFromComment(req.CommentUniqueName)

	if user.Email != email {
		c.JSON(404, struct {
			Message string `json:"message"`
		}{"Unauthorized access"})
		return
	}

	err = updateComment(req.CommentUniqueName, req.NewComment)

	if err != nil {
		c.JSON(500, struct {
			Message string `json:"message"`
		}{"The comment could not be updated"})
		return
	}

	c.JSON(200, struct {
		Message string `json:"message"`
	}{"Updated comment"})
}

func deleteUserComment(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	email := claims["email"].(string)

	if email == "" {
		c.String(400, "Token not valid.")
		return
	}

	var req struct {
		CommentUniqueName string `json:"CommentUniqueName"`
	}

	err := c.BindJSON(&req)
	check(err)

	user, err := fetchUserFromComment(req.CommentUniqueName)

	if user.Email != email {
		c.JSON(404, struct {
			Message string `json:"message"`
		}{"Unauthorized access"})
		return
	}

	err = deleteComment(req.CommentUniqueName)

	if err != nil {
		c.JSON(500, struct {
			Message string `json:"message"`
		}{"The comment could not be deleted"})
		return
	}

	c.JSON(200, struct {
		Message string `json:"message"`
	}{"Deleted comment"})
}
