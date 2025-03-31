package ratelimittedhttp_test

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	ratelimittedhttp "github.com/sumeshmurali/ratelimitted-http/ratelimitted_http"
)

const HTTPBUN_ENDPOINT = "http://localhost:80"

func TestClientAbleToSendRequests(t *testing.T) {
	lim := &ratelimittedhttp.NoOpRatelimitter{}
	c := ratelimittedhttp.NewRatelimittedHttpClient(ratelimittedhttp.NewGlobalRatelimiterPolicy(lim))
	req, _ := http.NewRequest("GET", HTTPBUN_ENDPOINT + "/get", nil)
	resp, err := c.Do(req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
}


func TestClientRespectsRateLimit(t *testing.T) {
	policy := ratelimittedhttp.NewDomainRatelimittingPolicy()
	domain, _ := url.ParseRequestURI(HTTPBUN_ENDPOINT)
	policy.AddDomainLimit(domain.Host, ratelimittedhttp.NewTokenBucketRatelimitter(1, 1))
	c := ratelimittedhttp.NewRatelimittedHttpClient(policy)
	req, _ := http.NewRequest("GET", HTTPBUN_ENDPOINT + "/get", nil)
	_, err := c.Do(req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	res := make(chan *http.Response)
	go func() {
		resp, err := c.Do(req)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		res <- resp
	}()
	select {
	case <- time.After(2 * time.Second):
		t.Fatal("Expected request to complete after 1 second, but run into timeout")
	case <- res:
		// do nothing
	}
}
