package handlers

import (
	"math/rand"
	"net/http"
	"strconv"
	"web-server/models"

	"github.com/gin-gonic/gin"
)

// ShowRegistrationPage -
func ShowRegistrationPage(c *gin.Context) {
	token, _ := c.Cookie("token")
	if len(token) > 0 {
		user := models.FindByToken(token)
		if user != nil {
			c.Redirect(http.StatusFound, "/u/user")
			return
		}
	}

	data := gin.H{
		"title": "Register",
	}
	render(c, data, "register.html")
}

// ShowUserPage -
func ShowUserPage(c *gin.Context) {
	token, _ := c.Cookie("token")
	if len(token) > 0 {
		user := models.FindByToken(token)
		if user != nil {
			data := gin.H{
				"title":    "User",
				"username": user.Username,
			}
			render(c, data, "user.html")
			return
		}
	}

	c.Redirect(http.StatusFound, "/u/register")
}

// Register -
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user, err := models.RegisterUser(username, password)
	if err == nil {
		token := generateToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		models.AddToken(token, user)

		data := gin.H{
			"title": "Successful registration & Login",
		}
		render(c, data, "login-successful.html")
	} else {
		data := gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error(),
		}
		c.HTML(http.StatusBadRequest, "register.html", data)
	}
}

//Logout -
func Logout(c *gin.Context) {
	token, _ := c.Cookie("token")
	if len(token) > 0 {
		models.RemoveToken(token)
	}
	c.SetCookie("token", "", 0, "", "", false, true)
	c.Redirect(http.StatusFound, "/")
}

func generateToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}
