package handlers

import (
	"gogin/models"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func AddUser(c *gin.Context) {
	user := &models.User{}
	user.Name = c.Query("name")
	user.Email = c.Query("email")
	user.Password = c.Query("password")
	lastRowID, err := user.Insert()
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	c.String(http.StatusOK, "insert success, the id is: %d", lastRowID)
}
