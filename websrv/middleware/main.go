package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/admin/", adminIndex)
	adminMux.HandleFunc("/admin/panic", panicPage)
	// set middleware
	adminHandler := authMiddleware(adminMux)
	siteMux := http.NewServeMux()
	siteMux.Handle("/admin/", adminHandler)
	siteMux.HandleFunc("/login", loginPage)
	siteMux.HandleFunc("/logout", logoutPage)
	siteMux.HandleFunc("/", mainPage)
	// set middleware
	siteHandler := accessLogMiddleware(siteMux)
	siteHandler = panicMiddleware(siteHandler)
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", siteHandler)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	// учебный пример! это не проверка авторизации!
	loggedIn := (err != http.ErrNoCookie)

	if loggedIn {
		fmt.Fprintln(w, `<a href="/logout">logout</a>`)
		fmt.Fprintln(w, "Welcome, "+session.Value)
	} else {
		fmt.Fprintln(w, `<a href="/login">login</a>`)
		fmt.Fprintln(w, "You need to login")
	}
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(10 * time.Hour)
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   "vzateychuk",
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func logoutPage(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	http.Redirect(w, r, "/", http.StatusFound)
}

func adminIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<a href="/">site index</a>`)
	fmt.Fprintln(w, "Admin main page")
}

func panicPage(w http.ResponseWriter, r *http.Request) {
	panic("this must me recovered")
}

/*
Возвращаем функцию, которая будет принимать и обрабатывать входящий запрос, но при этом мы сделаем в defer обработку паники,
которая может возникнуть при обработке входящего запроса а потом только вызовем функцию обработчик, передаваемый в параметре next.
*/
func panicMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("panicMiddleware", r.URL.Path)
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered", err)
				http.Error(w, "Internal server error", 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

/*
Аналогично создается функция - обработчик, "оборачивающая" передаваемую функцию handler в дополнительный обработчик,
логирующий входной запрос
*/
func accessLogMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("accessLogMiddleware", r.URL.Path)
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("[%s], %s, %s, %s\n", r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))

	})
}

/*
Аналогично создается функция - обработчик, "оборачивающая" передаваемую функцию handler в дополнительный обработчик,
проверяющий авторизацию входного запроса и если проверка не прошла просто возвращающий 404.
*/
func authMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("authMiddleware", r.URL.Path)
		// это учебный пример, просто проверка наличия cookie
		_, err := r.Cookie("session_id")
		if err != nil {
			fmt.Println("Not authenticated at: ", r.URL.Path)
			http.Redirect(w, r, "/", http.StatusNotFound)
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
