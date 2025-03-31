package ratelimittedhttp

import (
	"net/http"
)


type RatelimittedHttpClient struct {
	client *http.Client
	limiter Ratelimitter
}


func NewRatelimittedHttpClient(limiter Ratelimitter) *RatelimittedHttpClient {
	return &RatelimittedHttpClient{
		client: &http.Client{},
		limiter: limiter,
	}
}

func (c *RatelimittedHttpClient) Do(req *http.Request) (*http.Response, error) {
	c.limiter.Wait()
	return c.client.Do(req)
}