package dao

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func RdbOpCreate[T any](r *redis.Client, ctx context.Context, value RdbKeyFieldValue[T]) error {
	cmd := r.HSet(ctx, value.Key, value.Field, value.Value)
	err := cmd.Err()
	if err != nil {
		return err
	}
	return nil
}

func RdbOpRead(r *redis.Client, ctx context.Context, key RdbKeyField) (map[string]string, error) {
	cmd := r.HGetAll(ctx, key.Key)
	result, err := cmd.Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func RdbOpDelete(r *redis.Client, ctx context.Context, key RdbKeyField) error {
	cmd := r.HDel(ctx, key.Key, key.Field)
	err := cmd.Err()
	if err != nil {
		return err
	}
	return nil
}
