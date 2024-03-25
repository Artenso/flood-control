package service

import (
	"context"
	"fmt"
)

func (s *Service) Check(ctx context.Context, userID int64) (bool, error) {

	key := fmt.Sprintf("user:%d", userID)

	res, err := s.limiter.Allow(ctx, key, *s.limit)
	if err != nil {
		return false, err
	}

	return res.Allowed >= 1, nil
}
