package main

import (
	"confdReWrite/src/backends"
	"confdReWrite/src/log"
	"confdReWrite/src/resource/template"
	"errors"
	"flag"
	"fmt"
	"path/filepath"
)

// confd config
type Config struct {
	backends.BackendsConfig
	template.ResourceConfig
	PrintVersion bool
	OneTime      bool
	LogLevel     string
	Watch        bool
}

var config Config

func init() {

	//confd basic config
	flag.BoolVar(&config.PrintVersion, "version", false, "print version and exit")
	flag.StringVar(&config.ConfDir, "confdir", "/etc/confd", "confd root dir")
	flag.BoolVar(&config.OneTime, "onetime", false, "run once and exit")
	flag.StringVar(&config.LogLevel, "log-level", "", "level which confd should log messages")
	flag.BoolVar(&config.Watch, "watch", false, "enable watch support")
	//confd backend config
	flag.StringVar(&config.Backend, "backend", "", "backend to use")
	flag.Var(&config.BackendNodes, "node", "list of backend nodes")
	flag.IntVar(&config.Interval, "interval", 600, "backend polling interval")
	flag.StringVar(&config.Username, "username", "", "the username to authenticate")
	flag.StringVar(&config.Password, "password", "", "the password to authenticate")
	//nacos
	flag.StringVar(&config.Namespace, "namespace", "", "the namespace in nacos ")
	flag.StringVar(&config.Group, "group", "DEFAULT_GROUP", "the group in nacos ")
	flag.StringVar(&config.AccessKey, "accessKey", "", "the accessKey to authenticate in nacos")
	flag.StringVar(&config.SecretKey, "secretKey", "", "the secretKey to authenticate in nacos ")

}

// initConfig initializes the confd configuration
func initConfig() error {

	// seting log level
	if config.LogLevel != "" {
		log.SetLevel(config.LogLevel)
	}
	log.Info(fmt.Sprintf("Basic config init,confdir:%s,log-level:%s,oneTime:%v,watch:%v",
		config.ConfDir, config.LogLevel, config.OneTime, config.Watch))

	if config.Backend == "" {
		return errors.New("please input -backend ")
	}
	if len(config.BackendNodes) == 0 {
		return errors.New("please input -node ")
	}

	switch config.Backend {
	case "nacos":
		if err := validNacosParam(&config); err != nil {
			return err
		}
		log.Info(fmt.Sprintf("Backend is %s,node:%s,interval:%d,namespace:%s,group:%s,accessKey:%s,secretKey:%s",
			config.Backend, config.BackendNodes, config.Interval, config.Namespace, config.Group, config.AccessKey, config.SecretKey))
	default:
		return errors.New("Unknow backend ")

	}
	//setting resource path
	config.ConfigDir = filepath.Join(config.ConfDir, "conf.d")
	config.TemplateDir = filepath.Join(config.ConfDir, "templates")
	return nil
}

func validNacosParam(config *Config) error {

	if config.Namespace == "" {
		return errors.New("please input -namespace ")
	}
	return nil

}
