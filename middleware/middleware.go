package middleware

import "net/http"

/*
	Esta função é um middleware que adiciona o Content-type application/json
	http.Handler é uma interface que define um método ServeHTTP que recebe um http.ResponseWriter e um *http.Request

*/
func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})

}
