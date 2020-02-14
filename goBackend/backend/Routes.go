package main

import (
	"Nicolas-MacBeth/main/backend/generated/prisma-client"
	"io"
	"os"

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

	file, _, err := c.Request.FormFile("file")
	caption := c.PostForm("caption")

	randomId := uuid.New().String()

	out, err := os.Create("./photos/" + randomId)
	check(err)
	defer out.Close()
	_, err = io.Copy(out, file)

	postPicture(email, randomId, caption)

	c.JSON(200, struct {
		Message string `json:"message"`
	}{"Thank you photo"})
}
