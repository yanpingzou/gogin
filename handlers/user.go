package handlers

import (
	"gogin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewUser(c *gin.Context) {
	user := &models.User{}
	user.Name = c.Query("name")
	user.Email = c.Query("email")
	user.Password = c.Query("password")
	//if err := c.Bind(user); err != nil {
	//	c.String(400, err.Error())
	//	return
	//}
	lastRowID, err := user.Insert()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, "insert success, the id is: %d", lastRowID)
}
