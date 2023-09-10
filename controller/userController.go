package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/fur"
)

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func GetUser() gin.HandlerFunc{
	return func (c *gin.Context)  {
		
	}
}
func SignUp() gin.HandlerFunc{
	return func (c *gin.Context)  {
		
	}
}
func Login() gin.HandlerFunc{
	return func(ctx *gin.Context) {

	}
}
func HashPassword(password string) string {

}
func VerifyPassword(userPassowrd string , providedPassword string)(bool,string){

}
