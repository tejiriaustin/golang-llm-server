package config

import (
	"fmt"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	conf *sync.Map
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
}

func New() *Config {
	conf := new(sync.Map)
	return &Config{conf: conf}
}

func (c *Config) SetEnv(key string, value any) *Config {
	c.conf.Store(key, value)

	fmt.Println(c.conf)
	return c
}

func (c *Config) GetEnv(key string) any {
	value, ok := c.conf.Load(key)
	if !ok {
		return nil
	}
	return value
}

func (c *Config) GetString(key string) string {
	value := c.GetEnv(key)
	if value == nil {
		return ""
	}

	if valueString, ok := value.(string); ok {
		return valueString
	}
	return ""
}

func (c *Config) GetInt(key string) int {
	value := c.GetEnv(key)
	if value == nil {
		return 0
	}

	if valueString, ok := value.(int); ok {
		return valueString
	}
	return 0
}

func (c *Config) GetFloat64(key string) float64 {
	value := c.GetEnv(key)
	if value == nil {
		return 0
	}

	if valueString, ok := value.(float64); ok {
		return valueString
	}
	return 0
}
