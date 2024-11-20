package internal

import "github.com/gin-gonic/gin"

type ServerConfig struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}

func NewServer(config *Config) *gin.Engine {
	server := gin.Default()
	return server
}