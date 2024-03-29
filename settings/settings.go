package settings

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Init() (err error) {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		zap.L().Fatal("Fatel err config file:%v\n", zap.Error(err))
		return
	}

	//实时监控配置文件并更新
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		zap.L().Info("配置文件被修改...")
	})
	return
}
