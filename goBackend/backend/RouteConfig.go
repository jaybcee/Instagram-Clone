package main

import (
	"github.com/gin-gonic/gin"
)

func createRoutes() *gin.Engine {
	//Create router
	r := gin.Default()
	//Import cors for now
	r.Use(CORSMiddleware())
	//Example unsecure route
	r.POST("/api/test", testRoute)
	//Secure Routes
	r, _ = GetAuth(r)
	//Safety in case of crash
	r.Use(gin.Recovery())
	return r
}

//GetSecureRoutes gives us our protected routes. Most things should go here.
func GetSecureRoutes(auth *gin.RouterGroup) *gin.RouterGroup {
	//Example secure route
	auth.POST("/api/secure_test", protectedTestRoute)

	auth.POST("/api/uploadPhoto", postPhoto)
	//Return secure routes
	return (auth)
}
