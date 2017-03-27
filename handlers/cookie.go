package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetCookie(c *gin.Context) {
	if cookie, err := c.Request.Cookie("testcookie"); err == nil {
		c.String(http.StatusOK, cookie.Value)
	} else {
		cookie := &http.Cookie{
			Name:  "testcookie",
			Value: "123",
		}
		http.SetCookie(c.Writer, cookie)
		c.String(http.StatusOK, "cookie set success")
	}
}