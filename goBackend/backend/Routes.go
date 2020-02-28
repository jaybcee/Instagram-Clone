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

func userRoute(c *gin.Context) {
	username := c.Param("id")

	posts, err := getPostsByName(username)

	if err != nil || posts == nil {
		c.JSON(404, struct {
			UserNotFound bool `json:"userNotFound"`
		}{true})
		return
	}

	c.JSON(200, struct {
		Posts        []prisma.Post `json:"posts"`
		UserNotFound bool          `json:"userNotFound"`
	}{posts, false})
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
