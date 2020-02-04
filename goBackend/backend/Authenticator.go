package main

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

//Key could be removed and put into the file.
var identityKey = "email"

type newUserRequest struct {
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Phone    string `json:"Phone"`
}

type loginRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

//GetAuth gives us our authentication routes
func GetAuth(r *gin.Engine) (*gin.Engine, *gin.RouterGroup) {
	//Create Authentication middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		//Name
		Realm: "instagram",
		//Secret signing key
		Key: []byte("CHANGE THIS :)"),
		//JWT key timeout for one year
		Timeout:     time.Hour * 24 * 365,
		IdentityKey: identityKey,
		//Refresh key for one year
		MaxRefresh: time.Hour * 24 * 365,
		//Function to check the JWT key
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			fmt.Println(data)
			v, ok := data.(string)
			if ok {
				return jwt.MapClaims{
					identityKey: v,
				}
			}
			//Return an empty one if we can't parse it
			return jwt.MapClaims{}
		},
		//Function to check the JWT key
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return claims["email"].(string)
		},
		//Function that is executed when someone tries to login
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var req loginRequest

			//Bind request to the struct
			err := c.BindJSON(&req)
			check(err)

			print(req.Email, req.Password)

			if len(req.Email) == 0 || len(req.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			return AuthenticateUser(req)
		},
		//check that the token is valid
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(string); ok {
				return true
			}
			return false
		},
		//What to do when the user is not authorized
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		//Where to look for the authorization token
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		//Prefix for the token
		TokenHeadName: "Bearer",
		//Time right now
		TimeFunc: time.Now,
	})

	//Crash if we have an error here, should never occur
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	//Route for logging in
	r.POST("/login", authMiddleware.LoginHandler)

	//Create a route in case of a page not found
	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	//Create a secure route group for our secured routes
	auth := r.Group("/secure")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	//Pass our routes to the secure group
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		//Imported Routes
		auth = GetSecureRoutes(auth)
	}

	return r, auth
}
