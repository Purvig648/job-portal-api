package caching

import (
	"context"
	"encoding/json"
	"job-application-api/internal/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

//go:generate mockgen -source=cache.go -destination=cache_mock.go -package=caching
type Redis struct {
	client *redis.Client
}
type Cache interface {
	AddCache(ctx context.Context, jid uint, jobData models.Job) error
	FetchCache(ctx context.Context, jid uint) (models.Job, error)
	AddToCacheOtp(ctx context.Context, email string, otp string) error
}

func NewRdbCache(client *redis.Client) Cache {
	return &Redis{
		client: client,
	}
}
func (r *Redis) AddToCacheOtp(ctx context.Context, email string, otp string) error {
	err := r.client.Set(email, otp, 1*time.Minute).Err()
	return err
}

func (r *Redis) AddCache(ctx context.Context, jid uint, jobData models.Job) error {
	jobId := strconv.FormatUint(uint64(jid), 10)
	val, err := json.Marshal(jobData)
	if err != nil {
		return err
	}
	err = r.client.Set(jobId, val, 1*time.Minute).Err()
	return err
}

func (r *Redis) FetchCache(ctx context.Context, jid uint) (models.Job, error) {
	jobId := strconv.FormatUint(uint64(jid), 10)
	str, err := r.client.Get(jobId).Result()
	if err != nil {
		return models.Job{}, err
	}
	var job models.Job
	err = json.Unmarshal([]byte(str), &job)
	if err != nil {
		return models.Job{}, err
	}
	return job, nil
}
