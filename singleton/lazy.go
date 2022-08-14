package singleton

import "sync"

type LazySingleton struct{}

var (
	lazySingleton *LazySingleton
	once          = &sync.Once{}
)

func GetLazyInstance() *LazySingleton {
	if singleton == nil {
		once.Do(func() {
			lazySingleton = &LazySingleton{}
		})
	}
	return lazySingleton
}
