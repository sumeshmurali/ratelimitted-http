# ratelimitted-http
A golang requesting client that allows user to ratelimit out going requests to a server

# How to Use

Get the library
```
go get github.com/sumeshmurali/ratelimitted-http/ratelimitted_http
```

```go

package main

import (
    "net/http"

    ratelimittedhttp "github.com/sumeshmurali/ratelimitted-http/ratelimitted_http"
)

func main() {
    // create a domain ratelimiting policy
    policy := ratelimittedhttp.NewDomainRatelimittingPolicy()
    // create a ratelimitter with a capacity ( burst ) and a rate of refil per second
    limiter := ratelimittedhttp.NewTokenBucketRatelimitter(10, 2)
    // add the limit to domain ratelimitting policy
	policy.AddDomainLimit("httpbun.com", limiter)
    
    // use the limiter to create ratelimited client
    client := ratelimittedhttp.NewRatelimittedHttpClient(policy)

    // create a request
   	req, _ := http.NewRequest("GET", "https://httpbun.com/get", nil)
    
    // send the request
    res, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    fmt.Println("successfully recieved response with status code: ", res.StatusCode)
}
```