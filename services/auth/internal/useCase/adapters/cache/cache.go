package cache

import "time"

type Cache interface {
	Set(key string, value interface{}, duration time.Duration)
	Get(key string) (value interface{}, err error)
	Delete(key string) (err error)
	Stop() (err error)
}
