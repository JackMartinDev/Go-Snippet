package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.Handle("GET /{$}", sessionWrapper(app.home, app))
	mux.Handle("GET /snippet/view/{id}", sessionWrapper(app.snippetView, app))
	mux.Handle("GET /snippet/create", sessionWrapper(app.snippetCreate, app))
	mux.Handle("POST /snippet/create", sessionWrapper(app.snippetCreatePost, app))

	mux.Handle("GET /user/signup", sessionWrapper(app.userSignup, app))
	mux.Handle("POST /user/signup", sessionWrapper(app.userSignupPost, app))
	mux.Handle("GET /user/login", sessionWrapper(app.userLogin, app))
	mux.Handle("POST /user/login", sessionWrapper(app.userLoginPost, app))
	mux.Handle("POST /user/logout", sessionWrapper(app.userLogoutPost, app))

	return app.recoverPanic(app.logRequest(commonHeaders(mux)))
}

func sessionWrapper(route http.HandlerFunc, app *application) http.Handler {
	return app.sessionManager.LoadAndSave(http.HandlerFunc(route))
}
