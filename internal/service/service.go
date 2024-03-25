package service

import (
	"github.com/go-redis/redis_rate/v10"
)

type Service struct {
	limit   *redis_rate.Limit
	limiter *redis_rate.Limiter
}

func New(limit *redis_rate.Limit, limiter *redis_rate.Limiter) *Service {
	return &Service{
		limit:   limit,
		limiter: limiter,
	}
}
