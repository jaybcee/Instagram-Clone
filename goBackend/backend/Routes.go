package main

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
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

	c.JSON(200, struct {
		Status int `json:"status"`
	}{200})
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
