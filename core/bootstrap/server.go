package bootstrap

import "github.com/gin-gonic/gin"

type ServerConfig struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}

func NewServer() *gin.Engine {
	server := gin.Default()
	return server
}
