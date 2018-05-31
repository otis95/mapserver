package config

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/ini.v1"
)

type Config struct {
	//HTTP
	HttpServerWin   string `ini:"http_server_win"`
	HttpServerLinux string `ini:"http_server_linux"`
	//LOG
	LogDirWin   string `ini:"log_dir_win"`
	LogDirLinux string `ini:"log_dir_linux"`
	LogPrefix   string `ini:"log_prefix"`
}

func (c Config) String() string {

	http := fmt.Sprintf("HTTP:[%v]/[%v]", c.HttpServerWin, c.HttpServerLinux)

	log := fmt.Sprintf("LOG:[win:%v]/[linux:%v]:[prefix:%v]", c.LogDirWin, c.LogDirLinux, c.LogPrefix)

	return http + ", " + log
}

//Read Server's Config Value from "path"
func ReadConfig(path string) (Config, error) {
	var config Config
	conf, err := ini.Load(path)
	if err != nil {
		log.Println("load config file fail!")
		return config, err
	}
	conf.BlockMode = false
	err = conf.MapTo(&config)
	if err != nil {
		log.Println("mapto config file fail!")
		return config, err
	}
	return config, nil
}
