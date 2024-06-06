package helper

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func EnvLoder() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to connect env")
	} 
}
//.....................
func AuthMiddleware(c *gin.Context){
	session:= sessions.Default(c)

	auth, ok := session.Get("authenticate").(bool)
	if !ok || !auth{
		c.Redirect(http.StatusSeeOther,"/")
		c.Abort()
	}
	c.Next()
}
//.......................