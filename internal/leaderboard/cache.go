package leaderboard

import (
	"time"
)

type cacheValue struct {
	data interface{}
	exp  time.Time
}

type inMemoryCache map[string]cacheValue

func (i inMemoryCache) Set(k string, v interface{}, exp time.Time) {

	if exp.Before(time.Now()) {
		return
	}

	i[k] = cacheValue{
		data: v,
		exp:  exp,
	}
}

func (i inMemoryCache) Get(k string) interface{} {

	if _, ok := i[k]; !ok {
		return nil
	}

	v := i[k]

	if v.exp.Before(time.Now()) {
		delete(i, k)
		return nil
	}

	return v.data
}
