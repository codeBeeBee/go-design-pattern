package singleton

type HungrySingleton struct{}

var singleton *HungrySingleton

func init() {
	singleton = &HungrySingleton{}
}

func GetInstance() *HungrySingleton {
	return singleton
}
