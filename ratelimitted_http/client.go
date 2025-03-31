package ratelimittedhttp

import (
	"net/http"
)


type RatelimittedHttpClient struct {
	client *http.Client
	policy RatelimittingPolicy
}


func NewRatelimittedHttpClient(policy RatelimittingPolicy) *RatelimittedHttpClient {
	return &RatelimittedHttpClient{
		client: &http.Client{},
		policy: policy,
	}
}

func (c *RatelimittedHttpClient) Do(req *http.Request) (*http.Response, error) {
	c.policy.GetLimiter(req).Wait()
	return c.client.Do(req)
}