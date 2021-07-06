# go get - download and use external libraries

This is a simple web server which uses a third-party package called "mux." In the import
section, you will see that it has been imported from "github.com/gorilla/mux."
However, since we don't have this package stored locally, an error will occur when
we try to run the program. To get the third-party package, you can use go get. This will download it locally so
that our Go code can make use of it: 
` go get github.com/gorilla/mux `
