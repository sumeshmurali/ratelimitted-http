package ratelimittedhttp_test

import (
	"net/http"
	"testing"
	"time"

	ratelimittedhttp "github.com/sumeshmurali/ratelimitted-http/ratelimitted_http"
)

const HTTPBUN_ENDPOINT = "http://localhost:80"

func TestClientAbleToSendRequests(t *testing.T) {
	c := ratelimittedhttp.NewRatelimittedHttpClient(&ratelimittedhttp.NoOpRatelimitter{})
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
	c := ratelimittedhttp.NewRatelimittedHttpClient(ratelimittedhttp.NewTokenBucketRatelimitter(1, 1))
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