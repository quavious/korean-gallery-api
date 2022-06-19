package main

import "net/http"

func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET, OPTIONS, HEAD")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		next.ServeHTTP(w, r)
	})
}
