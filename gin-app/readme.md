# Go Web Application Using Gin lib

## Application Functionality

The application we’ll build is a simple article manager. This application should:

* Let users register with a username and a password (non-logged in users only),
* Let users login with a username and a password (non-logged in users only),
* Let users log out (logged in users only),
* Let users create new articles (logged in users only),
Display the list of all articles on the home page (for all users), and
Display a single article on its own page (for all users).
In addition to this functionality, the list of articles and a single article should be accessible in the HTML, JSON and XML formats.

This will allow us to illustrate how Gin can be used to design traditional web applications, API servers, and microservices.

To achieve this, we will make use of the following functionalities offered by Gin:

Routing — to handle various URLs,
Custom rendering — to handle the response format, and
Middleware — to implement authentication.
We’ll also write tests to validate that all the features work as intended.

## Links
[Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin)