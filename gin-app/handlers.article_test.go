// handlers.article_test.go

package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test that a GET request to the home page returns the home page with
// the HTTP code 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		// You can carry out a lot more detailed tests using libraries that can
		// parse and process HTML pages
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}

// Test that the application returns a JSON list of articles when the Accept header is set to application/json
func TestArticleListJSON(t *testing.T) {
	r := getRouter(true)

	// Define the route similar to its definition in the routes file
	r.GET("/", showIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Accept", "application/json")

	// The function asserts that the response is JSON which can be converted to an array of Article structs
	assertRespList := func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		resp, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}

		var articles []article
		err = json.Unmarshal(resp, &articles)

		return err == nil && statusOK && len(articles) >= 2
	}

	// test Http Response with assertRespListFunc
	testHTTPResponse(t, r, req, assertRespList)
}

// Test that the application returns a JSON list of articles when the Accept header is set to application/json
func TestArticleListXML(t *testing.T) {
	r := getRouter(true)

	// Define the route similar to its definition in the routes file
	r.GET("/", showIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Accept", "application/xml")

	// The function asserts that the response is JSON which can be converted to an array of Article structs
	assertRespList := func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		resp, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}

		var articles []article
		// К сожалению это не работает из за https://stackoverflow.com/questions/27553274/unmarshal-xml-array-in-golang-only-getting-the-first-element
		// Нужно дополнительно либо заворачивать в доп.тег <elems>...</elems> либо делать дополнительное декодирование
		err = xml.Unmarshal(resp, &articles)

		return err == nil && statusOK && len(articles) >= 2
	}

	// test Http Response with assertRespListFunc
	testHTTPResponse(t, r, req, assertRespList)
}

// Test the application returns an article in XML format when the Accept header is set to application/xml
func TestArticleXML(t *testing.T) {
	r := getRouter(true)

	r.GET("/article/view/:article_id", getArticle)

	req, _ := http.NewRequest("GET", "/article/view/1", nil)
	req.Header.Add("Accept", "application/xml")

	// The function asserts that the response is JSON which can be converted to an array of Article structs
	assertRespArticle := func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		resp, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}

		var arte article
		err = xml.Unmarshal(resp, &arte)

		return err == nil && arte.ID == 1 && statusOK
	}

	testHTTPResponse(t, r, req, assertRespArticle)
}

// Test the application returns an article in JSON format when the Accept header is set to application/json
func TestArticleJSON(t *testing.T) {
	r := getRouter(true)

	r.GET("/article/view/:article_id", getArticle)

	req, _ := http.NewRequest("GET", "/article/view/1", nil)
	req.Header.Add("Accept", "application/json")

	// The function asserts that the response is JSON which can be converted to an array of Article structs
	assertRespArticle := func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		resp, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}

		var arte article
		err = json.Unmarshal(resp, &arte)

		return err == nil && arte.ID == 1 && statusOK
	}

	testHTTPResponse(t, r, req, assertRespArticle)
}
