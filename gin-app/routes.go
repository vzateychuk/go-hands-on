package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)

	// Group article related routes
	articleRoutes := router.Group("/article")
	{
		// Handle GET requests at /article/view/some_article_id
		articleRoutes.GET("/view/:article_id", getArticle)
	}
}
