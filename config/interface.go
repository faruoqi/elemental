package config

type Config interface {
	GetConfig(key string) string
	Info() string
	Init()
}
