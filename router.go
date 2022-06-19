package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func photoListRoute(mux *http.ServeMux, apiKey string) {
	mux.HandleFunc("/photos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
			return
		}
		query := r.URL.Query()
		page, err := strconv.Atoi(query.Get("page"))
		if err != nil {
			log.Printf("invalid query - page: %s", err.Error())
			page = 1
		}
		size, err := strconv.Atoi(query.Get("size"))
		if err != nil {
			log.Printf("invalid query - size: %s", err.Error())
			size = 10
		}
		order := query.Get("order")
		order = orderMap[order]
		if order == "" {
			order = "B"
		}
		response := listRequest(size, page, order, apiKey)
		if response == nil {
			log.Println("nil data")
		}
		b, err := json.Marshal(response)
		w.Header().Add("content-type", "application/json")
		w.Write(b)
	})
}

func photoSearchRoute(mux *http.ServeMux, apiKey string, id string, key string) {
	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
			return
		}
		query := r.URL.Query()
		page, err := strconv.Atoi(query.Get("page"))
		if err != nil {
			log.Printf("invalid query - page: %s", err.Error())
			page = 1
		}
		size, err := strconv.Atoi(query.Get("size"))
		if err != nil {
			log.Printf("invalid query - size: %s", err.Error())
			size = 10
		}
		order := query.Get("order")
		order = orderMap[order]
		if order == "" {
			order = "B"
		}
		keyword := query.Get("keyword")
		if len(keyword) < 1 {
			log.Println("invalid query - keyword")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid keyword"))
			return
		}
		keyword = translateRequest(id, key, keyword)
		if len(keyword) < 1 {
			log.Printf("translate error")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid keyword"))
			return
		}

		response := searchRequest(size, page, order, keyword, apiKey)
		if response == nil {
			log.Println("nil data")
		}
		b, err := json.Marshal(response)
		w.Header().Add("content-type", "application/json")
		w.Write(b)
	})
}
