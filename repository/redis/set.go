package redis

import (
	"context"
	"encoding/json"
	"log"
	"time"
)

// SetWithExp save data to redis with key and exp time
func (r *Redis) SetWithExp(ctx context.Context, key string, val interface{}, duration string) error {
	dur, err := time.ParseDuration(duration)
	if err != nil {
		// logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "parse_duration", uc.ReqID)
		return err
	}

	b, err := json.Marshal(val)
	if err != nil {
		// logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "json_marshal", uc.ReqID)
		return err
	}

	err = r.redis.Set(key, string(b), dur).Err()
	if err != nil {
		// logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "redis_set", uc.ReqID)
		log.Println(err)
		return err
	}

	return err
}
