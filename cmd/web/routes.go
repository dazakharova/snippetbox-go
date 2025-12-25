package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//Unprotected routes
	mux.Handle("GET /{$}", app.sessionManager.LoadAndSave(preventCSRF(app.authenticate(http.HandlerFunc(app.home)))))
	mux.Handle("GET /snippet/view/{id}", app.sessionManager.LoadAndSave(preventCSRF(app.authenticate(http.HandlerFunc(app.snippetView)))))
	mux.Handle("GET /user/signup", app.sessionManager.LoadAndSave(preventCSRF(app.authenticate(http.HandlerFunc(app.userSignup)))))
	mux.Handle("POST /user/signup", app.sessionManager.LoadAndSave(preventCSRF(app.authenticate(http.HandlerFunc(app.userSignupPost)))))
	mux.Handle("GET /user/login", app.sessionManager.LoadAndSave(preventCSRF(app.authenticate(http.HandlerFunc(app.userLogin)))))
	mux.Handle("POST /user/login", app.sessionManager.LoadAndSave(preventCSRF(app.authenticate(http.HandlerFunc(app.userLoginPost)))))

	//Protected routes
	mux.Handle("GET /snippet/create", app.sessionManager.LoadAndSave(preventCSRF(app.authenticate(app.requireAuthentication(http.HandlerFunc(app.snippetCreate))))))
	mux.Handle("POST /snippet/create", app.sessionManager.LoadAndSave(preventCSRF(app.authenticate(app.requireAuthentication(http.HandlerFunc(app.snippetCreatePost))))))
	mux.Handle("POST /user/logout", app.sessionManager.LoadAndSave(preventCSRF(app.authenticate(app.requireAuthentication(http.HandlerFunc(app.userLogoutPost))))))

	return app.recoverPanic(app.logRequest(commonHeader(mux)))
}
