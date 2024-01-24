package util

import (
	"npos-demo/models"
	"npos-demo/pkg/msg"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Viper() {
	v := viper.New()
	v.AddConfigPath("config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		log.Errorf("Fatal error that read the global config file failed: %s", err.Error())
		panic(msg.PMSG)
	}
	if err := v.Unmarshal(&models.Config); err != nil {
		log.Errorf("%s when obtain the global config: %s", msg.DUNFMSG, err.Error())
		panic(msg.PMSG)
	}
}
