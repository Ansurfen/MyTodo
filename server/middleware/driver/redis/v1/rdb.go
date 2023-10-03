package rdb

import (
	"MyTodo/conf"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type RedisDB struct {
	*redis.Client
	ctx context.Context
}

func New(opt conf.RedisOption) *RedisDB {
	return &RedisDB{
		Client: redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d", opt.Host, opt.Port),
		}),
		ctx: context.Background(),
	}
}

func (rdb *RedisDB) Ping() error {
	_, err := rdb.Client.Ping(rdb.ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RedisDB) Set(key string, value any, expire time.Duration) error {
	err := rdb.Client.Set(rdb.ctx, key, value, expire).Err()
	if err != nil {
		zap.S().Warn(err)
	}
	return err
}

func (rdb *RedisDB) Get(key string) string {
	val, err := rdb.Client.Get(rdb.ctx, key).Result()
	if err != nil {
		zap.S().Warn(err)
	}
	return val
}

func (rdb *RedisDB) Del(keys ...string) error {
	err := rdb.Client.Del(rdb.ctx, keys...).Err()
	if err != nil {
		zap.S().Warn(err)
	}
	return err
}

func (rdb *RedisDB) SetNX(key string, value any, expire time.Duration) error {
	err := rdb.Client.SetNX(rdb.ctx, key, value, expire).Err()
	if err != nil {
		zap.S().Warn(err)
	}
	return err
}

func (rdb *RedisDB) Do(args ...any) error {
	err := rdb.Client.Do(rdb.ctx, args...).Err()
	if err != nil {
		zap.S().Warn(err)
	}
	return err
}

func (rdb *RedisDB) Eval(script string, keys []string, args ...any) (any, error) {
	res, err := rdb.Client.Eval(rdb.ctx, script, keys, args...).Result()
	if err != nil {
		zap.S().Warn(err)
		return res, err
	}
	return res, nil
}

func (rdb *RedisDB) HSet(key string, values ...any) error {
	err := rdb.Client.HSet(rdb.ctx, key, values...).Err()
	if err != nil {
		zap.S().Warn(err)
	}
	return err
}

func (rdb *RedisDB) HGet(key, filed string) string {
	val, err := rdb.Client.HGet(rdb.ctx, key, filed).Result()
	if err != nil {
		zap.S().Warn(err)
	}
	return val
}

func (rdb *RedisDB) HDel(key string, fields ...string) error {
	err := rdb.Client.HDel(rdb.ctx, key, fields...).Err()
	if err != nil {
		zap.S().Warn(err)
	}
	return err
}

func (rdb *RedisDB) LPush(key string, values ...any) error {
	err := rdb.Client.LPush(rdb.ctx, key, values).Err()
	if err != nil {
		zap.S().Warn(err)
	}
	return err
}

func (rdb *RedisDB) RPush(key string, values ...any) error {
	err := rdb.Client.RPush(rdb.ctx, key, values).Err()
	if err != nil {
		zap.S().Warn(err)
	}
	return err
}

func (rdb *RedisDB) LRange(key string, start, end int64) []string {
	return rdb.Client.LRange(rdb.ctx, key, start, end).Val()
}

func (rdb *RedisDB) Incr(key string) error {
	err := rdb.Client.Incr(rdb.ctx, key).Err()
	if err != nil {
		zap.S().Warn(err)
	}
	return err
}

func (rdb *RedisDB) GetWithPrefix(match string, count int64) []string {
	var cursor uint64
	match = "*" + match + "*"
	keys, _, err := rdb.Client.Scan(rdb.ctx, cursor, match, count+1).Result()
	if err != nil {
		zap.S().Warn(err)
		return nil
	}
	var ret []string
	for _, key := range keys {
		ret = append(ret, rdb.Get(key))
	}
	return ret
}

func (rdb *RedisDB) GetKeysWithPrefix(match string, count int64) []string {
	match = "*" + match + "*"
	var cursor uint64
	keys, _, err := rdb.Client.Scan(rdb.ctx, cursor, match, count+1).Result()
	if err != nil {
		zap.S().Warn(err)
		return nil
	}
	var ret []string
	for _, key := range keys {
		ret = append(ret, rdb.Get(key))
	}
	return ret
}

func (rdb *RedisDB) GetWithKeyAndPrefix(key, match string, count int64) []string {
	var cursor uint64
	keys, _, err := rdb.Client.SScan(rdb.ctx, key, cursor, match, count).Result()
	if err != nil {
		zap.S().Warn(err)
		return nil
	}
	return keys
}

func (rdb *RedisDB) GetKeysWithKeyAndPrefix(key, match string, count int64) []string {
	var cursor uint64
	keys, _, err := rdb.Client.SScan(rdb.ctx, key, cursor, match, count).Result()
	if err != nil {
		zap.S().Warn(err)
		return nil
	}
	return keys
}

func (rdb *RedisDB) ZRange(key, start, end string, count int64) []string {
	res, err := rdb.ZRangeByScore(rdb.ctx, key, &redis.ZRangeBy{
		Min:   start,
		Max:   end,
		Count: count,
	}).Result()
	if err != nil {
		zap.S().Warn(err)
		return nil
	}
	return res
}
