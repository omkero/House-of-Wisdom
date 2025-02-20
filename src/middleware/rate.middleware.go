package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type rateLimiter struct {
	limiter   *rate.Limiter
	lastSeen  time.Time
	blockedAt time.Time // Timestamp when the client was blocked
}

var (
	rateLimiters = make(map[string]*rateLimiter)
	mu           sync.Mutex
)

func getRateLimiter(ip string, requests, seconds, until int) *rateLimiter {
	mu.Lock()
	defer mu.Unlock()

	// Clean up old entries
	for key, rl := range rateLimiters {
		if time.Since(rl.lastSeen) > 3*time.Minute {
			delete(rateLimiters, key)
		}
	}

	if _, exists := rateLimiters[ip]; !exists {
		rateLimiters[ip] = &rateLimiter{
			limiter:  rate.NewLimiter(rate.Every(time.Duration(seconds)*time.Second)/rate.Limit(requests), requests),
			lastSeen: time.Now(),
		}
	}

	// Update last seen time
	rateLimiters[ip].lastSeen = time.Now()
	return rateLimiters[ip]
}

func RateLimitMiddleware(requests, seconds, until int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		rl := getRateLimiter(ip, requests, seconds, until)

		// Check if the client is blocked
		if !rl.blockedAt.IsZero() && time.Since(rl.blockedAt) < time.Duration(until)*time.Second {
			remaining := int(time.Duration(until)*time.Second-time.Since(rl.blockedAt)) / int(time.Second)
			ctx.AbortWithStatusJSON(429, gin.H{
				"message":     "blocked temporarily",
				"reason":      "too many requests",
				"retry_after": remaining,
			})
			return
		}

		// Reset the block status if the block duration has passed
		if !rl.blockedAt.IsZero() && time.Since(rl.blockedAt) >= time.Duration(until)*time.Second {
			rl.blockedAt = time.Time{} // Reset block time
		}

		// Allow or block requests based on the rate limiter
		if !rl.limiter.Allow() {
			rl.blockedAt = time.Now()
			ctx.AbortWithStatusJSON(429, gin.H{
				"message":     "blocked temporarily",
				"reason":      "too many requests",
				"retry_after": until,
			})
			return
		}

		ctx.Next()
	}
}
