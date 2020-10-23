package config

import (
	"github.com/spf13/viper"
	"log"
)

type ConfigViper struct {
	engine *viper.Viper
	env    string
}

func NewConfig(env string) *ConfigViper {
	engine := viper.New()
	return &ConfigViper{engine: engine, env: env}
}

func (c *ConfigViper) GetConfig(key string) string {

	val := c.engine.Get(c.env + "." + key)
	return val.(string)

}

func (c *ConfigViper) Init() {

	c.engine.AddConfigPath(".")
	c.engine.SetConfigName("config")
	c.engine.SetConfigType("json")

	if err := c.engine.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file %v\n", err)
	}
	log.Printf("Using config file %v \n", c.engine.ConfigFileUsed())

}

func (c ConfigViper) Info() string {
	return "Using Viper Framework"
}
