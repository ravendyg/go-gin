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
	_user, _ := c.Get("user")
	if _user != nil {
		c.Redirect(http.StatusFound, "/user")
	} else {
		data := gin.H{
			"title": "Register",
		}
		render(c, data, "register.html")
	}
}

// ShowLoginPage -
func ShowLoginPage(c *gin.Context) {
	token, _ := c.Cookie("token")
	if len(token) > 0 {
		user := models.FindByToken(token)
		if user != nil {
			c.Redirect(http.StatusFound, "/user")
			return
		}
	}

	data := gin.H{
		"title": "Login",
	}
	render(c, data, "login.html")
}

// Register -
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user, err := models.RegisterUser(username, password)
	if err == nil {
		setToken(c, user)
		c.Redirect(http.StatusFound, "/")
	} else {
		data := gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error(),
		}
		c.HTML(http.StatusBadRequest, "register.html", data)
	}
}

// Login -
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := models.FindUser(username)
	if user == nil {
		data := gin.H{
			"ErrorTitle": "Not found",
		}
		c.HTML(http.StatusBadRequest, "login.html", data)
		return
	}
	if user.Password != password {
		data := gin.H{
			"ErrorTitle": "Unauthorized",
		}
		c.HTML(http.StatusBadRequest, "login.html", data)
		return
	}

	setToken(c, user)
	c.Redirect(http.StatusFound, "/")
}

//Logout -
func Logout(c *gin.Context) {
	removeToken(c)
	c.Redirect(http.StatusFound, "/")
}

func setToken(c *gin.Context, user *models.User) {
	token := generateToken()
	c.SetCookie("token", token, 3600, "", "", false, true)
	models.AddToken(token, user)
}

func removeToken(c *gin.Context) {
	token, _ := c.Cookie("token")
	if len(token) > 0 {
		models.RemoveToken(token)
		c.SetCookie("token", "", 0, "", "", false, true)
	}
}

func generateToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}
