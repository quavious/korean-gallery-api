package external

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/quavious/gallery-api/internal"
)

type PhotoListResponse struct {
	Response struct {
		Body struct {
			Items struct {
				Item []struct {
					ID        int    `json:"galContentId,omitempty"`
					TypeID    int    `json:"galContentTypeId,omitempty"`
					URL       string `json:"galWebImageUrl,omitempty"`
					Location  string `json:"galPhotographyLocation,omitempty"`
					Title     string `json:"galTitle,omitempty"`
					CreatedAt int    `json:"galCreatedtime,omitempty"`
					Count     int    `json:"galViewCount,omitempty"`
					Keyword   string `json:"galSearchKeyword,omitempty"`
				} `json:"item"`
			} `json:"items"`
		} `json:"body"`
	} `json:"response"`
}

type PapagoResponse struct {
	Message struct {
		Result struct {
			SourceType     string `json:"srcLangType,omitempty"`
			TargetType     string `json:"tarLangType,omitempty"`
			TranslatedText string `json:"translatedText,omitempty"`
		} `json:"result,omitempty"`
	} `json:"message,omitempty"`
}

func (list *PhotoListResponse) Convert() []internal.ListItem {
	convertedList := []internal.ListItem{}
	for _, item := range list.Response.Body.Items.Item {
		listItem := new(internal.ListItem)
		listItem.ID = fmt.Sprintf("%d_%d", item.TypeID, item.ID)
		listItem.Title = item.Title
		listItem.URL = strings.Replace(item.URL, "http:", "https:", 1)
		listItem.Location = item.Location
		listItem.CreatedAt = strconv.Itoa(item.CreatedAt)
		listItem.Count = item.Count
		convertedList = append(convertedList, *listItem)
	}
	return convertedList
}
