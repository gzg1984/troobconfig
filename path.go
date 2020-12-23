package troobconfig

import (
	"log"

	"gopkg.in/ini.v1"
)

var globalPathPrefix string
var globalProjectBasePath string
var globalIndexBasePath string

func getPathConfig() error {
	configfile := getPathConfigFile()
	if len(configfile) == 0 {
		return nil
	}
	cfg, err := ini.Load(configfile)
	if err != nil {
		log.Printf("Cannot Find Path Config file, use default values\n")
		return err
	}

	if cfg != nil {
		/*prefix=/opt/file_root*/
		prefix := cfg.Section("").Key("prefix").String()
		if len(prefix) != 0 {
			globalPathPrefix = prefix
		}

		/*rootPath=project_base*/
		rootPath := cfg.Section("").Key("rootPath").String()
		if len(rootPath) != 0 {
			globalProjectBasePath = globalPathPrefix + "/" + rootPath
		}

		/*indexPath=index_base*/
		indexPath := cfg.Section("").Key("indexPath").String()
		if len(indexPath) != 0 {
			globalIndexBasePath = globalPathPrefix + "/" + indexPath
		}

	}
	return nil
}

/*InitGlobalBasePath should call before use all other file operation*/
func InitGlobalBasePath() error {
	err := getPathConfig()
	if err != nil {
		log.Printf("config file exist but Get DB From configfile failed:%v\n", err)
	}

	return nil
}
