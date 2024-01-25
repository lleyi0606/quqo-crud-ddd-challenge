package config

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Configuration viper.Viper

func LoadConfiguration() {
	configuration := *viper.New()
	configuration.SetConfigName("application")                         // name of config file
	configuration.AddConfigPath(os.ExpandEnv("infrastructure/config")) // path to look for config file in in
	// configuration.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))     // replace . with _ in env
	// configuration.SetConfigType("yaml")                                // type of config file, not required if file has extension in name
	if err := configuration.ReadInConfig(); err != nil { // find and read config file
		panic(err)
	}
	log.Print("load configuration ran")
	configuration.WatchConfig()                           // read config file while running
	configuration.OnConfigChange(func(e fsnotify.Event) { // config file change
		//	glog.Info("App Config file changed %s:", e.Name)
	})
	Configuration = configuration
}
