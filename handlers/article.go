package handlers

import (
	"fmt"
	"gogin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}

func ShowPageIndex(c *gin.Context) {
	articles := models.GetAllArticles()
	render(c, gin.H{
		"title":   "index pages",
		"payload": articles,
	}, "index.html")
}

func GetArticle(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("article_id"))
	fmt.Println(articleID, err)
	if err == nil {
		article, err := models.GetArticleByID(articleID)
		if err == nil {
			c.HTML(http.StatusOK, "article.html", gin.H{
				"title":   article.Title,
				"payload": article,
			})
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
