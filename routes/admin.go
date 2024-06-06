package routes

import (
	"login/database"
	"login/jwt"
	"login/model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var AdFeach model.AdminModel
var AdError string
var UserTable []model.UserModel

const RoleAdmin = "admin"

func AdminLogin(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	check := session.Get(RoleAdmin)
	if check == nil {

		//Updatefeach = model.UserModel{}
		c.HTML(http.StatusSeeOther, "adminlogin.html", AdError)
		AdError = ""
	} else {
		c.Redirect(http.StatusSeeOther, "/adminhome")
	}

}

func AdminPost(c *gin.Context) {

	AdFeach = model.AdminModel{}

	database.DB.First(&AdFeach, "email=?", c.Request.FormValue("username"))

	if AdFeach.Password != c.Request.FormValue("password") {
		AdError = "Invalid Username or Password"
		c.Redirect(http.StatusSeeOther, "/admin")

	} else {
		jwt.JwtToken(c, AdFeach.Email, RoleAdmin)
		c.Redirect(http.StatusSeeOther, "/adminhome")
	}

}

func AdminHome(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	check := session.Get(RoleAdmin)
	if check != nil {
		database.DB.Find(&UserTable)
		c.HTML(http.StatusSeeOther, "admin.html", gin.H{
			"Datas": UserTable,
			"Admin": AdFeach.Name,
			"Erro":  AdError,
		})
		AdError = ""
	} else {
		c.Redirect(http.StatusSeeOther, "/admin")
	}
}

func DeleteUsers(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	check := session.Get(RoleAdmin)
	if check != nil {
		User := c.Param("ID")

		database.DB.First(&Updatefeach, "ID=?", User)

		database.DB.Delete(&Updatefeach)
		Updatefeach = model.UserModel{}
		c.Redirect(http.StatusSeeOther, "/adminhome")
	} else {

		c.Redirect(http.StatusSeeOther, "/admin")
	}

}

func BlockUsers(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	check := session.Get(RoleAdmin)
	if check != nil {

		User := c.Param("ID")

		database.DB.First(&Updatefeach, "ID=?", User)

		if Updatefeach.Status == "Active" {
			Updatefeach.Status = "Blocked"
			database.DB.Save(&Updatefeach)
			Updatefeach = model.UserModel{}
			c.Redirect(http.StatusSeeOther, "/adminhome")
		} else {
			Updatefeach.Status = "Active"
			database.DB.Save(&Updatefeach)
			Updatefeach = model.UserModel{}
			c.Redirect(http.StatusSeeOther, "/adminhome")
		}
	} else {
		c.Redirect(http.StatusSeeOther, "/admin")
	}

}

// func UnBlockUsers(c *gin.Context) {
// 	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
// 	session := sessions.Default(c)
// 	check := session.Get(RoleAdmin)
// 	if check != nil {
// 		User := c.Param("ID")

// 		database.DB.First(&Updatefeach, "ID=?", User)

// 		Updatefeach.Status = "Active"
// 		database.DB.Save(&Updatefeach)
// 		Updatefeach = model.UserModel{}
// 		c.Redirect(http.StatusSeeOther, "/adminhome")
// 	} else {
// 		c.Redirect(http.StatusSeeOther, "/admin")
// 	}

// }

func Edit(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	check := session.Get(RoleAdmin)
	if check != nil {
		User := c.Param("ID")
		c.HTML(http.StatusSeeOther, "edit.html", User)
	} else {
		c.Redirect(http.StatusSeeOther, "/admin")
	}

}

func EditUser(c *gin.Context) {
	User := c.Param("ID")

	database.DB.First(&Updatefeach, "ID=?", User)
	Updatefeach.Name = c.Request.FormValue("name")
	Updatefeach.Email = c.Request.FormValue("email")
	database.DB.Save(&Updatefeach)
	AdError = "Saved"
	Updatefeach = model.UserModel{}
	c.Redirect(http.StatusSeeOther, "/adminhome")

}

func Adlogout(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	session.Delete(RoleAdmin)
	session.Save()
	c.Redirect(http.StatusSeeOther, "/admin")
}
