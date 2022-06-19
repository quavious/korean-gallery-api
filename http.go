package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/quavious/gallery-api/external"
	"github.com/quavious/gallery-api/internal"
)

const galleryURL = "http://api.visitkorea.or.kr/openapi/service/rest/PhotoGalleryService"
const listURL = galleryURL + "/galleryList"
const detailURL = galleryURL + "/galleryDetailList"
const searchURL = galleryURL + "/gallerySearchList"
const papagoURL = "https://openapi.naver.com/v1/papago/n2mt"

func listRequest(size int, page int, order string, apiKey string) *internal.ListResponse {
	if !(order == "A" || order == "B" || order == "C" || order == "D") {
		log.Println("invalid parameter - order")
		return nil
	}
	req, err := http.NewRequest("GET", listURL, nil)
	if err != nil {
		log.Printf("new request error: %s", err.Error())
		return nil
	}
	query := req.URL.Query()
	query.Add("numOfRows", strconv.Itoa(size))
	query.Add("pageNo", strconv.Itoa(page))
	query.Add("MobileOS", "ETC")
	query.Add("MobileApp", "GalleryAPI")
	query.Add("ServiceKey", apiKey)
	query.Add("arrange", order)
	query.Add("_type", "json")
	req.URL.RawQuery = query.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("request error: %s", err.Error())
		return nil
	}
	defer resp.Body.Close()

	rawData := new(external.PhotoListResponse)
	err = json.NewDecoder(resp.Body).Decode(rawData)
	if err != nil {
		log.Printf("decode error: %s", err.Error())
		return nil
	}
	responseData := new(internal.ListResponse)
	responseData.Items = rawData.Convert()

	return responseData
}

func searchRequest(size int, page int, order string, keyword string, apiKey string) *internal.ListResponse {
	if order != "A" && order != "B" && order != "D" {
		log.Println("invalid parameter - order")
		return nil
	}
	if len(keyword) < 2 {
		log.Println("invalid parameter - keyword")
		return nil
	}
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		log.Printf("new request error: %s", err.Error())
		return nil
	}
	query := req.URL.Query()
	query.Add("numOfRows", strconv.Itoa(size))
	query.Add("pageNo", strconv.Itoa(page))
	query.Add("MobileOS", "ETC")
	query.Add("MobileApp", "GalleryAPI")
	query.Add("ServiceKey", apiKey)
	query.Add("arrange", order)
	query.Add("keyword", keyword)
	query.Add("_type", "json")
	req.URL.RawQuery = query.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("request error: %s", err.Error())
		return nil
	}
	defer resp.Body.Close()

	rawData := new(external.PhotoListResponse)
	err = json.NewDecoder(resp.Body).Decode(rawData)
	if err != nil {
		log.Printf("decode error: %s", err.Error())
		return nil
	}
	responseData := new(internal.ListResponse)
	responseData.Items = rawData.Convert()

	return responseData
}

func translateRequest(id string, key string, keyword string) string {
	req, err := http.NewRequest(http.MethodPost, papagoURL, nil)
	if err != nil {
		log.Printf("new request error: %s", err.Error())
		return ""
	}
	query := req.URL.Query()
	query.Add("source", "en")
	query.Add("target", "ko")
	query.Add("text", keyword)
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("X-Naver-Client-Id", id)
	req.Header.Set("X-Naver-Client-Secret", key)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("request error: %s", err.Error())
		return ""
	}
	defer resp.Body.Close()

	response := new(external.PapagoResponse)
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		log.Printf("decode error: %s", err.Error())
		return ""
	}
	return response.Message.Result.TranslatedText
}
