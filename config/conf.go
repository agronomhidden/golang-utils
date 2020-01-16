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

func ReadConfigFilesV2(path, file string) {
		
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

	if file != "" {
		file = path + file
		var _, err = os.Stat(file)
		if !os.IsNotExist(err) {
			viper.SetConfigFile(file)
			viper.MergeInConfig()
		} else {
			panic(fmt.Sprintf("Config file \"%s\" was not found", file))
		}
	}
}

func GetPort() string {
	return viper.GetString("API.port")
}

func GetLogFile() string {
	return viper.GetString("logger.file")
}

func GetLogMaxSize() int {
	return viper.GetInt("logger.maxsize")
}

func GetLogMaxAge() int {
	return viper.GetInt("logger.maxage")
}
