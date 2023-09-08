package configurations

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configurations struct {
	Service1 struct {
		Port string
		Name string
	}
	Service2 struct {
		Port string
		Name string
	}
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
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.WithError(err).Errorf("error loading config file")
		}
	}
	v.AutomaticEnv()
	err := v.UnmarshalExact(&ApplicationConfiguration)
	if err != nil {
		log.WithError(err).Errorf("error mapping config file to type")
	}
}
