package ratelimittedhttp

import "net/http"

type RatelimittingPolicy interface {
	GetLimiter(req *http.Request) Ratelimitter
}

// GlobalRatelimiterPolicy is a policy that applies a single ratelimitter to all requests.
type GlobalRatelimiterPolicy struct {
	limiter Ratelimitter
}

func NewGlobalRatelimiterPolicy(limiter Ratelimitter) *GlobalRatelimiterPolicy {
	return &GlobalRatelimiterPolicy{
		limiter: limiter,
	}
}

func (p *GlobalRatelimiterPolicy) GetLimiter(req *http.Request) Ratelimitter {
	return p.limiter
}

type DomainRatelimittingPolicy struct {
	limits map[string]Ratelimitter
}

func NewDomainRatelimittingPolicy() *DomainRatelimittingPolicy {
	return &DomainRatelimittingPolicy{
		limits: make(map[string]Ratelimitter),
	}
}

func (p *DomainRatelimittingPolicy) GetLimiter(req *http.Request) Ratelimitter {
	l, ok := p.limits[req.URL.Host]
	if !ok {
		// panic if limiter is not found
		panic("No ratelimitter found for domain " + req.URL.Host)
	}
	return l
}

func (p *DomainRatelimittingPolicy) AddDomainLimit(domain string, limit Ratelimitter) {
	p.limits[domain] = limit
}