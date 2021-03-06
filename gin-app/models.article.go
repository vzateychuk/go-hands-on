package main

import (
	"errors"
	"strconv"
)

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// Return a list of all the articles
func findAll() []article {
	return articleList
}

// Fetches the article
func findByID(id int) (*article, error) {
	for _, article := range articleList {
		if article.ID == id {
			return &article, nil
		}
	}
	return nil, errors.New("Not found article by id: " + strconv.Itoa(id))
}

// Create a new article with the title and content provided
func saveNew(a article) (*article, error) {

	// Add the article to the list of articles
	articleList = append(articleList, a)

	return &a, nil
}
