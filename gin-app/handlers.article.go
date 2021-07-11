package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func getArticle(c *gin.Context) {
	// Check if the article ID is valid
	articleID, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err) // If the article is not found, abort with an error
	}

	arte, err := getArticleByID(articleID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound) // If an invalid article ID is specified in the URL, abort with an error
	}

	// Call the HTML method of the Context to render a template
	render(c, gin.H{"title": arte.Title, "payload": arte}, "article.html")

}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
