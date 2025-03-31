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

    ratelimitted_http "github.com/sumeshmurali/ratelimitted-http/ratelimitted_http"
)

func main() {
    // create a ratelimitter with a capacity ( burst ) and a rate of refil per second
    limiter := ratelimittedhttp.NewTokenBucketRatelimitter(10, 2)
    // use the limiter to create ratelimited client
    client := ratelimittedhttp.NewRatelimittedHttpClient(limiter)

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