package config

import (
	"os"

	"github.com/sevengit-wq/skill-test/pkg/utils"
	"github.com/sirupsen/logrus"
)

type ServerConfig struct {
	Port        string
	LoggerLevel string
}

type ApiService struct {
	URL      string
	Username string
	Password string
}

type EnvironmentConfig struct {
	ServerConfig
	ApiService
}

type CoreResources struct {
	Log *logrus.Logger
	*EnvironmentConfig
}

func SetupCoreResources(paths ...string) (*CoreResources, error) {
	var loggerLevel *utils.LogrusLogger
	var envCfg EnvironmentConfig
	envCfg.Port = os.Getenv("SERVER_PORT")
	envCfg.LoggerLevel = os.Getenv("LOG_LEVEL")
	level := utils.ParseLevel(envCfg.LoggerLevel)
	loggerLevel = utils.NewLogrusLogger()
	loggerLevel.SetLogLevel(level)
	envCfg.Username = os.Getenv("USERNAME")
	envCfg.Password = os.Getenv("PASSWORD")
	envCfg.URL = os.Getenv("API_SERVICE_URL")

	return &CoreResources{
		Log:               loggerLevel.Logger,
		EnvironmentConfig: &envCfg,
	}, nil
}
