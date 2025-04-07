package redis

func Save(key string, value string) error {
	return getClient().Set(ctx, key, value, 0).Err()
}

func Get(key string) (string, error) {
	return getClient().Get(ctx, key).Result()
}
