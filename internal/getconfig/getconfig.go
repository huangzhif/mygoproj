package getconfig

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"mygoproject/internal/logger"
)

func InitConfigure() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath("./configs/")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("file not found")
		} else {
			panic("something happend")
		}
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		logger.Info.Println("config file changed:", e.Name)
	})

	return v
}
