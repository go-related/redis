package settings

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configurations struct {
	Service1 Service1
	Service2 Service2
}
type Service1 struct {
	Port string
	Name string
}
type Service2 struct {
	Port string
	Name string
}

var ApplicationConfiguration Configurations

func init() {
	v := viper.New()
	v.SetConfigName("default")
	v.SetConfigType("yml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		// It's okay if there isn't a config file
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.Is(err, &configFileNotFoundError) {
			log.WithError(err).Errorf("error loading config file")
		}
	}
	v.AutomaticEnv()
	err := v.UnmarshalExact(&ApplicationConfiguration)
	if err != nil {
		log.WithError(err).Errorf("error mapping config file to type")
	}
}
