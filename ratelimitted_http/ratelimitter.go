package ratelimittedhttp

import (
	"context"

	"golang.org/x/time/rate"
)

type Ratelimitter interface {
	Wait()
	Allow() bool
}
// NoOpRatelimitter is a ratelimitter that does nothing. Rarely useful, except for testing.
type NoOpRatelimitter struct{}

func (r *NoOpRatelimitter) Wait() {}

func (r *NoOpRatelimitter) Allow() bool {
	return true
}

// TokenBucketRatelimitter is a ratelimitter that uses a token bucket algorithm.
// The implementation simply wraps the golang.org/x/time/rate.Limiter.
// Use NewTokenBucketRatelimitter to create a new instance.
type TokenBucketRatelimitter struct {
	limiter *rate.Limiter
}
func (r *TokenBucketRatelimitter) Wait() {
	_ = r.limiter.Wait(context.TODO())
}

func (r *TokenBucketRatelimitter) Allow() bool {
	return r.limiter.Allow()
}

// NewTokenBucketRatelimitter creates a new TokenBucketRatelimitter with the given capacity and fill rate.
// - The capacity (aka burst) is the maximum number of tokens that can be stored in the bucket. 
// - The fill rate is the number of tokens added to the bucket per second.
func NewTokenBucketRatelimitter(capacity int, fillRate float64) Ratelimitter {
	limiter := rate.NewLimiter(rate.Limit(fillRate), capacity)
	return &TokenBucketRatelimitter{
		limiter: limiter,
	}
}