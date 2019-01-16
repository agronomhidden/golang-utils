package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func ReadConfigFiles(path string) {
	var args = os.Args[1:]
	var host, hostErr = os.Hostname()

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	viper.AddConfigPath(path)
	viper.AddConfigPath(exPath)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("common")
	viper.ReadInConfig()

	if hostErr == nil {
		viper.SetConfigName(host)
		viper.MergeInConfig()
	}
	if len(args) > 0 {
		var _, err = os.Stat(args[0])
		if !os.IsNotExist(err) {
			viper.SetConfigFile(args[0])
			viper.MergeInConfig()
		} else {
			panic(fmt.Sprintf("Config file \"%s\" was not found", args[0]))
		}
	}
}

func GetPort() string {
	return viper.GetString("API.port")
}
