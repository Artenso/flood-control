package service_provider

import (
	"context"
	"log"
	"time"

	"github.com/Artenso/FloodControl/internal/config"
	"github.com/Artenso/FloodControl/internal/service"
	"github.com/go-redis/redis_rate/v10"

	"github.com/redis/go-redis/v9"
)

const (
	cfgPath = "/home/artenso/go/src/flood-control-vk-task/config.yml"
)

// serviceProvider di-container
type serviceProvider struct {
	config  *config.Config
	dbCli   *redis.Client
	limiter *redis_rate.Limiter
	service *service.Service
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) getConfig(_ context.Context) *config.Config {
	cfg, err := config.Read(cfgPath)
	if err != nil {
		log.Fatalf("failed to read config: %s", err.Error())
	}
	s.config = cfg
	return s.config
}

func (s *serviceProvider) getDbCli(ctx context.Context) *redis.Client {
	if s.dbCli == nil {
		opt, err := redis.ParseURL(s.getConfig(ctx).RedisURL)
		if err != nil {
			log.Fatalf("failed to parse url: %s", err.Error())
		}

		s.dbCli = redis.NewClient(opt)
		s.dbCli.FlushDB(ctx).Err()
	}

	return s.dbCli
}

func (s *serviceProvider) getLimiter(ctx context.Context) *redis_rate.Limiter {
	if s.limiter == nil {
		s.limiter = redis_rate.NewLimiter(s.getDbCli(ctx))
	}
	return s.limiter
}

func (s *serviceProvider) getService(ctx context.Context) *service.Service {
	if s.service == nil {
		limit := &redis_rate.Limit{
			Rate:   int(s.getConfig(ctx).CallsLimit),
			Period: time.Duration(s.config.TimeInterval * int64(time.Second)),
			Burst:  int(s.getConfig(ctx).CallsLimit),
		}
		s.service = service.New(limit, s.getLimiter(ctx))
	}

	return s.service
}
