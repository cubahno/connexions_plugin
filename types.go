package connexions_plugin

import "net/url"

type RequestedResource struct {
	Resource string
	Method   string
	URL      *url.URL
	Headers  map[string][]string
	Body     []byte
	Response *HistoryResponse
}

type HistoryResponse struct {
	Data           []byte
	StatusCode     int
	IsFromUpstream bool
}
