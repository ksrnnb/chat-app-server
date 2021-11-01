package controller

type ISession interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	Save() error
	Disable() error
}
