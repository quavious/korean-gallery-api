package internal

type ListResponse struct {
	Items []ListItem `json:"items"`
}

type ListItem struct {
	ID        string `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	URL       string `json:"url,omitempty"`
	Location  string `json:"location,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	Count     int    `json:"count,omitempty"`
}
