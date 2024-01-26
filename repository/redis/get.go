package redis

// Get get value from redis by key
func (r *Redis) Get(key string, cb interface{}) (string, error) {
	return r.redis.Get(key).Result()

}
