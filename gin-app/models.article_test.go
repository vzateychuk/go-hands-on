package main

import "testing"

// Test the function that fetches all articles.
// This test first makes sure that the article list fetched by this function and the article list present in the global
// variable articleList are identical. It then loops over the article list to verify that each article is identical.
// The test fails if either of these two checks fail.
func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// Check that the length of the list of articles returned is the
	// same as the length of the global variable holding the list
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range alist {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {

			t.Fail()
			break
		}
	}
}
